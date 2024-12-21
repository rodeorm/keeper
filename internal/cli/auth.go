package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// AuthScreen данные для аутентификации
type AuthScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err error // Ошибка при авторизации
}

func (m *Model) updateAuthInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.AuthScreen.Inputs))

	for i := range m.AuthScreen.Inputs {
		m.AuthScreen.Inputs[i], cmds[i] = m.AuthScreen.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// Update loop for the second view after a choice has been made
func updateAuthScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case authMsg:
		if !msg.auth {
			m.AuthScreen.err = fmt.Errorf("неправильный логин или пароль")
		} else {
			m.CurrentScreen = "verify"
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.AuthScreen.FocusIndex == len(m.AuthScreen.Inputs) {
				m.User = core.User{Login: m.AuthScreen.Inputs[0].Value(),
					Password: m.AuthScreen.Inputs[1].Value(),
				}

				return m, m.authUser
			}

			if s == "up" || s == "shift+tab" {
				m.AuthScreen.FocusIndex--
			} else {
				m.AuthScreen.FocusIndex++
			}

			if m.AuthScreen.FocusIndex > len(m.AuthScreen.Inputs) {
				m.AuthScreen.FocusIndex = 0
			} else if m.AuthScreen.FocusIndex < 0 {
				m.AuthScreen.FocusIndex = len(m.AuthScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.AuthScreen.Inputs))
			for i := 0; i <= len(m.AuthScreen.Inputs)-1; i++ {
				if i == m.AuthScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.AuthScreen.Inputs[i].Focus()
					m.AuthScreen.Inputs[i].PromptStyle = focusedStyle
					m.AuthScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.AuthScreen.Inputs[i].Blur()
				m.AuthScreen.Inputs[i].PromptStyle = noStyle
				m.AuthScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateAuthInputs(msg)
	return *m, cmd
}

// initAuthScreen инцицилизирует форму для регистрации по умолчанию
func initAuthScreen() AuthScreen {
	m := AuthScreen{
		Inputs: make([]textinput.Model, 2),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Login"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// authView - форма для авторизации
func authView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите зарегистрироваться?\n\nОтлично! Просто введите  %s, %s и %s.\n",
		keywordStyle.Render("логин"),
		keywordStyle.Render("пароль"),
		keywordStyle.Render("адрес электронной почты"))

	for i := range m.AuthScreen.Inputs {
		b.WriteString(m.AuthScreen.Inputs[i].View())
		if i < len(m.AuthScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.AuthScreen.FocusIndex == len(m.AuthScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите ") + keywordStyle.Render("ctrl+c")
	return msg
}
