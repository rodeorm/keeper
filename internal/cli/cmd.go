package cli

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

type SomeMsg struct {
	Text string
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

func (m *Model) RegUser() tea.Msg {
	ctx := context.TODO()
	err := client.RegUser(&m.User, ctx, m.sc)
	if err != nil {
		return SomeMsg{Text: err.Error()}
	}
	return SomeMsg{}
}
