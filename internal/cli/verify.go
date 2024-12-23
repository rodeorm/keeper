package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Verify данные для ввода одноразового пароля
type Verify struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err string
}

func (m *Model) updateVerifyInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Verify.Inputs))

	for i := range m.Verify.Inputs {
		m.Verify.Inputs[i], cmds[i] = m.Verify.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// updateVerify обновляет поля ввода экрана Verify
func updateVerify(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case verifyMsg:
		{
			if msg.valid {
				m.Authenticated = true
				m.CurrentScreen = "main"
			} else {
				m.Verify.err = "некорректный одноразовый пароль"
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.Verify.FocusIndex == len(m.Verify.Inputs) {
				//cmdNew := cmdWithArg(m.Verify.Inputs)

				m.User.OTP = m.Verify.Inputs[0].Value()

				return m, m.verifyOTP
			}

			if s == "up" || s == "shift+tab" {
				m.Verify.FocusIndex--
			} else {
				m.Verify.FocusIndex++
			}

			if m.Verify.FocusIndex > len(m.Verify.Inputs) {
				m.Verify.FocusIndex = 0
			} else if m.Verify.FocusIndex < 0 {
				m.Verify.FocusIndex = len(m.Verify.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Verify.Inputs))
			for i := 0; i <= len(m.Verify.Inputs)-1; i++ {
				if i == m.Verify.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.Verify.Inputs[i].Focus()
					m.Verify.Inputs[i].PromptStyle = focusedStyle
					m.Verify.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.Verify.Inputs[i].Blur()
				m.Verify.Inputs[i].PromptStyle = noStyle
				m.Verify.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateVerifyInputs(msg)
	return *m, cmd
}

// initVerify инцицилизирует форму для регистрации по умолчанию
func initVerify() Verify {
	m := Verify{
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

func viewVerify(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Для логина %s на адрес электронной почты %s выслан код\n\n",
		keywordStyle.Render(m.User.Login),
		keywordStyle.Render(m.User.Email))

	msg += "Введите его:"

	for i := range m.Verify.Inputs {
		b.WriteString(m.Verify.Inputs[i].View())
		if i < len(m.Verify.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.Verify.FocusIndex == len(m.Verify.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	return m.header() + msg + footer()
}
