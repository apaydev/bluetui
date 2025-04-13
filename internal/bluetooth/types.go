package bluetooth

import (
	"github.com/godbus/dbus/v5"
)

// device describes a discovered Bluetooth device. It is supposed
// to be OS-agnostic, but I will not be able to confirm that until
// I explore other platforms.
type device struct {
	Name    string
	Address string
	Path    dbus.ObjectPath
}
