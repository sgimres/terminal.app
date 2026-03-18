package main

import (
	"fmt"
	"os"

	"profile/internal/data"
	"profile/internal/tui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if err := data.InitGuestbookDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize guestbook database: %v\n", err)
		os.Exit(1)
	}
	defer data.CloseGuestbookDB()

	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
