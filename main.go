package main

import (
	"os"

	"github.com/Nemecus/neetcode/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func createModel() ui.Model {
	return ui.NewModel()
}

func main() {
	model := createModel()

	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
