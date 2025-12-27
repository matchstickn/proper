package components

import (
	"github.com/charmbracelet/lipgloss"
)

func EditingStyle() lipgloss.Style {
	e := lipgloss.NewStyle().
		Border(lipgloss.ASCIIBorder()).
		Background(lipgloss.Color("8")).
		Foreground(lipgloss.Color("9")).
		Padding(2).
		PaddingLeft(2).
		Margin(1)
	return e
}
