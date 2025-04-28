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
// key.Map interface. It is empty because we are going to use the default
// one for the list bubble, which has some functionalities integrated (like
// the clear filter option) that would be a paint to rebuild.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

// newKeyMap returns the default help keymap to be used in the app.
// NOTE: Need to add the missing keybindings
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
		filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		pair: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "pair"),
		),
		connect: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "connect"),
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

// filterKeyMap is only used for the filter input view
type filterKeyMap struct {
	apply  key.Binding
	cancel key.Binding
}

// ShortHelp returns keys for the mini help menu while filtering
func (f filterKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "apply filter")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel")),
	}
}

// FullHelp returns nothing because we don't need full help while filtering
func (f filterKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func newFilterKeyMap() filterKeyMap {
	return filterKeyMap{
		apply: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "apply filter"),
		),
		cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel filter"),
		),
	}
}
