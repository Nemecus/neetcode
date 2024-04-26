package strategy

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
		m.status = m.runStrategy()
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

func (m Model) runStrategy() string {
	var strategyString strings.Builder

	strategyString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	strategyString.WriteString("\n\n")

	people := initPeople()

	strategyString.WriteString("Adding Person Doe, Age 20, not married.\n")
	people.addPerson(Person{lastName: "Doe", age: 20, married: false})

	strategyString.WriteString("Adding Person Smith, Age 30, is married.\n")
	people.addPerson(Person{lastName: "Doe", age: 30, married: true})

	strategyString.WriteString("Adding Person Old, Age 70, is married.\n")
	people.addPerson(Person{lastName: "Old", age: 70, married: true})

	peopleCounter := &PeopleCounter{}
	adult := &AdultFilter{}
	peopleCounter.setFilter(adult)
	strategyString.WriteString(fmt.Sprintf(
		"Count of Adults: %d\n",
		peopleCounter.count(*people)))

	senior := &SeniorFilter{}
	peopleCounter.setFilter(senior)
	strategyString.WriteString(fmt.Sprintf(
		"Count of Seniors: %d\n",
		peopleCounter.count(*people)))

	married := &MarriedFilter{}
	peopleCounter.setFilter(married)
	strategyString.WriteString(fmt.Sprintf(
		"Count of Married: %d\n\n",
		peopleCounter.count(*people)))

	return strategyString.String()
}
