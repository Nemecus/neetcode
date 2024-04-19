package menu

import (
	"github.com/Nemecus/neetcode/ui/context"
	"github.com/Nemecus/neetcode/ui/keys"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx      *context.ProgramContext
	keys     keys.AppKeyMap
	MenuList list.Model
	Choice   string
}

type item struct {
	title, desc string
}

type AnswerMsg string

func menuDelegate(ctx *context.ProgramContext) list.DefaultDelegate {
	newDelegate := list.NewDefaultDelegate()

	newDelegate.Styles.SelectedTitle = ctx.Styles.Common.SelectTitleStyle.Copy()
	newDelegate.Styles.SelectedDesc = ctx.Styles.Common.SelectDescStyle.Copy()
	newDelegate.Styles.NormalTitle = ctx.Styles.Common.NormalTitleStyle.Copy()
	newDelegate.Styles.NormalDesc = ctx.Styles.Common.NormalDescStyle.Copy()
	return newDelegate
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewModel(ctx context.ProgramContext) Model {
	items := []list.Item{
		item{title: "Dynamic Array", desc: "Data Structure"},
		item{title: "Singled Linked List", desc: "Data Structure"},
		item{title: "Insertion Sort", desc: "Sorting"},
		item{title: "Factory", desc: "Design Pattern"},
		item{title: "Singleton", desc: "Design Pattern"},
		item{title: "Queue", desc: "Data Structure"},
		item{title: "Builder", desc: "Design Pattern"},
		item{title: "Prototype", desc: "Design Pattern"},
		item{title: "Adapter", desc: "Design Pattern"},
		item{title: "Decorator", desc: "Design Pattern"},
	}

	newMenuDelegate := menuDelegate(&ctx)

	newMenu := list.New(items, newMenuDelegate, 0, 0)
	newMenu.Title = "Choose Option:"
	newMenu.SetShowHelp(false)
	newMenu.SetShowStatusBar(false)
	newMenu.SetSize(66, 25)

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
				m.Choice = i.title
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
