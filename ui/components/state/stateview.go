package state

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
		m.status = m.runState()
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
		s = m.status + "\nPress Enter"
	}

	return m.ctx.Styles.Common.MainStyle.Copy().Render(s)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func (m Model) runState() string {
	var stateString strings.Builder

	stateString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	stateString.WriteString("\n\n")

	stateString.WriteString("New Document Created\n")
	document1 := newDocument()

	stateString.WriteString(fmt.Sprintf("Current state of doc: %#v\n", document1.getState()))

	stateString.WriteString("Running Publish\n")
	document1.publish()
	stateString.WriteString(fmt.Sprintf("Current state of doc: %#v\n", document1.getState()))

	stateString.WriteString("Running Publish\n")
	document1.publish()
	stateString.WriteString(fmt.Sprintf("Current state of doc: %#v\n", document1.getState()))

	stateString.WriteString("Setting approval to true\n")
	document1.setApproval(true)

	stateString.WriteString("Running Publish\n")
	document1.publish()
	stateString.WriteString("Running Publish\n")
	document1.publish()
	stateString.WriteString(fmt.Sprintf("Current state of doc: %#v\n", document1.getState()))

	stateString.WriteString("Running Publish\n")
	document1.publish()
	stateString.WriteString(fmt.Sprintf("Current state of doc: %#v\n\n", document1.getState()))

	return stateString.String()
}
