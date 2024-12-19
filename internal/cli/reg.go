package cli

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

// RegScreen данные регистрации
type RegScreen struct {
	// Для ввода данных с группы textInput
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
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
