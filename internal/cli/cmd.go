package cli

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

type SomeMsg struct {
	Text string
}

type CheckOTPMsg struct {
	Valid bool
}

func CmdWithArg(tm []textinput.Model) tea.Cmd {
	return func() tea.Msg {
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

func (m *Model) VerifiyOTP() tea.Msg {
	ctx := context.TODO()
	vrd := CheckOTPMsg{}
	vrd.Valid, m.Token = client.Verify(&m.User, ctx, m.sc)
	return vrd
}
