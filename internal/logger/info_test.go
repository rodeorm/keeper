package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	Initialize("info")

	tests := []struct {
		place   string
		message string
		err     string
	}{
		{},
		{
			place:   "какая-то функция",
			message: "какое-то сообщение",
			err:     "какой-то текст ошибки",
		},
	}
	for _, tt := range tests {
		t.Run(tt.place, func(t *testing.T) {

			// Используем require.NotPanics для проверки, что функция не паники.
			require.NotPanics(t, func() {
				Info(tt.place, tt.message, tt.err)
			}, "паника")
		})
	}
}
