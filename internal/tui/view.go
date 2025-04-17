package tui

import "fmt"

func (m model) View() string {
	// The header
	s := "Select a device to connect\n\n"

	// Iterate over our devices
	for i, choice := range m.devices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	helpView := m.help.View(m.keys)

	// Send the UI for rendering
	return s + "\n" + helpView
}
