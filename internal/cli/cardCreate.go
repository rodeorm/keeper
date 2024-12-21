package cli

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/rodeorm/keeper/internal/core"
)

// CardCreateScreen данные для создания записи о новой карте
type CardCreateScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode

	crd core.Card
	err error // Ошибка при создании карты
}
