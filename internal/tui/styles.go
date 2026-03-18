package tui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Styles contains all the UI styles for the application
type Styles struct {
	Container          lipgloss.Style
	Tab                lipgloss.Style
	ActiveTab          lipgloss.Style
	TabShortcut        lipgloss.Style
	TabLabel           lipgloss.Style
	TabsContainer      lipgloss.Style
	ContentArea        lipgloss.Style
	Header             lipgloss.Style
	Footer             lipgloss.Style
	KeyHint            lipgloss.Style
	Shortcut           lipgloss.Style
	Spinner            lipgloss.Style
	ProjectTitle       lipgloss.Style
	ProjectDescription lipgloss.Style
	DescriptionText    lipgloss.Style
	ProjectStack       lipgloss.Style
	ProjectListItem    lipgloss.Style
	ProjectListActive  lipgloss.Style
	ProjectListPane    lipgloss.Style
	ProjectDetailPane  lipgloss.Style
	ProjectBold        lipgloss.Style
	FilterPrompt       lipgloss.Style
	FilterInput        lipgloss.Style
	BentoBox           lipgloss.Style
	BentoTitle         lipgloss.Style
	BentoSkill         lipgloss.Style
	BentoSkillActive   lipgloss.Style
	DependencyBox      lipgloss.Style
	DependencyLine     lipgloss.Style
	LogDate            lipgloss.Style
	LogTitle           lipgloss.Style
	LogDesc            lipgloss.Style
	SectionHeader      lipgloss.Style
	BusinessCard       lipgloss.Style
	BusinessCardInner  lipgloss.Style
	CardName           lipgloss.Style
	CardTitle          lipgloss.Style
	CardRow            lipgloss.Style
	CardLabel          lipgloss.Style
	CardValue          lipgloss.Style
	CardStatus         lipgloss.Style
	Accent             color.Color
	Dim                color.Color
}

// InitialStyles returns the default styles for the application
func InitialStyles() Styles {
	accent := lipgloss.Color("252") // Off-white
	dim := lipgloss.Color("239")    // Much darker grey for better contrast with bold
	border := lipgloss.Color("237") // Dark grey border

	return Styles{
		Accent: accent,
		Dim:    dim,
		Container: lipgloss.NewStyle().
			Padding(1, 1),
		Tab: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(border).
			Padding(0, 3).
			Foreground(dim),
		ActiveTab: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(border).
			Padding(0, 3).
			Foreground(accent).
			Bold(true),
		TabShortcut: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		TabLabel: lipgloss.NewStyle().
			Foreground(dim),
		TabsContainer: lipgloss.NewStyle().
			Align(lipgloss.Center),
		ContentArea: lipgloss.NewStyle().
			Width(76),
		Header: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginBottom(1),
		Footer: lipgloss.NewStyle().
			Foreground(dim).
			Align(lipgloss.Center),
		KeyHint: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginRight(1),
		Shortcut: lipgloss.NewStyle().
			Foreground(dim).
			MarginRight(3),
		Spinner: lipgloss.NewStyle().
			Foreground(accent),
		ProjectTitle: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginBottom(1),
		ProjectDescription: lipgloss.NewStyle().
			Width(50).
			MarginBottom(1),
		DescriptionText: lipgloss.NewStyle().
			Foreground(dim),
		ProjectStack: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			Padding(0, 1).
			Background(lipgloss.Color("236")),
		ProjectListItem: lipgloss.NewStyle().
			Foreground(dim).
			PaddingLeft(2),
		ProjectListActive: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			Background(lipgloss.Color("237")).
			PaddingLeft(2),
		ProjectListPane: lipgloss.NewStyle().
			Width(25).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(border),
		ProjectDetailPane: lipgloss.NewStyle().
			PaddingLeft(4).
			Width(50),
		ProjectBold: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		FilterPrompt: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginRight(1).
			SetString("/"),
		FilterInput: lipgloss.NewStyle().
			Foreground(accent).
			Background(lipgloss.Color("236")).
			Padding(0, 1),
		BentoBox: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(border).
			Padding(0, 1).
			MarginRight(1).
			Width(18),
		BentoTitle: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginBottom(1),
		BentoSkill: lipgloss.NewStyle().
			Foreground(dim),
		BentoSkillActive: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			Background(lipgloss.Color("237")),
		DependencyBox: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(accent).
			Padding(0, 1),
		DependencyLine: lipgloss.NewStyle().
			Foreground(dim),
		LogDate: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			Width(10),
		LogTitle: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		LogDesc: lipgloss.NewStyle().
			Foreground(dim),
		SectionHeader: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginTop(1).
			MarginBottom(1).
			SetString("── "),
		BusinessCard: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(border).
			Padding(2, 4),
		BusinessCardInner: lipgloss.NewStyle().
			PaddingLeft(2).
			Align(lipgloss.Center),
		CardName: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginBottom(1),
		CardTitle: lipgloss.NewStyle().
			Foreground(dim).
			Italic(true).
			MarginBottom(2),
		CardRow: lipgloss.NewStyle().
			MarginBottom(1),
		CardLabel: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		CardValue: lipgloss.NewStyle().
			Foreground(dim),
		CardStatus: lipgloss.NewStyle().
			Foreground(accent).
			Background(lipgloss.Color("237")).
			Padding(0, 1),
	}
}
