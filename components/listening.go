package components

import (
	"github.com/charmbracelet/lipgloss"
)

func ListeningStyle() lipgloss.Style {
	l := lipgloss.NewStyle().
		Background(lipgloss.Color("3")).
		Foreground(lipgloss.Color("4")).
		Align(lipgloss.Center).
		Padding(2).
		PaddingLeft(2).
		Margin(1)
	return l
}
