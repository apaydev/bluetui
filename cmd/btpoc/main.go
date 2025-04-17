package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/apaydev/bluetui/internal/bluetooth"
)

func main() {
	validCmds := []string{"discover", "pair", "connect", "disconnect"}
	usageStr := fmt.Sprintf("Command to execute. Options: %s", strings.Join(validCmds, ", "))
	// I need to read the function to be executed through flags.
	cmd := flag.String("cmd", "discover", usageStr)
	flag.Parse()
	args := flag.Args()

	if !slices.Contains(validCmds, *cmd) {
		fmt.Fprintf(os.Stderr, "Invalid command: %s. Valid commands are: %s", *cmd, strings.Join(validCmds, ", "))
		os.Exit(1)
	}

	// I want to work with my bluetooth adapter. So, I need to get the
	// object for it, which will give me interfaces, devices and methods.
	adapter, err := bluetooth.NewAdapter("", "", bluetooth.NewSystemBusConnection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get bluetooth adapter: %v\n", err)
		os.Exit(1)
	}

	// I need to defer the cleaning of resources.
	defer func() {
		if cerr := adapter.Close(); cerr != nil {
			errors.Join(err, cerr)
		}
	}()

	// Discover
	err = adapter.Discover()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to discover devices: %v\n", err)
		os.Exit(1)
	}

	devices, err := adapter.Devices()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get devices: %v\n", err)
		os.Exit(1)
	}

	if devices == nil {
		fmt.Println("No devices found.")
		os.Exit(0)
	}

	if *cmd == "discover" {
		// Show all devices
		fmt.Println("\nDiscovered Devices:")
		i := 0
		for _, d := range devices {
			fmt.Printf("[%d] %s (%s)\n", i+1, d.Name(), d.Address())
			i++
		}
	}

	// Pairing
	if *cmd == "pair" {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "please, provide a device address to pair with.\n")
			os.Exit(1)
		}

		addr := args[0]
		err = adapter.Pair(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to pair with device %s: %v\n", addr, err)
			os.Exit(1)
		}

		os.Exit(0)

	}

	// Connecting
	if *cmd == "connect" {
		args := flag.Args()
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "please, provide a device address to connect with.\n")
			os.Exit(1)
		}

		addr := args[0]
		err = adapter.Connect(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to connect with device %s: %v\n", addr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Disconnecting
	if *cmd == "disconnect" {
		args := flag.Args()
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "please, provide a device address to disconnect from.\n")
			os.Exit(1)
		}

		addr := args[0]
		err = adapter.Disconnect(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to disconnect from device %s: %v\n", addr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	//
	//
	// // Try connecting (with retry)
	// for range 3 {
	// 	fmt.Println("Attempting to connect...")
	// 	err = device.Call(deviceInterface+".Connect", 0).Err
	// 	if err == nil {
	// 		fmt.Println("Connected to:", selected.Name)
	// 		break
	// 	}
	// 	fmt.Println("Connect failed:", err)
	// 	time.Sleep(2 * time.Second)
	// }
}
