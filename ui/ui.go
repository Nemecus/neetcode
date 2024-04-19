package ui

import (
	"strings"

	"github.com/Nemecus/neetcode/ui/components/dynamicarray"
	"github.com/Nemecus/neetcode/ui/components/menu"
	"github.com/Nemecus/neetcode/ui/context"
	"github.com/Nemecus/neetcode/ui/keys"
	"github.com/Nemecus/neetcode/ui/styles"
	"github.com/Nemecus/neetcode/ui/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type mainContentModel struct {
	menu         menu.Model
	dynamicArray dynamicarray.Model
}

type Model struct {
	ctx  context.ProgramContext
	keys keys.AppKeyMap
	mainContentModel
}

func NewModel() Model {
	m := Model{
		keys: keys.Keys,
	}

	m.ctx = context.ProgramContext{}
	m.ctx.Theme = theme.ParseTheme()
	m.ctx.Styles = styles.InitStyles(m.ctx.Theme)
	m.ctx.State = context.MenuView
	m.menu = menu.NewModel(m.ctx)

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd             tea.Cmd
		dynamicArrayCmd tea.Cmd
		menuCmd         tea.Cmd
		cmds            []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.ctx.Error = nil

		switch {
		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case menu.AnswerMsg:
		switch msg {
		case "Dynamic Array":
			m.ctx.State = context.DynamicArrayView
			m.dynamicArray = dynamicarray.NewModel(m.ctx)
		}
	case dynamicarray.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	}

	m.syncProgramContext()

	switch m.ctx.State {
	case context.MenuView:
		m.menu, menuCmd = m.menu.Update(msg)
	case context.DynamicArrayView:
		m.dynamicArray, dynamicArrayCmd = m.dynamicArray.Update(msg)
	}
	cmds = append(cmds, cmd, menuCmd, dynamicArrayCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString("Welcome to NeetCode Problems!\n")
	s.WriteString(m.menu.View())
	s.WriteString("\n")
	return m.ctx.Styles.Common.MainStyle.Render(s.String())
}

func (m *Model) syncProgramContext() {
	m.menu.UpdateProgramContext(&m.ctx)
	m.dynamicArray.UpdateProgramContext(&m.ctx)
}
