package components

import (
	"proper/models"

	"github.com/charmbracelet/bubbles/list"
)

const defaultWidth = 20

func ListComponentAndStyle(items []list.Item) list.Model {
	l := list.New(items, models.ItemDelegate{}, defaultWidth, models.ListHeight)
	l.Title = "Your Macros"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = models.TitleStyle
	l.Styles.PaginationStyle = models.PaginationStyle
	l.Styles.HelpStyle = models.HelpStyle

	return l
}
