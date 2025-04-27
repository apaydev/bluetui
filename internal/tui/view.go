package tui

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	listView := lipgloss.NewStyle().PaddingTop(2).Render(m.list.View())
	helpView := lipgloss.NewStyle().Padding(0, 2).Render(m.help.View(m.keys))

	return lipgloss.JoinVertical(lipgloss.Left, listView, helpView)
}
