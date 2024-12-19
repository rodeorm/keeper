package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SomeMsg struct {
}

func CmdWithArg(tm []textinput.Model) tea.Cmd {
	return func() tea.Msg {
		fmt.Println("Длина textInput.Model", len(tm))
		for i, v := range tm {
			fmt.Println(i, v.Value())
		}
		return SomeMsg{}
	}
}
