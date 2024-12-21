package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// VerifyScreen данные для ввода одноразового пароля
type VerifyScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err string
}

func (m *Model) updateVerifyInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.VerifyScreen.Inputs))

	for i := range m.VerifyScreen.Inputs {
		m.VerifyScreen.Inputs[i], cmds[i] = m.VerifyScreen.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// updateVerifyScreen обновляет поля ввода экрана Verify
func updateVerifyScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case verifyMsg:
		{
			if msg.valid {
				m.Authenticated = true
				m.CurrentScreen = "main"
			} else {
				m.VerifyScreen.err = "некорректный одноразовый пароль"
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.VerifyScreen.FocusIndex == len(m.VerifyScreen.Inputs) {
				//cmdNew := cmdWithArg(m.VerifyScreen.Inputs)

				m.User.OTP = m.VerifyScreen.Inputs[0].Value()

				return m, m.verifyOTP
			}

			if s == "up" || s == "shift+tab" {
				m.VerifyScreen.FocusIndex--
			} else {
				m.VerifyScreen.FocusIndex++
			}

			if m.VerifyScreen.FocusIndex > len(m.VerifyScreen.Inputs) {
				m.VerifyScreen.FocusIndex = 0
			} else if m.VerifyScreen.FocusIndex < 0 {
				m.VerifyScreen.FocusIndex = len(m.VerifyScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.VerifyScreen.Inputs))
			for i := 0; i <= len(m.VerifyScreen.Inputs)-1; i++ {
				if i == m.VerifyScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.VerifyScreen.Inputs[i].Focus()
					m.VerifyScreen.Inputs[i].PromptStyle = focusedStyle
					m.VerifyScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.VerifyScreen.Inputs[i].Blur()
				m.VerifyScreen.Inputs[i].PromptStyle = noStyle
				m.VerifyScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateVerifyInputs(msg)
	return *m, cmd
}

// initVerifyScreen инцицилизирует форму для регистрации по умолчанию
func initVerifyScreen() VerifyScreen {
	m := VerifyScreen{
		Inputs: make([]textinput.Model, 1),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		switch i {
		case 0:
			t.Placeholder = "OTP"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// verifyView - второе представление: либо для регистрации, либо для авторизации
func verifyView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Для логина %s на адрес электронной почты %s выслан код\n\n",
		keywordStyle.Render(m.User.Login),
		keywordStyle.Render(m.User.Email))

	msg += "Введите его:"

	for i := range m.VerifyScreen.Inputs {
		b.WriteString(m.VerifyScreen.Inputs[i].View())
		if i < len(m.VerifyScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.VerifyScreen.FocusIndex == len(m.VerifyScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	return m.header() + msg + footer()
}
