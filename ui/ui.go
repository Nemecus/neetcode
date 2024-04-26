package ui

import (
	"strings"

	"github.com/Nemecus/neetcode/ui/components/adapter"
	"github.com/Nemecus/neetcode/ui/components/builder"
	"github.com/Nemecus/neetcode/ui/components/decorator"
	"github.com/Nemecus/neetcode/ui/components/dynamicarray"
	"github.com/Nemecus/neetcode/ui/components/facade"
	"github.com/Nemecus/neetcode/ui/components/factorymethod"
	"github.com/Nemecus/neetcode/ui/components/insertionsort"
	"github.com/Nemecus/neetcode/ui/components/menu"
	"github.com/Nemecus/neetcode/ui/components/observer"
	"github.com/Nemecus/neetcode/ui/components/prototype"
	"github.com/Nemecus/neetcode/ui/components/queue"
	"github.com/Nemecus/neetcode/ui/components/singleton"
	"github.com/Nemecus/neetcode/ui/components/singlylinkedlist"
	"github.com/Nemecus/neetcode/ui/components/state"
	"github.com/Nemecus/neetcode/ui/components/strategy"
	"github.com/Nemecus/neetcode/ui/context"
	"github.com/Nemecus/neetcode/ui/keys"
	"github.com/Nemecus/neetcode/ui/styles"
	"github.com/Nemecus/neetcode/ui/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type mainContentModel struct {
	menu             menu.Model
	dynamicArray     dynamicarray.Model
	singlyLinkedList singlylinkedlist.Model
	insertionSort    insertionsort.Model
	factoryMethod    factorymethod.Model
	singleton        singleton.Model
	queue            queue.Model
	builder          builder.Model
	prototype        prototype.Model
	adapter          adapter.Model
	decorator        decorator.Model
	facade           facade.Model
	strategy         strategy.Model
	observer         observer.Model
	state            state.Model
}

type Model struct {
	ctx  context.ProgramContext
	keys keys.AppKeyMap
	mainContentModel
}

func NewModel() Model {
	m := Model{
		keys: keys.Keys,
	}

	m.ctx = context.ProgramContext{}
	m.ctx.Theme = theme.ParseTheme()
	m.ctx.Styles = styles.InitStyles(m.ctx.Theme)
	m.ctx.State = context.MenuView
	m.menu = menu.NewModel(m.ctx)

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd                 tea.Cmd
		dynamicArrayCmd     tea.Cmd
		singlyLinkedListCmd tea.Cmd
		insertionSortCmd    tea.Cmd
		factoryMethodCmd    tea.Cmd
		singletonCmd        tea.Cmd
		queueCmd            tea.Cmd
		menuCmd             tea.Cmd
		builderCmd          tea.Cmd
		prototypeCmd        tea.Cmd
		adapterCmd          tea.Cmd
		decoratorCmd        tea.Cmd
		facadeCmd           tea.Cmd
		strategyCmd         tea.Cmd
		observerCmd         tea.Cmd
		stateCmd            tea.Cmd
		cmds                []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.ctx.Error = nil

		switch {
		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case menu.AnswerMsg:
		switch msg {
		case "Dynamic Array":
			m.ctx.State = context.DynamicArrayView
			m.dynamicArray = dynamicarray.NewModel(m.ctx)
		case "Singled Linked List":
			m.ctx.State = context.SinglyLinkedListView
			m.singlyLinkedList = singlylinkedlist.NewModel(m.ctx)
		case "Insertion Sort":
			m.ctx.State = context.InsertionSortView
			m.insertionSort = insertionsort.NewModel(m.ctx)
		case "Factory":
			m.ctx.State = context.FactoryView
			m.factoryMethod = factorymethod.NewModel(m.ctx)
		case "Singleton":
			m.ctx.State = context.SingletonView
			m.singleton = singleton.NewModel(m.ctx)
		case "Queue":
			m.ctx.State = context.QueueView
			m.queue = queue.NewModel(m.ctx)
		case "Builder":
			m.ctx.State = context.BuilderView
			m.builder = builder.NewModel(m.ctx)
		case "Prototype":
			m.ctx.State = context.PrototypeView
			m.prototype = prototype.NewModel(m.ctx)
		case "Adapter":
			m.ctx.State = context.AdapterView
			m.adapter = adapter.NewModel(m.ctx)
		case "Decorator":
			m.ctx.State = context.DecoratorView
			m.decorator = decorator.NewModel(m.ctx)
		case "Facade":
			m.ctx.State = context.FacadeView
			m.facade = facade.NewModel(m.ctx)
		case "Strategy":
			m.ctx.State = context.StrategyView
			m.strategy = strategy.NewModel(m.ctx)
		case "Observer":
			m.ctx.State = context.ObserverView
			m.observer = observer.NewModel(m.ctx)
		case "State":
			m.ctx.State = context.StateView
			m.state = state.NewModel(m.ctx)
		}
	case dynamicarray.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case singlylinkedlist.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case insertionsort.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case factorymethod.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case singleton.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case queue.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case builder.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case prototype.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case adapter.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case decorator.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case facade.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case strategy.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case observer.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	case state.AnswerMsg:
		switch msg {
		case "Return":
			m.ctx.State = context.MenuView
		}
	}

	m.syncProgramContext()

	switch m.ctx.State {
	case context.MenuView:
		m.menu, menuCmd = m.menu.Update(msg)
	case context.DynamicArrayView:
		m.dynamicArray, dynamicArrayCmd = m.dynamicArray.Update(msg)
	case context.SinglyLinkedListView:
		m.singlyLinkedList, singlyLinkedListCmd = m.singlyLinkedList.Update(msg)
	case context.InsertionSortView:
		m.insertionSort, insertionSortCmd = m.insertionSort.Update(msg)
	case context.FactoryView:
		m.factoryMethod, factoryMethodCmd = m.factoryMethod.Update(msg)
	case context.SingletonView:
		m.singleton, singletonCmd = m.singleton.Update(msg)
	case context.QueueView:
		m.queue, queueCmd = m.queue.Update(msg)
	case context.BuilderView:
		m.builder, builderCmd = m.builder.Update(msg)
	case context.PrototypeView:
		m.prototype, prototypeCmd = m.prototype.Update(msg)
	case context.AdapterView:
		m.adapter, adapterCmd = m.adapter.Update(msg)
	case context.DecoratorView:
		m.decorator, decoratorCmd = m.decorator.Update(msg)
	case context.FacadeView:
		m.facade, facadeCmd = m.facade.Update(msg)
	case context.StrategyView:
		m.strategy, strategyCmd = m.strategy.Update(msg)
	case context.ObserverView:
		m.observer, observerCmd = m.observer.Update(msg)
	case context.StateView:
		m.state, stateCmd = m.state.Update(msg)
	}
	cmds = append(
		cmds,
		cmd,
		menuCmd,
		dynamicArrayCmd,
		singlyLinkedListCmd,
		insertionSortCmd,
		factoryMethodCmd,
		singletonCmd,
		queueCmd,
		builderCmd,
		prototypeCmd,
		adapterCmd,
		decoratorCmd,
		facadeCmd,
		strategyCmd,
		observerCmd,
		stateCmd,
	)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.ctx.Styles.Common.MainTitleStyle.Render("Welcome to NeetCode Problems!"))
	s.WriteString("\n\n")

	switch m.ctx.State {
	case context.MenuView:
		s.WriteString(m.menu.View())
	case context.DynamicArrayView:
		s.WriteString(m.dynamicArray.View())
	case context.SinglyLinkedListView:
		s.WriteString(m.singlyLinkedList.View())
	case context.InsertionSortView:
		s.WriteString(m.insertionSort.View())
	case context.FactoryView:
		s.WriteString(m.factoryMethod.View())
	case context.SingletonView:
		s.WriteString(m.singleton.View())
	case context.QueueView:
		s.WriteString(m.queue.View())
	case context.BuilderView:
		s.WriteString(m.builder.View())
	case context.PrototypeView:
		s.WriteString(m.prototype.View())
	case context.AdapterView:
		s.WriteString(m.adapter.View())
	case context.DecoratorView:
		s.WriteString(m.decorator.View())
	case context.FacadeView:
		s.WriteString(m.facade.View())
	case context.StrategyView:
		s.WriteString(m.strategy.View())
	case context.ObserverView:
		s.WriteString(m.observer.View())
	case context.StateView:
		s.WriteString(m.state.View())
	}

	s.WriteString("\n")
	return m.ctx.Styles.Common.MainStyle.Render(s.String())
}

func (m *Model) syncProgramContext() {
	m.menu.UpdateProgramContext(&m.ctx)
	m.dynamicArray.UpdateProgramContext(&m.ctx)
}
