package method

import "github.com/charmbracelet/bubbles/key"

var (
	toggleKey = key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("space", "toggle"),
	)

	selectKey = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	)
)
