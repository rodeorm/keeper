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

// BinaryAdd
type BinaryAdd struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	err error

	b *core.Binary
}

type binaryAddMsg struct {
	err error
}

type saveBinaryMsg struct {
	err error
}

func (m *Model) addBinary() tea.Msg {
	ctx := context.TODO()
	msg := binaryAddMsg{}
	msg.err = client.CreateBinary(ctx, m.Token, *m.BinaryAdd.b, m.sc)
	return msg
}

func (m *Model) saveBinary(bin core.Binary, dir string) tea.Cmd {
	return func() tea.Msg {
		ctx := context.TODO()
		err := client.ReadBinary(ctx, m.Token, &bin, m.sc)
		if err != nil {
			return saveBinaryMsg{err: err}
		}
		err = core.SaveBinaryToFile(bin, dir)
		return saveBinaryMsg{err: err}
	}
}

func (m *Model) updateBinaryAddInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.BinaryAdd.Inputs))

	for i := range m.BinaryAdd.Inputs {
		m.BinaryAdd.Inputs[i], cmds[i] = m.BinaryAdd.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func updateBinaryAdd(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case binaryAddMsg:
		if msg.err != nil {
			m.BinaryAdd.err = fmt.Errorf("ошибка при попытке добавить бинарный файл")
		} else {
			m.CurrentScreen = "main"
		}
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.BinaryAdd.FocusIndex == len(m.BinaryAdd.Inputs) {
				bf, err := core.GetBinaryFromFile(m.BinaryPick.selectedFile)
				if err != nil {
					m.BinaryAdd.err = err
				} else {
					m.BinaryAdd.b = bf
					m.BinaryAdd.b.Name = m.BinaryAdd.Inputs[0].Value()
					m.BinaryAdd.b.Meta = m.BinaryAdd.Inputs[1].Value()
					return m, m.addBinary
				}
			}

			if s == "up" || s == "shift+tab" {
				m.BinaryAdd.FocusIndex--
			} else {
				m.BinaryAdd.FocusIndex++
			}

			if m.BinaryAdd.FocusIndex > len(m.BinaryAdd.Inputs) {
				m.BinaryAdd.FocusIndex = 0
			} else if m.BinaryAdd.FocusIndex < 0 {
				m.BinaryAdd.FocusIndex = len(m.BinaryAdd.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.BinaryAdd.Inputs))
			for i := 0; i <= len(m.BinaryAdd.Inputs)-1; i++ {
				if i == m.BinaryAdd.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.BinaryAdd.Inputs[i].Focus()
					m.BinaryAdd.Inputs[i].PromptStyle = focusedStyle
					m.BinaryAdd.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.BinaryAdd.Inputs[i].Blur()
				m.BinaryAdd.Inputs[i].PromptStyle = noStyle
				m.BinaryAdd.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateBinaryAddInputs(msg)
	return *m, cmd
}

// initBinaryAdd инцицилизирует форму для регистрации по умолчанию
func initBinaryAdd() BinaryAdd {
	m := BinaryAdd{
		Inputs: make([]textinput.Model, 2),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Наименование"
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

// viewBinaryAdd - форма для авторизации
func viewBinaryAdd(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Отлично. Вы выбрали файл %s\n\nУкажите под каким %s и какое %s хотите добавить к нему.\n",
		keywordStyle.Render(m.BinaryPick.selectedFile),
		keywordStyle.Render("именем"),
		keywordStyle.Render("мета-описание"))

	for i := range m.BinaryAdd.Inputs {
		b.WriteString(m.BinaryAdd.Inputs[i].View())
		if i < len(m.BinaryAdd.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.BinaryAdd.FocusIndex == len(m.BinaryAdd.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	msg += "\n" + b.String()
	if m.BinaryAdd.err != nil {
		msg += "\n" + m.BinaryAdd.err.Error() + "\n"
	}

	return m.header() + msg + footer()
}
