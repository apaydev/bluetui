package bluetooth

import "context"

// Adapter defines the behavior of a platform-specific Bluetooth adapter.
//
// Implementations are expected to handle OS-specific logic behind methods like
// Discover, Connect, etc.
type Adapter interface {
	Discover(context.Context) error
	Pair(addr string) error
	Trust(addr string) error
	Connect(addr string) error
	Disconnect(addr string) error
	Devices() ([]Device, error)
	Close() error
	// These methods are used to get the adapter's properties.
	// NOTE: I have not found a way to make them generic for all implementations
	// of this interface. Tried placing them into the adapterBase struct, but it
	// does not work.
	Destination() string
	Path() string
}

// adapterBase provides a base implementation for our bluetooth adapters. It
// contains common fields that can be shared across different platforms.
type adapterBase struct {
	destination string
	path        string
	devices     map[string]Device
}

// Device describes a discovered Bluetooth Device.
type Device struct {
	name    string
	address string
	path    string
}

// Name returns the name of the Bluetooth device.
func (d Device) Name() string {
	return d.name
}

// Address returns the address of the Bluetooth device.
func (d Device) Address() string {
	return d.address
}

// Path returns the D-Bus path of the Bluetooth device.
func (d Device) Path() string {
	return d.path
}

// METHODS REQUIRED SO THAT THIS CAN BE USED AS A LIST ITEM

func (d Device) Title() string       { return d.name }
func (d Device) Description() string { return d.address }
func (d Device) FilterValue() string { return d.name }

// NOTE: Func used only to test the integration with the list bubble. This
// will be either modified or removed.
func GetDevices() []Device {
	return []Device{
		{name: "Device 1", address: "00:00:00:00:00:01", path: "/org/bluez/hci0/dev_00_00_00_00_00_01"},
		{name: "Device 2", address: "00:00:00:00:00:02", path: "/org/bluez/hci0/dev_00_00_00_00_00_02"},
		{name: "Device 3", address: "00:00:00:00:00:03", path: "/org/bluez/hci0/dev_00_00_00_00_00_03"},
		{name: "Device 4", address: "00:00:00:00:00:04", path: "/org/bluez/hci0/dev_00_00_00_00_00_04"},
		{name: "Device 5", address: "00:00:00:00:00:05", path: "/org/bluez/hci0/dev_00_00_00_00_00_05"},
		{name: "Device 6", address: "00:00:00:00:00:06", path: "/org/bluez/hci0/dev_00_00_00_00_00_06"},
	}
}
