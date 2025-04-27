package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	discover   key.Binding
	trust      key.Binding
	pair       key.Binding
	connect    key.Binding
	disconnect key.Binding
	filter     key.Binding
	quit       key.Binding
	up         key.Binding
	down       key.Binding
	nextPage   key.Binding
	prevPage   key.Binding
	home       key.Binding
	end        key.Binding
	help       key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.up, k.down, k.pair, k.connect, k.help, k.quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.up, k.down, k.prevPage, k.nextPage}, // first column
		{k.help, k.quit},                       // second column
	}
}

func newKeyMap() keyMap {
	return keyMap{
		up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		prevPage: key.NewBinding(
			key.WithKeys("left", "h", "pgup"),
			key.WithHelp("←/h/pgup", "prev page"),
		),
		nextPage: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l/pgdn", "next page"),
		),
		help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}
