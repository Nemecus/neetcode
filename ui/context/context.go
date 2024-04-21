package context

import (
	"github.com/Nemecus/neetcode/ui/styles"
	"github.com/Nemecus/neetcode/ui/theme"
)

type sessionState int

const (
	MenuView sessionState = iota
	DynamicArrayView
	SinglyLinkedListView
	InsertionSortView
	FactoryView
	SingletonView
	QueueView
	BuilderView
	PrototypeView
	AdapterView
	DecoratorView
)

type ProgramContext struct {
	State  sessionState
	Error  error
	Theme  theme.Theme
	Styles styles.Styles
}
