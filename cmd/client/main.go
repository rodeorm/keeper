package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/cli"
)

func main() {

	initModel := cli.InitialModel(newGRPCClient())
	p := tea.NewProgram(initModel)

	if _, err := p.Run(); err != nil {
		fmt.Println("Ошибка при старте клиентского приложения", err)
	}
}
