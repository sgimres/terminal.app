package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// View renders the UI
func (m Model) View() tea.View {
	if m.Loading {
		loading := fmt.Sprintf("\n  %s Loading Portfolio...", m.Spinner.View())
		if m.Width > 0 && m.Height > 0 {
			return tea.NewView(lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, loading))
		}
		return tea.NewView(loading)
	}

	// Tabs
	var tabs []string
	for i, section := range m.Sections {
		shortcutLetter := strings.ToLower(string(section[0]))
		fullName := strings.ToLower(section)

		var tabContent string
		if m.ActiveSection == i {
			tabContent = m.Styles.ActiveTab.Render(fmt.Sprintf("%s %s", shortcutLetter, fullName))
		} else {
			styledShortcut := m.Styles.TabShortcut.Render(shortcutLetter)
			styledLabel := m.Styles.TabLabel.Render(fullName)
			tabContent = m.Styles.Tab.Render(fmt.Sprintf("%s %s", styledShortcut, styledLabel))
		}
		tabs = append(tabs, tabContent)
	}
	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	tabBarCentered := lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, tabBar)

	// Copy URL Banner
	var copyURLBanner string
	if m.CopyURL != "" {
		copyURLBanner = m.Styles.CopyURLBanner.Render("Copy this URL: " + m.CopyURL)
		copyURLBanner = lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, copyURLBanner)
	}

	// Main Content
	var content string
	switch m.Sections[m.ActiveSection] {
	case "Projects":
		content = m.renderProjects()
	case "Skills":
		content = m.renderSkills()
	case "Contact":
		content = m.renderContact()
	default:
		// body := m.Content[m.Sections[m.ActiveSection]]
		// content = m.Styles.ContentArea.Render(m.parseDescription(body))
		content = m.Styles.ContentArea.Render(m.parseDescription(m.Content["About"]))
	}
	contentCentered := lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, content)

	// Footer
	footerCentered := lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, m.renderFooter())

	// Divider
	divider := lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		Render(strings.Repeat("─", 80))
	dividerCentered := lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, divider)

	secondaryFooter := m.Styles.Footer.Render("Zero AI was harmed in the making of this app • 2026")
	secondaryFooterCentered := lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, secondaryFooter)

	// Combined Layout
	finalView := m.Styles.Container.Render(lipgloss.JoinVertical(
		lipgloss.Center,
		tabBarCentered,
		copyURLBanner,
		contentCentered,
		secondaryFooterCentered,
		dividerCentered,
		footerCentered,
	))

	if m.SignModal {
		modal := m.renderSignModal()
		if m.Width > 0 && m.Height > 0 {
			return tea.NewView(lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal))
		}
		return tea.NewView(modal)
	}

	if m.Width > 0 && m.Height > 0 {
		return tea.NewView(lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, finalView))
	}
	return tea.NewView(finalView)
}

func (m Model) renderSignModal() string {
	if m.SignSuccess {
		content := lipgloss.JoinVertical(lipgloss.Center,
			m.Styles.ModalTitle.Render("Thank You!"),
			"",
			m.Styles.CardValue.Render("Your message has been saved."),
		)
		box := m.Styles.ModalBox.Render(content)
		return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, box)
	}

	nameLabel := m.Styles.CardLabel.Render("Your Name:")
	nameValue := m.Styles.FilterInput.Render(m.SignName)
	if m.SignField == 0 {
		nameValue = m.Styles.FilterInput.Background(lipgloss.Color("237")).Render(m.SignName + "_")
	}

	descLabel := m.Styles.CardLabel.Render("Message:")
	descValue := m.Styles.FilterInput.Render(m.SignDescription)
	if m.SignField == 1 {
		descValue = m.Styles.FilterInput.Background(lipgloss.Color("237")).Render(m.SignDescription + "_")
	}

	footerHint := "tab: next  •  enter: submit  •  esc: cancel"
	if m.SignError != "" {
		footerHint = m.SignError + "  •  press any key to continue"
	}

	content := lipgloss.JoinVertical(lipgloss.Left,
		m.Styles.ModalTitle.Render("Sign Guestbook"),
		"",
		nameLabel,
		nameValue,
		"",
		descLabel,
		descValue,
		"",
		m.Styles.Footer.Render(footerHint),
	)

	box := m.Styles.ModalBox.Render(content)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, box)
}

func (m Model) renderProjects() string {
	var projectsLayout string

	// Filter Bar
	if m.Filtering {
		filterBar := lipgloss.JoinHorizontal(lipgloss.Top,
			m.Styles.FilterPrompt.String(),
			m.Styles.FilterInput.Render(m.FilterText),
		)
		projectsLayout = lipgloss.NewStyle().MarginBottom(1).Render(filterBar) + "\n"
	}

	// Left Pane: Project List
	var projectList []string
	if len(m.FilteredProjects) == 0 {
		projectList = append(projectList, m.Styles.ProjectListItem.Render("no results"))
	} else {
		for i, p := range m.FilteredProjects {
			titleParts := strings.SplitN(p.Title, " ", 2)
			title := titleParts[len(titleParts)-1]

			if m.ActiveProject == i {
				projectList = append(projectList, m.Styles.ProjectListActive.Render(strings.ToLower(title)))
			} else {
				projectList = append(projectList, m.Styles.ProjectListItem.Render(strings.ToLower(title)))
			}
		}
	}
	leftPane := m.Styles.ProjectListPane.Render(lipgloss.JoinVertical(lipgloss.Left, projectList...))

	// Right Pane: Project Details
	var rightPane string
	if len(m.FilteredProjects) > 0 {
		p := m.FilteredProjects[m.ActiveProject]
		title := m.Styles.ProjectTitle.Render(p.Title)
		descRaw := m.parseDescription(p.Description)
		desc := m.Styles.ProjectDescription.Render(descRaw)
		stack := m.Styles.ProjectStack.Render(fmt.Sprintf("stack: %s", strings.ToLower(p.Stack)))

		details := lipgloss.JoinVertical(lipgloss.Left, title, desc, stack)
		if p.URL != "" {
			url := m.Styles.CardValue.Foreground(lipgloss.Color("226")).Render(p.URL)
			details += "\n" + url
		}
		rightPane = m.Styles.ProjectDetailPane.Render(details)
	} else {
		rightPane = m.Styles.ProjectDetailPane.Render(m.Styles.ProjectDescription.Render("try another search..."))
	}

	// Combine Panes
	projectsLayout += lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)
	return m.Styles.ContentArea.Render(projectsLayout)
}

func (m Model) renderSkills() string {
	// 1. Bento Box Grid
	var boxes []string
	for i, cat := range m.SkillCategories {
		var skillList []string
		for j, s := range cat.Skills {
			style := m.Styles.BentoSkill
			if m.SkillGridFocus && m.ActiveCategory == i && m.ActiveSkillIndex == j {
				style = m.Styles.BentoSkillActive
			}
			skillList = append(skillList, style.Render(s))
		}
		title := m.Styles.BentoTitle.Render(strings.ToLower(cat.Name))
		content := lipgloss.JoinVertical(lipgloss.Left, title, lipgloss.JoinVertical(lipgloss.Left, skillList...))
		boxes = append(boxes, m.Styles.BentoBox.Render(content))
	}
	bentoGrid := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)

	// Dependency Graph (if focused)
	var depGraph string
	if m.SkillGridFocus {
		selectedSkill := m.SkillCategories[m.ActiveCategory].Skills[m.ActiveSkillIndex]
		depGraph = m.renderDependencyGraph(selectedSkill)
	}

	// 2. Skill Acquisition Log
	var acqEntries []string
	acqEntries = append(acqEntries, m.Styles.SectionHeader.Render("SKILL ACQUISITION LOG"))
	for _, entry := range m.SkillAcquisitionLog {
		date := m.Styles.LogDate.Render(entry.Date)
		title := m.Styles.LogTitle.Render(strings.ToLower(entry.Title))
		desc := m.Styles.LogDesc.Render(entry.Description)

		row := lipgloss.JoinHorizontal(lipgloss.Top, date, lipgloss.JoinVertical(lipgloss.Left, title, desc))
		acqEntries = append(acqEntries, row)
	}
	acqLog := lipgloss.JoinVertical(lipgloss.Left, acqEntries...)

	// 3. OSS Contribution Log
	var ossLogEntries []string
	ossLogEntries = append(ossLogEntries, m.Styles.SectionHeader.Render("OSS CONTRIBUTION LOG"))
	for i, entry := range m.OSSContributionLog {
		date := m.Styles.LogDate.Render(entry.Date)
		title := m.Styles.LogTitle.Render(strings.ToLower(entry.Title))
		if !m.SkillGridFocus && i == m.ActiveSkillLog {
			title = m.Styles.LogTitle.Background(lipgloss.Color("237")).Render(strings.ToLower(entry.Title))
		}
		desc := m.Styles.LogDesc.Render(entry.Description)

		row := lipgloss.JoinHorizontal(lipgloss.Top, date, lipgloss.JoinVertical(lipgloss.Left, title, desc))
		if !m.SkillGridFocus && i == m.ActiveSkillLog && entry.URL != "" {
			url := m.Styles.LogDesc.Foreground(lipgloss.Color("226")).Render(entry.URL)
			row += " " + url
		}
		ossLogEntries = append(ossLogEntries, row)
	}
	ossLog := lipgloss.JoinVertical(lipgloss.Left, ossLogEntries...)

	mainView := lipgloss.JoinVertical(lipgloss.Left, bentoGrid, depGraph, acqLog, ossLog)
	return m.Styles.ContentArea.Render(mainView)
}

func (m Model) renderDependencyGraph(skill string) string {
	deps, ok := m.Dependencies[skill]
	if !ok || len(deps) == 0 {
		return ""
	}

	// Center main skill box
	mainBox := m.Styles.DependencyBox.Render(strings.ToLower(skill))

	// Create dependency boxes
	var depBoxes []string
	for _, d := range deps {
		depBoxes = append(depBoxes, m.Styles.DependencyBox.Render(strings.ToLower(d)))
	}

	// Lines and joining
	connector := m.Styles.DependencyLine.Render(" ──┬── ")
	if len(deps) == 1 {
		connector = m.Styles.DependencyLine.Render(" ─── ")
	}

	// Simple flow: [ Skill ] ──┬── [ Dep 1 ] [ Dep 2 ]
	row := lipgloss.JoinHorizontal(lipgloss.Center, mainBox, connector, lipgloss.JoinHorizontal(lipgloss.Top, depBoxes...))

	return lipgloss.NewStyle().MarginTop(1).MarginBottom(1).PaddingLeft(2).Render(row)
}

func (m Model) parseDescription(desc string) string {
	var plain strings.Builder
	parts := strings.Split(desc, "**")
	for i, part := range parts {
		if i%2 == 1 {
			plain.WriteString(m.Styles.ProjectBold.Render(part))
		} else {
			plain.WriteString(m.Styles.DescriptionText.Render(part))
		}
	}

	text := plain.String()

	var wrapped strings.Builder
	words := strings.Fields(text)
	lineLen := 0
	const maxWidth = 48

	for _, word := range words {
		stripped := stripANSI(word)
		wordLen := len(stripped)

		if lineLen+wordLen+1 > maxWidth && lineLen > 0 {
			wrapped.WriteString("\n")
			lineLen = 0
		}
		if lineLen > 0 {
			wrapped.WriteString(" ")
			lineLen++
		}
		wrapped.WriteString(word)
		lineLen += wordLen
	}
	return wrapped.String()
}

func stripANSI(s string) string {
	var result strings.Builder
	inEscape := false
	for _, c := range s {
		if c == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape && c == 'm' {
			inEscape = false
			continue
		}
		if !inEscape {
			result.WriteRune(c)
		}
	}
	return result.String()
}

func (m Model) renderFooter() string {
	var footerItems []string

	if m.Filtering {
		footerItems = append(footerItems, m.Styles.KeyHint.Render("esc"))
		footerItems = append(footerItems, m.Styles.Shortcut.Render("clear"))
		footerItems = append(footerItems, m.Styles.KeyHint.Render("enter"))
		footerItems = append(footerItems, m.Styles.Shortcut.Render("done"))
	} else {
		// Contextual Navigation
		switch m.Sections[m.ActiveSection] {
		case "Projects":
			footerItems = append(footerItems, m.Styles.KeyHint.Render("↑/↓"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("browse"))
			footerItems = append(footerItems, m.Styles.KeyHint.Render("/"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("filter"))
			footerItems = append(footerItems, m.Styles.KeyHint.Render("enter"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("copy url"))
		case "Skills":
			footerItems = append(footerItems, m.Styles.KeyHint.Render("tab"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("toggle focus"))

			footerItems = append(footerItems, m.Styles.Shortcut.Render("navigate"))
			if !m.SkillGridFocus {
				footerItems = append(footerItems, m.Styles.KeyHint.Render("enter"))
				footerItems = append(footerItems, m.Styles.Shortcut.Render("copy url"))
			}
		case "Contact":
			footerItems = append(footerItems, m.Styles.KeyHint.Render("↑/↓"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("navigate"))
			footerItems = append(footerItems, m.Styles.KeyHint.Render("enter"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("copy url"))
			footerItems = append(footerItems, m.Styles.KeyHint.Render("g"))
			footerItems = append(footerItems, m.Styles.Shortcut.Render("sign"))
		}

		footerItems = append(footerItems, m.Styles.KeyHint.Render("q"))
		footerItems = append(footerItems, m.Styles.Shortcut.Render("quit"))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, footerItems...)
}

func (m Model) renderContact() string {
	var contactRows []string
	for i, link := range m.ContactLinks {
		var label, value, url string
		if i == m.ActiveContactLink {
			label = m.Styles.CardLabel.Background(lipgloss.Color("237")).Render(link.Label + ":")
			value = m.Styles.CardValue.Background(lipgloss.Color("237")).Render(link.Value)
			if link.URL != "" {
				url = m.Styles.CardValue.Background(lipgloss.Color("237")).Foreground(lipgloss.Color("226")).Render(link.URL)
			}
		} else {
			label = m.Styles.CardLabel.Render(link.Label + ":")
			value = m.Styles.CardValue.Render(link.Value)
			if link.URL != "" {
				url = m.Styles.CardValue.Foreground(lipgloss.Color("246")).Render(link.URL)
			}
		}
		rowContent := m.Styles.CardRow.Render(label + " " + value)
		if url != "" {
			rowContent += " " + url
		}
		contactRows = append(contactRows, rowContent)
	}

	header := lipgloss.JoinVertical(lipgloss.Left,
		m.Styles.CardName.Render(m.ContactName),
		m.Styles.CardTitle.Render(m.ContactTitle),
	)
	contactInfo := lipgloss.JoinVertical(lipgloss.Left, contactRows...)
	status := m.Styles.CardStatus.Render(m.ContactStatus)

	inner := lipgloss.JoinVertical(lipgloss.Center, header, contactInfo, "", status)
	return m.Styles.BusinessCard.Render(m.Styles.BusinessCardInner.Render(inner))
}
