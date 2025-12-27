package cmd

import (
	"runtime"
	"time"

	"proper/components"
	"proper/models"

	"github.com/charmbracelet/bubbles/list"
	"github.com/micmonay/keybd_event"
)

func KeyBDInit() keybd_event.KeyBonding {
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	return kb
}

func InitialModel() models.Model {
	return models.Model{
		Macros: []models.Macro{
			{
				Binds: []models.Bind{
					{
						Keybind: "a",
						Output:  []string{"b", "i", "t"},
					},
				},
				Activated: true,
				Name:      "Bits",
			},
			{
				Binds: []models.Bind{
					{
						Keybind: "/",
						Output:  []string{"b", "y", "t", "e"},
					},
				},
				Activated: true,
				Name:      "Bytes",
			},
		},
		Cursor:    0,
		Listening: false,
		Listener:  components.ListeningStyle(),
		Editing:   false,
		Editor:    components.EditingStyle(),
	}
}

func GenerateListItemsFromMacros(macros []models.Macro) []list.Item {
	items := []list.Item{}
	// loop through all provided macros and add them to the items list after converting their name strings into type Item
	for _, macro := range macros {
		items = append(items, models.Item(macro.Name))
	}

	return items
}
