package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// WaitScreen данные для ввода одноразового пароля
type WaitScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

func (m *Model) updateWaitInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.WaitScreen.Inputs))

	for i := range m.WaitScreen.Inputs {
		m.WaitScreen.Inputs[i], cmds[i] = m.WaitScreen.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// updateWaitScreen обновляет поля ввода экрана wait
func updateWaitScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case CheckOTPMsg:
		{
			if msg.Valid {
				m.Authenticated = true
				m.CurrentScreen = "main"
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.WaitScreen.FocusIndex == len(m.WaitScreen.Inputs) {
				//cmdNew := cmdWithArg(m.WaitScreen.Inputs)

				m.User.OTP = m.WaitScreen.Inputs[0].Value()

				return m, m.VerifiyOTP
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.WaitScreen.FocusIndex--
			} else {
				m.WaitScreen.FocusIndex++
			}

			if m.WaitScreen.FocusIndex > len(m.WaitScreen.Inputs) {
				m.WaitScreen.FocusIndex = 0
			} else if m.WaitScreen.FocusIndex < 0 {
				m.WaitScreen.FocusIndex = len(m.WaitScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.WaitScreen.Inputs))
			for i := 0; i <= len(m.WaitScreen.Inputs)-1; i++ {
				if i == m.WaitScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.WaitScreen.Inputs[i].Focus()
					m.WaitScreen.Inputs[i].PromptStyle = focusedStyle
					m.WaitScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.WaitScreen.Inputs[i].Blur()
				m.WaitScreen.Inputs[i].PromptStyle = noStyle
				m.WaitScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateWaitInputs(msg)
	return *m, cmd
}

// initWaitScreen инцицилизирует форму для регистрации по умолчанию
func initWaitScreen() WaitScreen {
	m := WaitScreen{
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

// waitView - второе представление: либо для регистрации, либо для авторизации
func waitView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Для логина %s на адрес электронной почты %s выслан код\n\n",
		keywordStyle.Render(m.User.Login),
		keywordStyle.Render(m.User.Email))

	msg += "Введите его:"

	for i := range m.WaitScreen.Inputs {
		b.WriteString(m.WaitScreen.Inputs[i].View())
		if i < len(m.WaitScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.WaitScreen.FocusIndex == len(m.WaitScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите esc или ctrl + c")
	return msg
}
