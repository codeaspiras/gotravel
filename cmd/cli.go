package cmd

import (
	"gotravel/cmd/cli"
	"gotravel/pkg/stdio"
	"gotravel/pkg/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func CLI(io stdio.IO) {
	defer utils.PanicHandler(io)
	p := tea.NewProgram(
		cli.NewFormModel(),
		tea.WithAltScreen(),
	)

	p.SetWindowTitle("#gotravel")

	if _, err := p.Run(); err != nil {
		utils.ErrorHandler(io, err)
	}
}
