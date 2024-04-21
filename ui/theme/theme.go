package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Background      lipgloss.Color
	TitleText       lipgloss.Color
	PrimaryBorder   lipgloss.Color
	PrimaryText     lipgloss.Color
	SelectTitleText lipgloss.Color
	SelectDescText  lipgloss.Color
	SecondaryText   lipgloss.Color
	FaintText       lipgloss.Color
	HelpText        lipgloss.Color
}

type hexColor string

func ParseTheme() Theme {
	_shimHex := func(hex hexColor) lipgloss.Color {
		return lipgloss.Color(string(hex))

	}

	CurrentTheme := &Theme{
		Background:      _shimHex(hexColor("")),
		TitleText:       _shimHex(hexColor("#FB4F14")),
		PrimaryBorder:   _shimHex(hexColor("#FFFFFF")),
		PrimaryText:     _shimHex(hexColor("#FFFFFF")),
		SelectTitleText: _shimHex(hexColor("#399CA1")),
		SelectDescText:  _shimHex(hexColor("#00556C")),
		SecondaryText:   _shimHex(hexColor("#53565A")),
		FaintText:       _shimHex(hexColor("#a09d9d")),
		HelpText:        _shimHex(hexColor("#FB4F14")),
	}
	return *CurrentTheme
}
