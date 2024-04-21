package factorymethod

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
		m.status = m.runFactoryMethod()
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

func (m Model) runFactoryMethod() string {
	var fmString strings.Builder

	fmString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Testing Vehicles"))
	fmString.WriteString("\n\n")
	carFactory := &CarFactory{}
	car := carFactory.CreateVehicle()
	fmString.WriteString(fmt.Sprintf("Vehicle is a %s\n", car.GetType()))

	truckFactory := &TruckFactory{}
	truck := truckFactory.CreateVehicle()
	fmString.WriteString(fmt.Sprintf("Vehicle is a %s\n", truck.GetType()))

	bikeFactory := &BikeFactory{}
	bike := bikeFactory.CreateVehicle()
	fmString.WriteString(fmt.Sprintf("Vehicle is a %s\n\n", bike.GetType()))
	return fmString.String()
}
