package insertionsort

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
		m.status = m.runInsertionSort()
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

func (m Model) runInsertionSort() string {
	var isString strings.Builder

	pairs := map[int]string{5: "apple", 2: "banana", 9: "cherry"}

	isString.WriteString(m.ctx.Styles.Common.HeaderStyle.Render("Test 1"))
	isString.WriteString("\n\n")
	keys := make([]int, 0, len(pairs))

	for k := range pairs {
		keys = append(keys, k)
	}
	isString.WriteString("Unsorted Pairs: \n")
	for _, k := range keys {
		isString.WriteString(fmt.Sprintf("%d %s \n", k, pairs[k]))
	}

	isString.WriteString("\n")

	sortedPairs := InsertionSort(pairs)
	newKeys := make([]int, 0, len(sortedPairs))
	for nk := range sortedPairs {
		newKeys = append(newKeys, nk)
	}
	isString.WriteString("Sorted Pairs: \n")
	for _, nk := range newKeys {
		isString.WriteString(fmt.Sprintf("%d %s \n", nk, sortedPairs[nk]))
	}
	isString.WriteString("\n")

	/*
		 	maps don't allow for duplicate keys

			pairs2 := map[int]string{3: "cat", 3: "bird", 2: "dog"}

			fmt.Println("Test 1")
			keys2 := make([]int, 0, len(pairs2))

			for k := range pairs {
				keys2 = append(keys2, k)
			}
			fmt.Println("Unsorted Pairs: ")
			for _, k := range keys2 {
				fmt.Println(k, pairs[k])
			}

			sortedPairs2 := insertionsort.InsertionSort(pairs2)
			newKeys2 := make([]int, 0, len(sortedPairs2))
			for nk := range sortedPairs2 {
				newKeys2 = append(newKeys2, nk)
			}
			fmt.Println("Sorted Pairs: ")
			for _, nk := range newKeys2 {
				fmt.Println(nk, sortedPairs2[nk])
			}
	*/
	return isString.String()
}
