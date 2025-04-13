package bluetooth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/godbus/dbus/v5"
)

// These values were obtained through the busctl command:
const (
	// busctl list gives us this first value, which is the well-known name
	// name for the bluetooth D-Bus.
	bluezDestination = "org.bluez"
	// busctl tree org.bluez gives us the object path to our BT adapter,
	// as well as the paths to all of the devices connected to it.
	adapterPath = "/org/bluez/hci0"
	// busctl introspect org.bluez /org/bluez/hci0 gives us the interfaces
	// and methods available for the BT adapter, and for the Device adapter.
	adapterInterface = "org.bluez.Adapter1"
	deviceInterface  = "org.bluez.Device1"
	// Standard interface to work with properties of D-Bus objects.
	propertiesInterface = "org.freedesktop.DBus.Properties"
)

// linuxAdapter is a Linux-specific implementation of the Adapter interface.
// It uses D-Bus to communicate with the BlueZ stack.
type linuxAdapter struct {
	adapterBase
	conn       *dbus.Conn
	adapterObj dbus.BusObject
}

// newSystemBusConn creates a new connection to the system-wide D-Bus.
func newSystemBusConn() (*dbus.Conn, error) {
	// Conecting to the system-wide D-Bus
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}
	return conn, err
}

// NewAdapter creates a new Linux-specific Bluetooth adapter.
func NewAdapter(destination, path string) (Adapter, error) {
	conn, err := newSystemBusConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}

	dest := destination
	if dest == "" {
		dest = bluezDestination
	}

	pth := path
	if pth == "" {
		pth = adapterPath
	}

	return &linuxAdapter{
		conn:        conn,
		adapterBase: adapterBase{destination: dest, path: pth},
		adapterObj:  conn.Object(dest, dbus.ObjectPath(pth)),
	}, nil
}

// Discover starts the discovery process for Bluetooth devices.
func (b *linuxAdapter) Discover() (err error) {
	fmt.Println("Starting discovery. This may take a few seconds ...")
	// Pretty self explainatory. Begin scanning for devices, and defer
	// the call to stop that process.
	err = b.adapterObj.Call(adapterInterface+".StartDiscovery", 0).Err
	if err != nil {
		return fmt.Errorf("failed to start discovery process: %w", err)
	}

	// Since stopping the discovery process can yield an error, we should
	// handle it accordingly.
	defer func() {
		if cerr := b.adapterObj.Call(adapterInterface+".StopDiscovery", 0).Err; cerr != nil {
			errors.Join(err, errors.Join(err, fmt.Errorf("failed to stop discovery process: %w", cerr)))
		}
	}()

	// Give the process some time to run.
	time.Sleep(5 * time.Second)

	// Get the devices that were discovered.
	err = b.getDevicesInfo()
	if err != nil {
		return fmt.Errorf("failed to get info for discovered devices: %w", err)
	}

	return nil
}

// getDevicesInfo is a helper function that retrieves the information of discovered
// devices in our BlueZ object.
func (b *linuxAdapter) getDevicesInfo() error {
	var objs map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	// Call the GetManagedObjects method to get all the devices that BlueZ
	// knows about at this point.
	objManager := b.conn.Object(b.destination, "/")
	err := objManager.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&objs)
	if err != nil {
		return fmt.Errorf("failed to get managed objects: %w", err)
	}

	devices := make(map[string]device)

	for path, ifaceMap := range objs {
		if dev, ok := ifaceMap[deviceInterface]; ok {
			var name, addr string

			if val, ok := dev["Address"]; ok {
				addr = val.Value().(string)
			} else {
				continue
			}

			if val, ok := dev["Name"]; ok {
				name = val.Value().(string)
			} else {
				name = "<unknown>"
			}

			devices[addr] = device{Name: name, Address: addr, Path: path}
		}
	}

	b.devices = devices

	return nil
}

// Pair attempts to pair with a Bluetooth device using its address.
func (b *linuxAdapter) Pair(deviceAddress string) error {
	if deviceAddress == "" {
		return errors.New("a device address is required")
	}

	devicePath := b.devices[deviceAddress].Path
	device := b.conn.Object(b.destination, devicePath)

	// Try pairing
	err := device.Call(deviceInterface+".Pair", 0).Err
	if err != nil {
		if strings.Contains(err.Error(), "Already Exists") {
			fmt.Println("Already paired. Skipping pairing step.")
		} else {
			return fmt.Errorf("pairing with device at addr %s failed: %w", deviceAddress, err)
		}
	} else {
		fmt.Println("Paired successfully.")
	}

	return nil
}

// Trust attempts to trust a Bluetooth device using its address.
func (b *linuxAdapter) Trust(deviceAddress string) error {
	if deviceAddress == "" {
		return errors.New("a device address is required")
	}

	devicePath := b.devices[deviceAddress].Path
	device := b.conn.Object(b.destination, devicePath)

	// Trust the device
	err := device.Call(propertiesInterface+".Set", 0, deviceInterface, "Trusted", dbus.MakeVariant(true)).Err
	if err != nil {
		return fmt.Errorf("trusting device at addr %s failed: %w", deviceAddress, err)
	}

	fmt.Println("Device marked as trusted.")
	return nil
}

// Connect attempts to connect to a Bluetooth device using its address.
func (b *linuxAdapter) Connect(deviceAddress string) error {
	if deviceAddress == "" {
		return errors.New("a device address is required")
	}

	devicePath := b.devices[deviceAddress].Path
	device := b.conn.Object(b.destination, devicePath)

	// Try connecting
	err := device.Call(deviceInterface+".Connect", 0).Err
	if err != nil {
		return fmt.Errorf("connecting to device at addr %s failed: %w", deviceAddress, err)
	}

	fmt.Println("Connected successfully.")
	return nil
}

// Disconnect attempts to disconnect from a Bluetooth device using its address.
func (b *linuxAdapter) Disconnect(deviceAddress string) error {
	if deviceAddress == "" {
		return errors.New("a device address is required")
	}

	devicePath := b.devices[deviceAddress].Path
	device := b.conn.Object(b.destination, devicePath)

	err := device.Call(deviceInterface+".Disconnect", 0).Err
	if err != nil {
		return fmt.Errorf("disconnecting from device at addr %s failed: %w", deviceAddress, err)
	}

	fmt.Println("Disconnected successfully.")
	return nil
}

// Devices returns a list of discovered Bluetooth devices.
func (b *linuxAdapter) Devices() ([]device, error) {
	if len(b.devices) == 0 {
		return nil, errors.New("no devices found")
	}

	devices := make([]device, 0, len(b.devices))
	for _, dev := range b.devices {
		devices = append(devices, dev)
	}

	return devices, nil
}

// Close closes the connection to the D-Bus used by the adapter.
func (b *linuxAdapter) Close() error {
	if b.conn != nil {
		err := b.conn.Close()
		if err != nil {
			return fmt.Errorf("failed to close D-Bus connection: %w", err)
		}
	}
	return nil
}
