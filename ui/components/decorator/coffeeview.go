package decorator

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
		m.status = m.runCoffeeDecorator()
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

func (m Model) runCoffeeDecorator() string {
	var coffeeString strings.Builder

	coffeeString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Decorator"))
	coffeeString.WriteString("\n\n")

	coffeeString.WriteString("Let's order a simple coffee!\n")
	coffee := &simpleCoffee{}
	coffeeString.WriteString(fmt.Sprintf("Cost: %.2f\n", coffee.getCost()))

	coffeeString.WriteString("Let's add some milk!\n")
	milkCoffee := &milkCoffee{coffee: coffee}
	coffeeString.WriteString(fmt.Sprintf("Cost: %.2f\n", milkCoffee.getCost()))

	coffeeString.WriteString("Let's add some milk!\n")
	sugarCoffee := &sugarCoffee{coffee: milkCoffee}
	coffeeString.WriteString(fmt.Sprintf("Cost: %.2f\n", sugarCoffee.getCost()))

	coffeeString.WriteString("Let's add some milk!\n")
	creamCoffee := &creamCoffee{coffee: sugarCoffee}
	coffeeString.WriteString(fmt.Sprintf("Cost: %.2f\n\n", creamCoffee.getCost()))

	return coffeeString.String()
}
