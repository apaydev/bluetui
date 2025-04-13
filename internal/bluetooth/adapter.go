package bluetooth

// Adapter defines the behavior of a platform-specific Bluetooth adapter.
//
// Implementations are expected to handle OS-specific logic behind methods like
// Discover, Connect, etc.
type Adapter interface {
	Discover() error
	Pair(addr string) error
	Trust(addr string) error
	Connect(addr string) error
	Disconnect(addr string) error
	Devices() ([]device, error)
}

// NewAdapter creates a new Bluetooth adapter instance based on the
// platform. By default, it returns a Linux adapter.
func NewAdapter(destination, path string) (Adapter, error) {
	return NewLinuxAdapter(destination, path)
}

// adapterBase provides a base implementation for our bluetooth adapters. It
// contains common fields that can be shared across different platforms.
type adapterBase struct {
	destination string
	path        string
	devices     map[string]device
}
