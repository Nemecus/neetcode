package adapter

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
		m.status = m.runAdapter()
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

func (m Model) runAdapter() string {
	var adapterString strings.Builder

	adapterString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Adapter"))
	adapterString.WriteString("\n\n")

	adapterString.WriteString("Test 1\n")
	adapterString.WriteString("Creating Squarehole\n")
	newSquareHole := &squareHole{sideLenght: 5}

	adapterString.WriteString("Creating Square\n")
	newSquare := &square{sideLength: 5}

	adapterString.WriteString(fmt.Sprintf("Can the square fit in the squarehole? %t\n", newSquareHole.canFit(*newSquare)))

	adapterString.WriteString("Creating Circle\n")
	newCircle := &circle{radius: 5}

	newCircleAdapter := &circleAdapter{circleItem: newCircle}
	convertedSquare := newCircleAdapter.circleToSquareAdapter()
	adapterString.WriteString(fmt.Sprintf("Can the square fit in the squarehole? %t\n\n", newSquareHole.canFit(convertedSquare)))

	adapterString.WriteString("Test 2\n")
	adapterString.WriteString("Creating Squarehole\n")
	newSquareHole2 := &squareHole{sideLenght: 5}

	adapterString.WriteString("Creating Square\n")
	newSquare2 := &square{sideLength: 6}

	adapterString.WriteString(fmt.Sprintf("Can the square fit in the squarehole? %t\n", newSquareHole2.canFit(*newSquare2)))

	adapterString.WriteString("Creating Circle\n")
	newCircle2 := &circle{radius: 2}

	newCircleAdapter2 := &circleAdapter{circleItem: newCircle2}
	convertedSquare2 := newCircleAdapter2.circleToSquareAdapter()
	adapterString.WriteString(fmt.Sprintf("Can the square fit in the squarehole? %t\n\n", newSquareHole.canFit(convertedSquare2)))

	return adapterString.String()
}
