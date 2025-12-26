package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Macros    []Macro
	Cursor    int
	Listening bool
	List      list.Model
	Editing   bool
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Proper")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			selectedMacroFromList, ok := m.List.SelectedItem().(Item)
			if !ok {
				return m, tea.Quit
			}

			selectedMacro := fmt.Sprint(selectedMacroFromList)
			selectedMacro = strings.ReplaceAll(selectedMacro, " ", "")

			macro, index, err := m.GetMacroAndIndexInListByName(selectedMacro)
			if err != nil {
				fmt.Print(err)
				return m, nil
			}

			m.List.RemoveItem(index)
			if macro.Activated {
				m.List.InsertItem(index, Item(" "+macro.Name))
			} else {
				m.List.InsertItem(index, Item(strings.ReplaceAll(macro.Name, " ", "")))
			}

			macro.toggleActivated()
		case "tab":
			if m.Listening {
				m.Listening = false
			} else {
				m.Listening = true
			}
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.Listening == true {
		return "\n" + "Listening..."
	}
	return "\n" + m.List.View()
}

func (m Model) GetMacroAndIndexInListByName(name string) (*Macro, int, error) {
	name = strings.TrimSpace(strings.ReplaceAll(name, " ", ""))
	items := m.RemovedAllXInNames()

	for i, macro := range items {
		if name == strings.TrimSpace(macro) {
			return &m.Macros[i], i, nil
		}
	}
	return &Macro{}, 0, errors.New("name doesn't match to any macro")
}

func (m Model) RemovedAllXInNames() []string {
	var items []string
	for _, macro := range m.Macros {
		items = append(items, fmt.Sprint(macro.Name))
	}
	return items
}
