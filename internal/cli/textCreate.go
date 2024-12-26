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

// TextCreate данные пары логин-пароль
type TextCreate struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
	Text       core.Text
	err        string
}

type TextCreateMsg struct {
	err error
}

func (m *Model) createText() tea.Msg {
	ctx := context.TODO()
	msg := TextCreateMsg{}
	msg.err = client.CreateText(ctx, m.Token, m.TextCreate.Text, m.sc)
	return msg
}

func (m *Model) updateTextCreateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.TextCreate.Inputs))

	for i := range m.TextCreate.Inputs {
		m.TextCreate.Inputs[i], cmds[i] = m.TextCreate.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateTextCreate(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TextCreateMsg:
		{
			if msg.err != nil {
				m.TextCreate.err = msg.err.Error()
			}
			m.CurrentScreen = "textList"
			return m, m.listText
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.TextCreate.FocusIndex == len(m.TextCreate.Inputs) {
				m.TextCreate.Text = core.Text{Value: m.TextCreate.Inputs[0].Value(),
					Meta: m.TextCreate.Inputs[1].Value(),
				}

				return m, m.createText
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.TextCreate.FocusIndex--
			} else {
				m.TextCreate.FocusIndex++
			}

			if m.TextCreate.FocusIndex > len(m.TextCreate.Inputs) {
				m.TextCreate.FocusIndex = 0
			} else if m.TextCreate.FocusIndex < 0 {
				m.TextCreate.FocusIndex = len(m.TextCreate.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.TextCreate.Inputs))
			for i := 0; i <= len(m.TextCreate.Inputs)-1; i++ {
				if i == m.TextCreate.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.TextCreate.Inputs[i].Focus()
					m.TextCreate.Inputs[i].PromptStyle = focusedStyle
					m.TextCreate.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.TextCreate.Inputs[i].Blur()
				m.TextCreate.Inputs[i].PromptStyle = noStyle
				m.TextCreate.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateTextCreateInputs(msg)
	return *m, cmd
}

// initTextCreate
func initTextCreate() TextCreate {
	m := TextCreate{
		Inputs: make([]textinput.Model, 2),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle

		switch i {
		case 0:
			t.Placeholder = "Текст"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Мета"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// viewTextCreate
func viewTextCreate(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите добавить новый %s?\n\nПросто заполните форму ниже \n",
		keywordStyle.Render("текст"))

	for i := range m.TextCreate.Inputs {
		b.WriteString(m.TextCreate.Inputs[i].View())
		if i < len(m.TextCreate.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.TextCreate.FocusIndex == len(m.TextCreate.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	msg += "\n" + b.String()

	if m.TextCreate.err != "" {
		return msg + err(m.TextCreate.err) + footer()
	}

	return m.header() + msg + footer()
}
