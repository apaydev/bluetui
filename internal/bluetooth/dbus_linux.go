package bluetooth

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

// device describes a discovered Bluetooth device.
type device struct {
	Name    string
	Address string
	Path    dbus.ObjectPath
}

// dbusObject abstracts the functions that are currently used in the code
// to interact with D-Bus objects. Useful for dependency injection and mocking.
type dbusObject interface {
	Call(method string, flags dbus.Flags, args ...any) *dbus.Call
}

// dbusConn abstracts the functions that are currently used in the code
// to interact with D-Bus connections. Also useful for mocking.
type dbusConn interface {
	Object(dest string, path dbus.ObjectPath) dbusObject
	Close() error
}

// defaltDbusConn is a wrapper around the dbus.Conn type to implement the
// dbusConn interface.
type defaultDbusConn struct {
	*dbus.Conn
}

// Object is required to give me controll on the return type of the conn.Object
// method. My linux adapter interface expects this to return a dbusObject, not
// a *dbus.BusObject.
func (c *defaultDbusConn) Object(dest string, path dbus.ObjectPath) dbusObject {
	return c.Conn.Object(dest, path)
}

// DbusConnectionFactory is a function type that creates a dbusConn. Useful
// for creating mocks.
type DbusConnectionFactory func() (dbusConn, error)

// NewSystemBusConnection creates a real connection to the system bus.
func NewSystemBusConnection() (dbusConn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}
	return &defaultDbusConn{conn}, nil
}
