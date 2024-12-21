package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
)

// Model— это состояние клиентского приложения. Она содержит данные, которые отображаются пользователю.
type Model struct {
	User             core.User // Данные текущего пользователя
	CurrentScreen    string    // Текущий экран для отображения
	Token            string    // Токен
	OTPMessageSended bool      // Было послано сообщение с OTP
	Authenticated    bool      // Авторизован пользователь или нет
	Verified         bool      // Подтвержден OTP или нет
	Quitting         bool

	LogoScreen
	RegScreen
	AuthScreen
	VerifyScreen
	MainScreen
	CardCreateScreen
	CardSelectScreen

	sc proto.KeeperServiceClient
}

// InitialModel инициализирует модель со значениями по умолчанию
func InitialModel(sc proto.KeeperServiceClient) Model {
	var m Model
	m.CurrentScreen = "logo"
	m.User = core.User{}
	m.Token = ""
	m.OTPMessageSended = false
	m.Authenticated = false
	m.Verified = false
	m.Quitting = false
	m.LogoScreen = initLogoScreen()
	m.AuthScreen = initAuthScreen()
	m.RegScreen = initRegScreen()
	m.VerifyScreen = initVerifyScreen()
	m.MainScreen = initMainScreen()
	m.sc = sc

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
