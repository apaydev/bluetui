package bluetooth

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

// mockDbusConn is a fake dbusConn used only for tests.
type mockDbusConn struct {
	objects map[string]mockBusObject
	closed  bool
}

func (m *mockDbusConn) Object(destination string, path dbus.ObjectPath) dbusObject {
	// Return a no-op object if not found (can be improved to simulate errors)
	return &mockBusObject{}
}

func (m *mockDbusConn) Close() error {
	m.closed = true
	return nil
}

// mockBusObject simulates a dbus.BusObject for testing.
type mockBusObject struct {
	CallHistory []string
}

func (m *mockBusObject) Call(method string, flags dbus.Flags, args ...any) *dbus.Call {
	m.CallHistory = append(m.CallHistory, method)
	// If the args contain the string "error", simulate an error
	if len(args) > 0 && args[0] == "error" {
		return &dbus.Call{Err: fmt.Errorf("simulated error")}
	}
	// Fake successful call
	return &dbus.Call{}
}

// NewMockConnection is a factory that returns a mock connection for testing.
func NewMockConnection() (dbusConn, error) {
	return &mockDbusConn{}, nil
}
