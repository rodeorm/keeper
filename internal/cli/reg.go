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

// Reg данные регистрации
type Reg struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err string
}

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

func (m *Model) updateRegInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Reg.Inputs))

	for i := range m.Reg.Inputs {
		m.Reg.Inputs[i], cmds[i] = m.Reg.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateReg(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case regMsg:
		{
			if msg.reg {
				m.CurrentScreen = "verify"
				m.OTPMessageSended = true
			} else {
				m.Reg.err = msg.err.Error()
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.Reg.FocusIndex == len(m.Reg.Inputs) {

				m.User = core.User{Login: m.Reg.Inputs[0].Value(),
					Password: m.Reg.Inputs[1].Value(),
					Email:    m.Reg.Inputs[2].Value(),
					Name:     m.Reg.Inputs[3].Value(),
				}

				return m, m.regUser
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.Reg.FocusIndex--
			} else {
				m.Reg.FocusIndex++
			}

			if m.Reg.FocusIndex > len(m.Reg.Inputs) {
				m.Reg.FocusIndex = 0
			} else if m.Reg.FocusIndex < 0 {
				m.Reg.FocusIndex = len(m.Reg.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Reg.Inputs))
			for i := 0; i <= len(m.Reg.Inputs)-1; i++ {
				if i == m.Reg.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.Reg.Inputs[i].Focus()
					m.Reg.Inputs[i].PromptStyle = focusedStyle
					m.Reg.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.Reg.Inputs[i].Blur()
				m.Reg.Inputs[i].PromptStyle = noStyle
				m.Reg.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateRegInputs(msg)
	return *m, cmd
}

// initReg инцицилизирует форму для регистрации по умолчанию
func initReg() Reg {
	m := Reg{
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

// viewReg - форма для регистрации
func viewReg(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите зарегистрироваться?\n\nОтлично! Просто введите  %s, %s, %s и %s \n",
		keywordStyle.Render("логин"),
		keywordStyle.Render("пароль"),
		keywordStyle.Render("адрес электронной почты"),
		keywordStyle.Render("имя"))

	for i := range m.Reg.Inputs {
		b.WriteString(m.Reg.Inputs[i].View())
		if i < len(m.Reg.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.Reg.FocusIndex == len(m.Reg.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	msg += "\n" + b.String()

	if m.Reg.err != "" {
		return msg + err(m.Reg.err) + footer()
	}

	return m.header() + msg + footer()
}
