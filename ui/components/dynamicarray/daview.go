package dynamicarray

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
		m.status = m.runDynamicArray()
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
	} else {
		s = "Starting Dynamic Array!"
	}

	return m.ctx.Styles.Common.MainStyle.Copy().Render(s)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func (m Model) runDynamicArray() string {
	var daString strings.Builder

	// Test 1
	daString.WriteString("Test 1\n")
	example1 := SetupArray(1)
	daString.WriteString(fmt.Sprintf("Current Size: %d\n", example1.GetSize()))
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example1.GetCapacity()))

	// Test 2
	daString.WriteString("Test 2\n")
	example2 := SetupArray(1)
	daString.WriteString("Running Pushback\n")
	example2.Pushback(1)
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example2.GetCapacity()))
	daString.WriteString("Running Pushback\n")
	example2.Pushback(2)
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example2.GetCapacity()))

	// Test 3
	daString.WriteString("Test 3\n")
	example3 := SetupArray(1)
	daString.WriteString(fmt.Sprintf("Current Size: %d\n", example3.GetSize()))
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example3.GetCapacity()))
	daString.WriteString("Running Pushback\n")
	example3.Pushback(1)
	daString.WriteString(fmt.Sprintf("Current Size: %d\n", example3.GetSize()))
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example3.GetCapacity()))
	daString.WriteString("Running Pushback\n")
	example3.Pushback(2)
	daString.WriteString(fmt.Sprintf("Current Size: %d\n", example3.GetSize()))
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example3.GetCapacity()))
	daString.WriteString(fmt.Sprintf("Running get: %d\n", example3.Get(1)))
	daString.WriteString("Running set\n")
	example3.Set(1, 3)
	daString.WriteString(fmt.Sprintf("Running get: %d\n", example3.Get(1)))
	daString.WriteString(fmt.Sprintf("Running popback: %d\n", example3.Popback()))
	daString.WriteString(fmt.Sprintf("Current Size: %d\n", example3.GetSize()))
	daString.WriteString(fmt.Sprintf("Current Capacity: %d\n", example3.GetCapacity()))
	return daString.String()
}
