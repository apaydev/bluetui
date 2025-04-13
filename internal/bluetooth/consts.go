package bluetooth

// This values were obtained through the busctl command
const (
	// busctl list gives us this first value, which is the well known name
	// name for the bluetooth D-Bus.
	bluezDestination = "org.bluez"
	// busctl tree org.bluez gives us the object path to our BT adapter,
	// as well as the paths to all of the devices connected to it.
	adapterPath = "/org/bluez/hci0"
	// busctl introspect org.bluez /org/bluez/hci0 gives us the interfaces
	// and methods available for the BT adapter.
	adapterInterface = "org.bluez.Adapter1"
	deviceInterface  = "org.bluez.Device1"
	// Standard interfae to work with properties of D-Bus objects.
	propertiesInterface = "org.freedesktop.DBus.Properties"
)
