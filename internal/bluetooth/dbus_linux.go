package bluetooth

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

// dbusObject abstracts the functions that are currently used in the code
// to interact with D-Bus objects. Useful for dependency injection and mocking.
// An example of an object is our bluetooth adapter, through which we call
// all of our methods and properties.
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
// dbusConn interface. dbus.Conn already implements both the Object and
// Close methods. I just need to make some wraps in my custom types so that
// this works.
type defaultDbusConn struct {
	*dbus.Conn
}

// Object is required to give me control on the return type of the conn.Object
// method. My linux adapter interface expects this to return a dbusObject, not
// a *dbus.BusObject.
func (c *defaultDbusConn) Object(dest string, path dbus.ObjectPath) dbusObject {
	return c.Conn.Object(dest, path)
}

// DbusConnectionFactory is a function type that creates a dbusConn. This will
// be fed to the NewAdapter functions, so that we can decide whether we should
// mock or not the connection.
type DbusConnectionFactory func() (dbusConn, error)

// NewSystemBusConnection creates a real connection to the system bus.
func NewSystemBusConnection() (dbusConn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}
	return &defaultDbusConn{conn}, nil
}
