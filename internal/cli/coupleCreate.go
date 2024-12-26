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

// CoupleCreate данные пары логин-пароль
type CoupleCreate struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
	couple     core.Couple
	err        string
}

type coupleCreateMsg struct {
	err error
}

func (m *Model) createCouple() tea.Msg {
	ctx := context.TODO()
	msg := coupleCreateMsg{}
	msg.err = client.CreateCouple(ctx, m.Token, m.CoupleCreate.couple, m.sc)
	return msg
}

func (m *Model) updateCoupleCreateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.CoupleCreate.Inputs))

	for i := range m.CoupleCreate.Inputs {
		m.CoupleCreate.Inputs[i], cmds[i] = m.CoupleCreate.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateCoupleCreate(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case coupleCreateMsg:
		{
			if msg.err != nil {
				m.CoupleCreate.err = msg.err.Error()
			}
			m.CurrentScreen = "coupleList"
			return m, m.listCouple
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.CoupleCreate.FocusIndex == len(m.CoupleCreate.Inputs) {
				m.CoupleCreate.couple = core.Couple{Source: m.CoupleCreate.Inputs[0].Value(),
					Login:    m.CoupleCreate.Inputs[1].Value(),
					Password: m.CoupleCreate.Inputs[2].Value(),
					Meta:     m.CoupleCreate.Inputs[3].Value(),
				}

				return m, m.createCouple
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.CoupleCreate.FocusIndex--
			} else {
				m.CoupleCreate.FocusIndex++
			}

			if m.CoupleCreate.FocusIndex > len(m.CoupleCreate.Inputs) {
				m.CoupleCreate.FocusIndex = 0
			} else if m.CoupleCreate.FocusIndex < 0 {
				m.CoupleCreate.FocusIndex = len(m.CoupleCreate.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.CoupleCreate.Inputs))
			for i := 0; i <= len(m.CoupleCreate.Inputs)-1; i++ {
				if i == m.CoupleCreate.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.CoupleCreate.Inputs[i].Focus()
					m.CoupleCreate.Inputs[i].PromptStyle = focusedStyle
					m.CoupleCreate.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.CoupleCreate.Inputs[i].Blur()
				m.CoupleCreate.Inputs[i].PromptStyle = noStyle
				m.CoupleCreate.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateCoupleCreateInputs(msg)
	return *m, cmd
}

// initCoupleCreate
func initCoupleCreate() CoupleCreate {
	m := CoupleCreate{
		Inputs: make([]textinput.Model, 4),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Источник"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Логин"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 2:
			t.Placeholder = "Пароль"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 3:
			t.Placeholder = "Мета"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// viewCoupleCreate
func viewCoupleCreate(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите добавить новые %s?\n\nПросто заполните форму ниже \n",
		keywordStyle.Render("логин-пароль"))

	for i := range m.CoupleCreate.Inputs {
		b.WriteString(m.CoupleCreate.Inputs[i].View())
		if i < len(m.CoupleCreate.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.CoupleCreate.FocusIndex == len(m.CoupleCreate.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	msg += "\n" + b.String()

	if m.CoupleCreate.err != "" {
		return msg + err(m.CoupleCreate.err) + footer()
	}

	return m.header() + msg + footer()
}
