package cli

import (
	"context"

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

	msg.cards = make([]core.Card, 0)

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
func (m *Model) listBinary() tea.Msg {
	ctx := context.TODO()

	bs, err := client.ReadAllBinaries(ctx, m.Token, m.sc)

	msg := binaryListMsg{}
	if err != nil {
		msg.err = err
		return msg
	}

	msg.bins = make([]core.Binary, 0)

	for _, v := range bs {
		msg.bins = append(msg.bins,
			core.Binary{
				Value: v.Value,
				Name:  v.Name,
				Meta:  v.Meta,
				ID:    int(v.Id),
			})
	}

	return msg
}

func (m *Model) addBinary() tea.Msg {
	ctx := context.TODO()
	msg := binaryAddMsg{}
	msg.err = client.CreateBinary(ctx, m.Token, *m.BinaryAdd.b, m.sc)
	return msg
}

func (m *Model) saveBinary(bin core.Binary, dir string) tea.Cmd {
	return func() tea.Msg {
		ctx := context.TODO()
		err := client.ReadBinary(ctx, m.Token, &bin, m.sc)
		if err != nil {
			return saveBinaryMsg{err: err}
		}
		err = core.SaveBinaryToFile(bin, dir)
		return saveBinaryMsg{err: err}
	}
}

/*


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



*/
