package cli

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

// AuthScreen данные для авторизации
type AuthScreen struct {
	// Для ввода данных с группы textInput
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

// initAuthScreen инициализирует по умолчанию
func initAuthScreen() AuthScreen {
	m := AuthScreen{
		Inputs: make([]textinput.Model, 2),
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
			t.EchoCharacter = '•'
		}

		m.Inputs[i] = t
	}

	return m
}
