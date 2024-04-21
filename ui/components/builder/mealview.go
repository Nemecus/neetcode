package builder

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
		m.status = m.runMealBuilder()
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

func (m Model) runMealBuilder() string {
	var mealString strings.Builder

	mealString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Builder"))
	mealString.WriteString("\n\n")

	mealBuilder := getBuilder()
	director := newDirector(mealBuilder)
	mealString.WriteString("Order In! Takeout Order for Burger and Coke!\n\n")
	meal := director.buildMeal(15.99, true, "Burger", "Coke")

	mealString.WriteString(fmt.Sprintf("Cost of meal: %.2f\n", meal.cost))
	mealString.WriteString(fmt.Sprintf("Takeout: %t\n", meal.takeout))
	mealString.WriteString(fmt.Sprintf("Main dish: %s\n", meal.main))
	mealString.WriteString(fmt.Sprintf("Drink: %s\n\n", meal.drink))

	return mealString.String()

}
