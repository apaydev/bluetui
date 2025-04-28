package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// NOTE: Need to check if I can simplify this.
func (m model) View() string {
	listView := m.list.View()
	helpView := lipgloss.NewStyle().PaddingLeft(2).Render(m.help.View(m.keys))

	switch {
	case m.list.ShowHelp():
		return lipgloss.NewStyle().
			PaddingTop(1).
			Render(listView)
	case m.list.SettingFilter():
		return lipgloss.NewStyle().
			PaddingTop(1).
			Render(listView + "\n" + lipgloss.NewStyle().
				Padding(0, 2).
				Render(m.help.View(m.filterKeys)),
			)
	default:
		return lipgloss.NewStyle().
			PaddingTop(1).
			Render(listView + "\n" + helpView)
	}
}
