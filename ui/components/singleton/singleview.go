package singleton

import (
	"fmt"
	"strings"

	"github.com/Nemecus/neetcode/ui/context"
	"github.com/Nemecus/neetcode/ui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx    *context.ProgramContext
	keys   keys.AppKeyMap
	done   bool
	status string
}

type AnswerMsg string

func NewModel(ctx context.ProgramContext) Model {
	return Model{
		ctx:    &ctx,
		keys:   keys.Keys,
		done:   false,
		status: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if !m.done {
		m.status = m.runSingletonPattern()
		m.done = true
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			return m, func() tea.Msg { return AnswerMsg("Return") }
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s string
	if m.done {
		s = m.status + "Press Enter"
	}

	return m.ctx.Styles.Common.MainStyle.Copy().Render(s)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func (m Model) runSingletonPattern() string {
	var singleString strings.Builder

	singleString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Singleton"))
	singleString.WriteString("\n\n")
	singleString.WriteString("Creating Singleton Instance\n")

	var singleton ValueServiceSingleton

	s1 := singleton.GetService()
	singleString.WriteString(fmt.Sprintf("Getting Value: %s\n", s1.GetValue()))
	singleString.WriteString("Setting Value to 'a value string'\n")
	s1.SetValue("a value string")
	singleString.WriteString(fmt.Sprintf("Getting Value: %s\n", s1.GetValue()))

	singleString.WriteString("Creating Second Singleton Instance\n")
	s2 := singleton.GetService()
	singleString.WriteString(fmt.Sprintf("Getting Value from second instance: %s\n\n", s2.GetValue()))

	return singleString.String()
}
