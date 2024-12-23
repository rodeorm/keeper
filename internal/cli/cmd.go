package cli

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// regUser - команда на регистрацию пользователя
func (m *Model) regUser() tea.Msg {
	ctx := context.TODO()
	err := client.RegUser(&m.User, ctx, m.sc)
	r := regMsg{}
	if err != nil {
		r.err = err
		r.reg = false
		return r
	}
	r.reg = true
	return r
}

// authUser - команда на авторизацию пользователя
func (m *Model) authUser() tea.Msg {
	ctx := context.TODO()
	err := client.AuthUser(&m.User, ctx, m.sc)
	if err != nil {
		return authMsg{auth: false}
	}
	return authMsg{auth: true}
}

// verifyOTP - команда на верификацию одноразового пароля
func (m *Model) verifyOTP() tea.Msg {
	ctx := context.TODO()
	vrd := verifyMsg{}
	vrd.valid, m.Token = client.Verify(&m.User, ctx, m.sc)
	return vrd
}

func (m *Model) createCard() tea.Msg {
	return nil
}

func (m *Model) listCard() tea.Msg {
	cards := make([]core.Card, 0)

	cards = append(cards, core.Card{CardNumber: "1231 2311 3112 2212", OwnerName: "Alexander 1", ExpMonth: 10, ExpYear: 22, Meta: "this is desc 1"})
	cards = append(cards, core.Card{CardNumber: "1231 2312 3112 2212", OwnerName: "Alexander 2", ExpMonth: 10, ExpYear: 22, Meta: "this is desc 2"})
	cards = append(cards, core.Card{CardNumber: "1231 2315 3112 2212", OwnerName: "Alexander 3", ExpMonth: 10, ExpYear: 22, Meta: "this is desc 3"})
	cards = append(cards, core.Card{CardNumber: "1231 2312 3112 2212", OwnerName: "Alexander 4", ExpMonth: 10, ExpYear: 22, Meta: "this is desc 4"})

	return cardListMsg{cards: cards}
}

func (m *Model) createBin() tea.Msg {
	return nil
}

func (m *Model) listBin() tea.Msg {
	return nil
}

func (m *Model) createCouple() tea.Msg {
	return nil
}

func (m *Model) listCouple() tea.Msg {
	return nil
}

func (m *Model) createText() tea.Msg {
	return nil
}

func (m *Model) listText() tea.Msg {
	return nil
}

// CmdWithArg - на случай необходимости передавать параметры в cmd
func CmdWithArg(tm []textinput.Model) tea.Cmd {
	return func() tea.Msg {
		return SomeMsg{}
	}
}

type SomeMsg struct{}
