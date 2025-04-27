package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

const (
	wrapperEllipsis = lipgloss.Color("#FFCCE1")
	separatorDot    = lipgloss.Color("#FF66A6")
	helpKey         = lipgloss.Color("#999999")
	helpDesc        = lipgloss.Color("#808080")
)

var appStyle = lipgloss.NewStyle().Padding(1, 2)

func styledHelp(help help.Model) help.Model {
	// The ellipsis is the "..." shown when text is truncated.
	help.Styles.Ellipsis = lipgloss.NewStyle().Foreground(lipgloss.Color(wrapperEllipsis))
	// Styles shown for the short help menu.
	help.Styles.ShortSeparator = lipgloss.NewStyle().Foreground(lipgloss.Color(separatorDot))
	help.Styles.ShortKey = lipgloss.NewStyle().Foreground(lipgloss.Color(helpKey))
	help.Styles.ShortDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(helpDesc))

	// Styles shown for the full help menu.
	help.Styles.FullSeparator = lipgloss.NewStyle().Foreground(lipgloss.Color(separatorDot))
	help.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color(helpKey))
	help.Styles.FullDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(helpDesc))

	return help
}
