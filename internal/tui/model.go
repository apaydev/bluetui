package tui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// TODO: Change this for the actual device struct.
	devices []string
	// TODO: This will be used to render the selected option with a different
	// style.
	cursor int
	keys   keyMap
	help   help.Model
}

// NewModel defines the app's initial state
func NewModel() model {
	m := model{
		devices: []string{"Willen II", "Galaxy Buds", "Logitech MX Master 3"},
		keys:    newKeyMap(),
		help:    help.New(),
	}

	m.help = styledHelp(m.help)
	return m
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
