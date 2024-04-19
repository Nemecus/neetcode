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
	SelectTitleStyle       lipgloss.Style
	SelectDescStyle        lipgloss.Style
	NormalTitleStyle       lipgloss.Style
	NormalDescStyle        lipgloss.Style
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
		Foreground(theme.PrimaryText).
		Padding(0, 0, 0, 4)

	s.SelectTitleStyle = lipgloss.NewStyle().
		BorderStyle(customBorder).
		BorderForeground(theme.SelectTitleText).
		BorderLeft(true).
		Foreground(theme.SelectTitleText).
		Bold(true).
		MarginLeft(2).
		Padding(0, 0, 0, 1)

	s.SelectDescStyle = lipgloss.NewStyle().
		Foreground(theme.SelectDescText).
		Padding(0, 0, 0, 5)

	s.NormalTitleStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Padding(0, 0, 0, 4)

	s.NormalDescStyle = s.NormalTitleStyle.Copy().
		Foreground(theme.SecondaryText).
		Padding(0, 0, 0, 5)

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
