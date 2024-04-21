package styles

import (
	"github.com/Nemecus/neetcode/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Help struct {
		Text    lipgloss.Style
		KeyText lipgloss.Style
	}
	Common CommonStyles
}

type CommonStyles struct {
	MainStyle              lipgloss.Style
	MainTitleStyle         lipgloss.Style
	HeaderStyle            lipgloss.Style
	TitleStyle             lipgloss.Style
	SelectTitleStyle       lipgloss.Style
	NormalTitleStyle       lipgloss.Style
	HelpBorderStyle        lipgloss.Style
	HelpPrimaryTextStyle   lipgloss.Style
	HelpSeparatorTextStyle lipgloss.Style
	HelpDescTextStyle      lipgloss.Style
}

var customBorder = lipgloss.Border{
	Top:         "*",
	Bottom:      "*",
	Left:        "⇒",
	Right:       "*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

func InitStyles(theme theme.Theme) Styles {
	var s Styles

	s.Common = BuildStyles(theme)

	s.Help.Text = lipgloss.NewStyle().Foreground(theme.FaintText)
	s.Help.KeyText = lipgloss.NewStyle().Foreground(theme.SecondaryText)

	return s
}

func BuildStyles(theme theme.Theme) CommonStyles {
	var s CommonStyles

	s.MainStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText)

	s.MainTitleStyle = lipgloss.NewStyle().
		Background(theme.TitleText).
		Foreground(theme.PrimaryText)

	s.HeaderStyle = lipgloss.NewStyle().
		Foreground(theme.TitleText)

	s.TitleStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		PaddingLeft(2)

	s.SelectTitleStyle = lipgloss.NewStyle().
		Foreground(theme.SelectTitleText).
		Bold(true).
		MarginLeft(2).
		Padding(0, 0, 0, 1)

	s.NormalTitleStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Padding(0, 0, 0, 4)

	s.HelpBorderStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(theme.PrimaryBorder).
		Height(5).
		Width(27)

	s.HelpPrimaryTextStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(theme.HelpText)

	s.HelpSeparatorTextStyle = lipgloss.NewStyle().
		Foreground(theme.FaintText)

	s.HelpDescTextStyle = lipgloss.NewStyle().
		Foreground(theme.SecondaryText)

	return s
}
