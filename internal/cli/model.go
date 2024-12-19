package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// Model— это состояние клиентского приложения. Она содержит данные, которые отображаются пользователю.
type Model struct {
	User             core.User // Данные текущего пользователя
	CurrentScreen    string    // Текущий экран для отображения
	OTPMessageSended bool      // Было послано сообщение с OTP
	LogoScreen
	RegScreen
	AuthScreen
}

// InitialModel инициализирует модель со значениями по умолчанию
func InitialModel() Model {
	var m Model
	m.CurrentScreen = "logo"
	m.LogoScreen = initLogoScreen()
	m.AuthScreen = initAuthScreen()
	m.RegScreen = initRegScreen()
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
