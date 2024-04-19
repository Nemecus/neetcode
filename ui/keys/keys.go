package keys

import "github.com/charmbracelet/bubbles/key"

type AppKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Help  key.Binding
	Enter key.Binding
	Quit  key.Binding
}

var Keys = AppKeyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
	Help:  key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "toggle help")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select item")),
	Quit:  key.NewBinding(key.WithKeys("esc", "ctrl+c"), key.WithHelp("esc/ctrl+c", "quit")),
}

func (k AppKeyMap) SelectHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Quit}
}

func (k AppKeyMap) AnswerHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Quit}
}

func (k AppKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Quit}
}

func (k AppKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.NavigationKeys(),
		k.EnterAndQuitKeys(),
	}
}

func (k AppKeyMap) NavigationKeys() []key.Binding {
	return []key.Binding{k.Up, k.Down}
}

func (k AppKeyMap) QuitAndHelpKeys() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k AppKeyMap) EnterAndQuitKeys() []key.Binding {
	return []key.Binding{k.Enter, k.Quit}
}
