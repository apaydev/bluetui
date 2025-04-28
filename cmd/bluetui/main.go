package main

import (
	"fmt"
	"os"

	"github.com/apaydev/bluetui/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Logging functionality
	if os.Getenv("DEBUG") == "true" {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
