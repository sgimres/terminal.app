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
		fmt.Printf("Failed to initialize guestbook database: %v", err)
		os.Exit(1)
	}
	defer data.CloseGuestbookDB()

	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
