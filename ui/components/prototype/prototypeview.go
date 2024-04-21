package prototype

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
		m.status = m.runPrototype()
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

func (m Model) runPrototype() string {
	var shapeString strings.Builder

	shapeString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Prototype"))
	shapeString.WriteString("\n\n")

	shapeString.WriteString("Creating square of length 10\n")
	square := &square{width: 10, height: 10}

	shapeString.WriteString("Cloning the square\n")
	anotherSquare := square.clone()

	shapeString.WriteString(fmt.Sprintf("Are the squares equal? %t\n", (square == anotherSquare)))

	shapeString.WriteString("Creating rectangle of width 10 and height 20\n")
	rectangle := &rectangle{width: 10, height: 20}

	shapeString.WriteString("Cloning the rectangle\n")
	anotherRectangle := rectangle.clone()

	shapeString.WriteString(fmt.Sprintf("Are the squares equal? %t\n\n", (rectangle == anotherRectangle)))

	return shapeString.String()
}
