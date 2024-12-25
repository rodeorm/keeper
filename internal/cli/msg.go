package cli

import "github.com/rodeorm/keeper/internal/core"

type binaryAddMsg struct {
	err error
}

type saveBinaryMsg struct {
	err error
}

type binaryListMsg struct {
	bins []core.Binary
	err  error
}

// cardCreateMsg - сообщение о результате аутентификации по логину и паролю
type cardCreateMsg struct {
	err error
}

// cardListMsg - сообщение о результате аутентификации по логину и паролю
type cardListMsg struct {
	cards []core.Card
	err   error
}

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
