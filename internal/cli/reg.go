package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// RegScreen данные регистрации
type RegScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err string
}

func (m *Model) updateRegInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.RegScreen.Inputs))

	for i := range m.RegScreen.Inputs {
		m.RegScreen.Inputs[i], cmds[i] = m.RegScreen.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateRegScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case regMsg:
		{
			if msg.reg {
				m.CurrentScreen = "verify"
				m.OTPMessageSended = true
			} else {
				m.RegScreen.err = msg.err.Error()
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.RegScreen.FocusIndex == len(m.RegScreen.Inputs) {

				m.User = core.User{Login: m.RegScreen.Inputs[0].Value(),
					Password: m.RegScreen.Inputs[1].Value(),
					Email:    m.RegScreen.Inputs[2].Value(),
					Name:     m.RegScreen.Inputs[3].Value(),
				}

				return m, m.regUser
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.RegScreen.FocusIndex--
			} else {
				m.RegScreen.FocusIndex++
			}

			if m.RegScreen.FocusIndex > len(m.RegScreen.Inputs) {
				m.RegScreen.FocusIndex = 0
			} else if m.RegScreen.FocusIndex < 0 {
				m.RegScreen.FocusIndex = len(m.RegScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.RegScreen.Inputs))
			for i := 0; i <= len(m.RegScreen.Inputs)-1; i++ {
				if i == m.RegScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.RegScreen.Inputs[i].Focus()
					m.RegScreen.Inputs[i].PromptStyle = focusedStyle
					m.RegScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.RegScreen.Inputs[i].Blur()
				m.RegScreen.Inputs[i].PromptStyle = noStyle
				m.RegScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateRegInputs(msg)
	return *m, cmd
}

// initRegScreen инцицилизирует форму для регистрации по умолчанию
func initRegScreen() RegScreen {
	m := RegScreen{
		Inputs: make([]textinput.Model, 4),
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
		case 2:
			t.Placeholder = "Email"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 3:
			t.Placeholder = "Name"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// regView - форма для регистрации
func regView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите зарегистрироваться?\n\nОтлично! Просто введите  %s, %s, %s и %s \n",
		keywordStyle.Render("логин"),
		keywordStyle.Render("пароль"),
		keywordStyle.Render("адрес электронной почты"),
		keywordStyle.Render("имя"))

	for i := range m.RegScreen.Inputs {
		b.WriteString(m.RegScreen.Inputs[i].View())
		if i < len(m.RegScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.RegScreen.FocusIndex == len(m.RegScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	msg += "\n" + b.String()

	if m.RegScreen.err != "" {
		return msg + err(m.RegScreen.err) + footer()
	}

	return m.header() + msg + footer()
}
