package cmd

import (
	"github.com/codeaspiras/gotravel/cmd/views"
	"github.com/codeaspiras/gotravel/stdio"
	"github.com/codeaspiras/gotravel/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func CLI(io stdio.IO) {
	defer utils.PanicHandler(io)
	p := tea.NewProgram(
		views.NewFormModel(),
		tea.WithAltScreen(),
	)

	p.SetWindowTitle("#github.com/codeaspiras/gotravel")

	if _, err := p.Run(); err != nil {
		io.Echo("# Isso gerou um erro:\n- %s", err)
		io.End()
	}
}
