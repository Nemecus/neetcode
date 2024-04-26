package observer

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
		m.status = m.runObserver()
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

func (m Model) runObserver() string {
	var observerString strings.Builder

	observerString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	observerString.WriteString("\n\n")

	observerString.WriteString("Welcome to the Online Store!!!\n\n")

	observerString.WriteString("New item was added but not currently available: Awesome Gadget\n")
	item1 := newItem("Awesome Gadget", 0)

	observerString.WriteString(fmt.Sprintf("John Doe subscribes to %s\n", item1.name))
	customer1 := &customer{name: "John Doe"}
	item1.subscribe(customer1)

	observerString.WriteString(fmt.Sprintf("John Doe subscribes to %s\n", item1.name))
	customer2 := &customer{name: "Jane Doe"}
	item1.subscribe(customer2)

	observerString.WriteString(fmt.Sprintf("%d stock added to %s\n", 5, item1.name))
	item1.updateStock(5)

	observerString.WriteString(fmt.Sprintf("John Doe unsubscribes from %s\n", item1.name))
	item1.unsubscribe(customer1)

	observerString.WriteString(fmt.Sprintf("No stock available for %s\n", item1.name))
	item1.updateStock(0)

	observerString.WriteString(fmt.Sprintf("%d stock added to %s\n", 3, item1.name))
	item1.updateStock(3)

	observerString.WriteString(fmt.Sprintf("%d stock added to %s\n", 2, item1.name))
	item1.updateStock(2)

	observerString.WriteString(fmt.Sprintf("%s was notified %d times\n", customer1.name, customer1.countNotifications()))
	observerString.WriteString(fmt.Sprintf("%s was notified %d times\n\n", customer2.name, customer2.countNotifications()))

	return observerString.String()
}
