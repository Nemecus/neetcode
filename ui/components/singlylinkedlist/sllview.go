package singlylinkedlist

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
		m.status = m.runSinglyLinkedList()
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

func (m Model) runSinglyLinkedList() string {
	var sllString strings.Builder

	sllString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	sllString.WriteString("\n\n")

	var example1 LinkedList
	sllString.WriteString("Inserting Head value 1\n")
	example1.InsertHead(1)
	sllString.WriteString("Inserting Tail value 2\n")
	example1.InsertTail(2)
	sllString.WriteString("Inserting Head value 0\n")
	example1.InsertHead(0)
	sllString.WriteString("Removing 1\n")
	example1.Remove(1)
	sllString.WriteString(fmt.Sprintf("Full Values List: %d\n\n", example1.GetValues()))

	sllString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 2"))
	sllString.WriteString("\n\n")
	var example2 LinkedList
	sllString.WriteString("Inserting Head value 1\n")
	example2.InsertHead(1)
	sllString.WriteString("Inserting Head value 2\n")
	example2.InsertHead(2)
	sllString.WriteString(fmt.Sprintf("Getting value 5: %d\n\n", example2.Get(5)))

	return sllString.String()
}
