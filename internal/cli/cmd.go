package cli

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// authMsg - сообщение о результате аутентификации по логину и паролю
type authMsg struct {
	auth bool //Прошел аутентификаию по логину и паролю или нет
}

// verifyMsg - сообщение о результате аутентификаци по одноразовому паролю
type verifyMsg struct {
	valid bool
}

// regMsg - сообщение о рзеультате регистрации
type regMsg struct {
	reg bool  // Да - Был зарегистрирован; Нет - не был зарегистрирован
	err error // Причина, по которой мог не зарегистрироваться
}

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

// CmdWithArg - на случай необходимости передавать параметры в cmd
func CmdWithArg(tm []textinput.Model) tea.Cmd {
	return func() tea.Msg {
		return SomeMsg{}
	}
}

type SomeMsg struct{}
