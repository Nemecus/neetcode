package facade

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
		m.status = m.runFacade()
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

func (m Model) runFacade() string {
	var facadeString strings.Builder

	facadeString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	facadeString.WriteString("\n\n")

	driveThru := newDriveThruFacade()
	facadeString.WriteString("I would like to order a Burger and Fries.\n")
	dineInOrder := driveThru.takeOrder("Burger and Fries", false)

	facadeString.WriteString(fmt.Sprintf("I have your %s\n\n", dineInOrder.getFood()))

	facadeString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 2"))
	facadeString.WriteString("\n\n")
	facadeString.WriteString("I would like to order a Pizza for takeout.\n")
	takeOutOrder := driveThru.takeOrder("Pizza", true)

	facadeString.WriteString(fmt.Sprintf("I have your %s\n\n", takeOutOrder.getFood()))

	return facadeString.String()
}
