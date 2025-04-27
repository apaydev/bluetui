package tui

import (
	"github.com/apaydev/bluetui/internal/bluetooth"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	list list.Model
	// TODO: This will be used to render the selected option with a different
	// style.
	cursor int
	keys   keyMap
	help   help.Model
}

// NewModel defines the app's initial state
func NewModel() model {
	m := model{
		keys: newKeyMap(),
		help: help.New(),
	}

	// Setup help
	m.help = styledHelp(m.help)

	// Make initial list of items
	devices := bluetooth.GetDevices()

	items := make([]list.Item, len(devices))
	for i := range devices {
		items[i] = devices[i]
	}

	// Setup List
	deviceList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	deviceList.Title = "Bluetooth Devices"
	deviceList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#25A065")).
		Padding(0, 1)
	deviceList.SetShowHelp(false)

	m.list = deviceList

	return m
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
