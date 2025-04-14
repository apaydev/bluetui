package bluetooth

import "testing"

func TestNewAdapterWithMock(t *testing.T) {
	testCases := []struct {
		name        string
		destination string
		path        string
	}{
		{name: "Custom destination/path", destination: "org.bluez", path: "/org/bluez/hci0"},
		{name: "Default destination/path", destination: "", path: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			adapter, err := NewAdapter(tc.destination, tc.path, NewMockConnection)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			if adapter == nil {
				t.Fatal("expected adapter to be created")
			}
		})
	}
}
