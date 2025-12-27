package main

import (
	"fmt"
	"os"

	"proper/cmd"
	"proper/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/micmonay/keybd_event"
)

func main() {
	kb := cmd.KeyBDInit()
	kb.SetKeys(keybd_event.VK_A)

	m := cmd.InitialModel()
	items := cmd.GenerateListItemsFromMacros(m.Macros)

	l := components.ListComponentAndStyle(items)
	m.List = l

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
