package menu

import (
	"fmt"
	"io"
	"strings"

	"github.com/Nemecus/neetcode/ui/context"
	"github.com/Nemecus/neetcode/ui/keys"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 20

var itemStyle lipgloss.Style
var selectedItemStyle lipgloss.Style

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Model struct {
	ctx      *context.ProgramContext
	keys     keys.AppKeyMap
	MenuList list.Model
	Choice   string
}

type AnswerMsg string

func NewModel(ctx context.ProgramContext) Model {
	items := []list.Item{
		item("Dynamic Array"),
		item("Singled Linked List"),
		item("Insertion Sort"),
		item("Factory"),
		item("Singleton"),
		item("Queue"),
		item("Builder"),
		item("Prototype"),
		item("Adapter"),
		item("Decorator"),
		item("Facade"),
		item("Strategy"),
		item("Observer"),
		item("State"),
	}

	const defaultWidth = 20
	itemStyle = ctx.Styles.Common.NormalTitleStyle
	selectedItemStyle = ctx.Styles.Common.SelectTitleStyle

	newMenu := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	newMenu.Title = "Choose Option:"
	newMenu.SetShowStatusBar(false)
	newMenu.SetFilteringEnabled(false)
	newMenu.Styles.Title = ctx.Styles.Common.MainStyle
	newMenu.Styles.PaginationStyle = ctx.Styles.Common.SelectTitleStyle
	newMenu.Styles.HelpStyle = ctx.Styles.Common.HelpPrimaryTextStyle

	return Model{
		ctx:      &ctx,
		keys:     keys.Keys,
		MenuList: newMenu,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			i, ok := m.MenuList.SelectedItem().(item)
			if ok {
				m.Choice = string(i)
				return m, func() tea.Msg { return AnswerMsg(m.Choice) }
			}
		}
	}

	newMenuList, cmd := m.MenuList.Update(msg)
	m.MenuList = newMenuList

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.MenuList.Styles.Title = m.ctx.Styles.Common.MainStyle
	return m.ctx.Styles.Common.MainStyle.Copy().Render(m.MenuList.View())
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
