package main

import (
	"fmt"
	"os"

	"proper/cmd"
	"proper/models"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/micmonay/keybd_event"
)

func main() {
	kb := cmd.KeyBDInit()
	kb.SetKeys(keybd_event.VK_A)

	items := cmd.GenerateListItemsFromMacros(cmd.InitialModel().Macros)

	const defaultWidth = 20

	l := list.New(items, models.ItemDelegate{}, defaultWidth, models.ListHeight)
	l.Title = "Your Macros"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = models.TitleStyle
	l.Styles.PaginationStyle = models.PaginationStyle
	l.Styles.HelpStyle = models.HelpStyle

	m := cmd.InitialModel()
	m.List = l

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
