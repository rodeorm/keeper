package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// CardCreate данные карты
type CardCreate struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
	crd        core.Card
	err        string
}

func (m *Model) updateCardCreateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.CardCreate.Inputs))

	for i := range m.CardCreate.Inputs {
		m.CardCreate.Inputs[i], cmds[i] = m.CardCreate.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateCardCreate(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cardCreateMsg:
		{
			if msg.err != nil {
				m.CardCreate.err = msg.err.Error()
			}
			m.CurrentScreen = "cardList"
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.CardCreate.FocusIndex == len(m.CardCreate.Inputs) {

				expMonth, _ := strconv.Atoi(m.CardCreate.Inputs[2].Value()) // Валидатор должен был отловить до этого момента, поэтому ошибки не ожидаем
				expYear, _ := strconv.Atoi(m.CardCreate.Inputs[3].Value())
				cvc, _ := strconv.Atoi(m.CardCreate.Inputs[4].Value())
				m.CardCreate.crd = core.Card{CardNumber: m.CardCreate.Inputs[0].Value(),
					OwnerName: m.CardCreate.Inputs[1].Value(),
					ExpMonth:  expMonth,
					ExpYear:   expYear,
					CVC:       cvc,
					Meta:      m.CardCreate.Inputs[5].Value(),
				}

				return m, m.createCard
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.CardCreate.FocusIndex--
			} else {
				m.CardCreate.FocusIndex++
			}

			if m.CardCreate.FocusIndex > len(m.CardCreate.Inputs) {
				m.CardCreate.FocusIndex = 0
			} else if m.CardCreate.FocusIndex < 0 {
				m.CardCreate.FocusIndex = len(m.CardCreate.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.CardCreate.Inputs))
			for i := 0; i <= len(m.CardCreate.Inputs)-1; i++ {
				if i == m.CardCreate.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.CardCreate.Inputs[i].Focus()
					m.CardCreate.Inputs[i].PromptStyle = focusedStyle
					m.CardCreate.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.CardCreate.Inputs[i].Blur()
				m.CardCreate.Inputs[i].PromptStyle = noStyle
				m.CardCreate.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateCardCreateInputs(msg)
	return *m, cmd
}

// initCardCreate инцицилизирует форму для регистрации по умолчанию
func initCardCreate() CardCreate {
	m := CardCreate{
		Inputs: make([]textinput.Model, 6),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Validate = ccnValidator
			t.Placeholder = "XXXX XXXX XXXX XXXX"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Name"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 2:
			t.Placeholder = "Exp.month"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 3:
			t.Placeholder = "Exp.year"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 4:
			t.Placeholder = "CVV"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 5:
			t.Placeholder = "Meta"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
}

// viewCardCreate - форма для регистрации
func viewCardCreate(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите добавить данные новой %s?\n\nПросто заполните форму ниже \n",
		keywordStyle.Render("кредитной карты"))

	for i := range m.CardCreate.Inputs {
		b.WriteString(m.CardCreate.Inputs[i].View())
		if i < len(m.CardCreate.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.CardCreate.FocusIndex == len(m.CardCreate.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	msg += "\n" + b.String()

	if m.CardCreate.err != "" {
		return msg + err(m.CardCreate.err) + footer()
	}

	return m.header() + msg + footer()
}
