package tui

import (
	"time"

	"profile/internal/data"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Model represents the application state
type Model struct {
	Sections            []string
	ActiveSection       int
	ActiveProject       int
	ActiveSkillLog      int
	Filtering           bool
	FilterText          string
	Content             map[string]string
	Projects            []data.Project
	FilteredProjects    []data.Project
	SkillCategories     []data.SkillCategory
	SkillAcquisitionLog []data.LogEntry
	OSSContributionLog  []data.LogEntry
	Dependencies        map[string][]string
	ContactName         string
	ContactTitle        string
	ContactStatus       string
	ContactLinks        []data.ContactLink
	ActiveContactLink   int
	ActiveCategory      int
	ActiveSkillIndex    int
	SkillGridFocus      bool
	Loading             bool
	Spinner             spinner.Model
	SignModal           bool
	SignName            string
	SignDescription     string
	SignField           int
	SignSuccess         bool
	SignError           string
	Styles              Styles
	Width               int
	Height              int
}

// InitialModel returns the initial state of the application
func InitialModel() Model {
	styles := InitialStyles()
	p := data.GetDefaultPortfolio()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.Accent)

	return Model{
		Sections:            p.Sections,
		Content:             p.Content,
		Projects:            p.Projects,
		FilteredProjects:    p.Projects,
		SkillCategories:     p.SkillCategories,
		SkillAcquisitionLog: p.SkillAcquisitionLog,
		OSSContributionLog:  p.OSSContributionLog,
		Dependencies:        p.Dependencies,
		ContactName:         p.ContactName,
		ContactTitle:        p.ContactTitle,
		ContactStatus:       p.ContactStatus,
		ContactLinks:        p.ContactLinks,
		Loading:             true,
		Spinner:             s,
		Styles:              styles,
	}
}

// LoadedMsg is sent when the application has finished loading
type LoadedMsg struct{}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		func() tea.Msg {
			time.Sleep(2 * time.Second)
			return LoadedMsg{}
		},
	)
}
