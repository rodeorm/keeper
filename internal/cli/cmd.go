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
	ctx := context.TODO()
	msg := cardCreateMsg{}
	msg.err = client.CreateCard(ctx, m.Token, m.CardCreate.crd, m.sc)
	return msg
}

func (m *Model) listCard() tea.Msg {
	ctx := context.TODO()

	cds, err := client.ReadAllCards(ctx, m.Token, m.sc)
	msg := cardListMsg{}
	if err != nil {
		msg.err = err
		return msg
	}

	msg.cards = make([]core.Card, len(cds))

	for _, v := range cds {
		msg.cards = append(msg.cards,
			core.Card{CardNumber: v.CardNumber,
				OwnerName: v.OwnerName,
				ExpMonth:  int(v.ExpMonth),
				ExpYear:   int(v.ExpYear),
				Meta:      v.Meta,
				ID:        int(v.Id),
				CVC:       int(v.CVC)})
	}

	return msg
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
