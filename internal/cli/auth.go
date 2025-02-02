package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// Auth данные для аутентификации
type Auth struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err error // Ошибка при авторизации
}

type authMsg struct {
	auth bool //Прошел аутентификаию по логину и паролю или нет
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

func (m *Model) updateAuthInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Auth.Inputs))

	for i := range m.Auth.Inputs {
		m.Auth.Inputs[i], cmds[i] = m.Auth.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateAuth(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case authMsg:
		if !msg.auth {
			m.Auth.err = fmt.Errorf("неправильный логин или пароль")
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
			if s == "enter" && m.Auth.FocusIndex == len(m.Auth.Inputs) {
				m.User = core.User{Login: m.Auth.Inputs[0].Value(),
					Password: m.Auth.Inputs[1].Value(),
				}

				return m, m.authUser
			}

			if s == "up" || s == "shift+tab" {
				m.Auth.FocusIndex--
			} else {
				m.Auth.FocusIndex++
			}

			if m.Auth.FocusIndex > len(m.Auth.Inputs) {
				m.Auth.FocusIndex = 0
			} else if m.Auth.FocusIndex < 0 {
				m.Auth.FocusIndex = len(m.Auth.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Auth.Inputs))
			for i := 0; i <= len(m.Auth.Inputs)-1; i++ {
				if i == m.Auth.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.Auth.Inputs[i].Focus()
					m.Auth.Inputs[i].PromptStyle = focusedStyle
					m.Auth.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.Auth.Inputs[i].Blur()
				m.Auth.Inputs[i].PromptStyle = noStyle
				m.Auth.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateAuthInputs(msg)
	return *m, cmd
}

// initAuth инцицилизирует форму для регистрации по умолчанию
func initAuth() Auth {
	m := Auth{
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

// viewAuth - форма для авторизации
func viewAuth(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите авторизоваться?\n\nОтлично! Просто введите  %s и %s.\n",
		keywordStyle.Render("логин"),
		keywordStyle.Render("пароль"))

	for i := range m.Auth.Inputs {
		b.WriteString(m.Auth.Inputs[i].View())
		if i < len(m.Auth.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.Auth.FocusIndex == len(m.Auth.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите ") + keywordStyle.Render("ctrl+c")
	return m.header() + msg + footer()
}
