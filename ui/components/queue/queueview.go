package queue

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
		m.status = m.runQueue()
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

func (m Model) runQueue() string {
	var queueString strings.Builder

	queueString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test Queue"))
	queueString.WriteString("\n\n")

	newQueue := createQueue()
	queueString.WriteString("Appending 10 to the back of the queue\n")
	newQueue.append(10)

	queueString.WriteString(fmt.Sprintf("Is the queue empty? %t\n", newQueue.isEmpty()))

	queueString.WriteString("Appending 20 to the front of the queue\n")
	newQueue.appendleft(20)

	queueString.WriteString(fmt.Sprintf("Popping from the front of the queue: %d\n", newQueue.popleft()))
	queueString.WriteString(fmt.Sprintf("Popping from the back of the queue: %d\n", newQueue.pop()))
	queueString.WriteString(fmt.Sprintf("Popping from the back of the queue: %d\n", newQueue.pop()))

	queueString.WriteString("Appending 30 to the back of the queue\n")
	newQueue.append(30)

	queueString.WriteString(fmt.Sprintf("Is the queue empty? %t\n\n", newQueue.isEmpty()))

	return queueString.String()
}
