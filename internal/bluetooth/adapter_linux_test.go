package bluetooth

import (
	"fmt"
	"testing"
)

func TestNewAdapterWithMock(t *testing.T) {
	type args struct {
		destination string
		path        string
	}

	testCases := []struct {
		name     string
		args     args
		expected Adapter
	}{
		{
			name: "Custom destination/path",
			args: args{
				destination: "org.bluez",
				path:        "/org/bluez/hci0",
			},
			expected: &linuxAdapter{
				adapterBase: adapterBase{
					destination: "org.bluez",
					path:        "/org/bluez/hci0",
					devices:     make(map[string]Device),
				},
				conn:       &mockDbusConn{},
				adapterObj: &mockBusObject{},
			},
		},
		{
			name: "Default destination/path",
			args: args{
				destination: "",
				path:        "",
			},
			expected: &linuxAdapter{
				adapterBase: adapterBase{
					destination: "org.bluez",
					path:        "/org/bluez/hci0",
					devices:     make(map[string]Device),
				},
				conn:       &mockDbusConn{},
				adapterObj: &mockBusObject{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			adapter, err := NewAdapter(tc.args.destination, tc.args.path, NewMockConnection)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			if adapter == nil {
				t.Fatal("expected adapter to be created")
			}

			if destination := adapter.Destination(); destination != tc.expected.Destination() {
				fmt.Printf("expected destination %s, got: %s", tc.expected.Destination(), destination)
			}

			if devicePath := adapter.Path(); devicePath != tc.expected.Path() {
				fmt.Printf("expected path %s, got: %s", tc.expected.Path(), devicePath)
			}
		})
	}
}
