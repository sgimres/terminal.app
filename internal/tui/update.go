package tui

import (
	"strings"
	"time"

	"profile/internal/data"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

// Update handles messages and updates the model state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case LoadedMsg:
		m.Loading = false
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case ClearCopyURLMsg:
		m.CopyURL = ""
		return m, nil

	case tea.KeyPressMsg:
		if m.SignModal {
			if m.SignSuccess || m.SignError != "" {
				m.SignModal = false
				m.SignSuccess = false
				m.SignError = ""
				return m, nil
			}
			switch msg.String() {
			case "esc":
				m.SignModal = false
				m.SignName = ""
				m.SignDescription = ""
				m.SignSuccess = false
				m.SignError = ""
				return m, nil
			case "tab":
				if m.SignField == 0 {
					m.SignField = 1
				} else {
					m.SignField = 0
				}
				return m, nil
			case "enter":
				if m.SignName == "" || m.SignDescription == "" {
					m.SignError = "please fill all fields"
					return m, nil
				}
				err := data.AddGuestbookEntry(m.SignName, m.SignDescription, data.GetDB())
				if err != nil {
					m.SignError = "error saving: " + err.Error()
					return m, nil
				}
				m.SignSuccess = true
				return m, nil
			case "backspace":
				if m.SignField == 0 {
					if len(m.SignName) > 0 {
						m.SignName = m.SignName[:len(m.SignName)-1]
					}
				} else {
					if len(m.SignDescription) > 0 {
						m.SignDescription = m.SignDescription[:len(m.SignDescription)-1]
					}
				}
				return m, nil
			default:
				keyStr := msg.String()
				if keyStr == "space" {
					keyStr = " "
				}
				if keyStr != "" && keyStr != "space" {
					if m.SignField == 0 {
						m.SignName += keyStr
					} else {
						m.SignDescription += keyStr
					}
				}
				return m, nil
			}
		}

		if m.Filtering {
			switch msg.String() {
			case "esc", "enter":
				m.Filtering = false
				m.FilterText = ""
				m.ActiveProject = 0
			case "backspace":
				if len(m.FilterText) > 0 {
					m.FilterText = m.FilterText[:len(m.FilterText)-1]
					m.updateFilter()
				}
			default:
				if len(msg.String()) == 1 {
					m.FilterText += msg.String()
					m.updateFilter()
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.Sections[m.ActiveSection] == "Skills" {
				m.SkillGridFocus = !m.SkillGridFocus
				return m, nil
			}
		case "/":
			if m.Sections[m.ActiveSection] == "Projects" {
				m.Filtering = true
				m.FilterText = ""
				m.updateFilter()
				return m, nil
			}
		case "enter":
			if m.Sections[m.ActiveSection] == "Projects" {
				if m.ActiveProject < len(m.FilteredProjects) {
					url := m.FilteredProjects[m.ActiveProject].URL
					if url != "" {
						m.CopyURL = url
						return m, clearCopyURLAfterDelay()
					}
				}
			}
			if m.Sections[m.ActiveSection] == "Skills" && !m.SkillGridFocus {
				if m.ActiveSkillLog < len(m.OSSContributionLog) {
					m.CopyURL = m.OSSContributionLog[m.ActiveSkillLog].URL
					return m, clearCopyURLAfterDelay()
				}
			}
			if m.Sections[m.ActiveSection] == "Contact" {
				if m.ActiveContactLink < len(m.ContactLinks) {
					m.CopyURL = m.ContactLinks[m.ActiveContactLink].URL
					return m, clearCopyURLAfterDelay()
				}
			}
		case "up", "k":
			if !m.Loading {
				switch m.Sections[m.ActiveSection] {
				case "Projects":
					if m.ActiveProject > 0 {
						m.ActiveProject--
					}
				case "Skills":
					if m.SkillGridFocus {
						if m.ActiveSkillIndex > 0 {
							m.ActiveSkillIndex--
						}
					} else {
						if m.ActiveSkillLog > 0 {
							m.ActiveSkillLog--
						}
					}
				case "Contact":
					if m.ActiveContactLink > 0 {
						m.ActiveContactLink--
					}
				default:
					if m.ActiveSection > 0 {
						m.ActiveSection--
					}
				}
			}
		case "down", "j":
			if !m.Loading {
				switch m.Sections[m.ActiveSection] {
				case "Projects":
					if m.ActiveProject < len(m.FilteredProjects)-1 {
						m.ActiveProject++
					}
				case "Skills":
					if m.SkillGridFocus {
						if m.ActiveSkillIndex < len(m.SkillCategories[m.ActiveCategory].Skills)-1 {
							m.ActiveSkillIndex++
						}
					} else {
						if m.ActiveSkillLog < len(m.OSSContributionLog)-1 {
							m.ActiveSkillLog++
						}
					}
				case "Contact":
					if m.ActiveContactLink < len(m.ContactLinks)-1 {
						m.ActiveContactLink++
					}
				default:
					if m.ActiveSection < len(m.Sections)-1 {
						m.ActiveSection++
					}
				}
			}
		case "left", "h":
			if !m.Loading {
				if m.Sections[m.ActiveSection] == "Skills" && m.SkillGridFocus {
					if m.ActiveCategory > 0 {
						m.ActiveCategory--
						m.ActiveSkillIndex = 0
					}
				} else if m.ActiveSection > 0 {
					m.ActiveSection--
				}
			}
		case "right", "l":
			if !m.Loading {
				if m.Sections[m.ActiveSection] == "Skills" && m.SkillGridFocus {
					if m.ActiveCategory < len(m.SkillCategories)-1 {
						m.ActiveCategory++
						m.ActiveSkillIndex = 0
					}
				} else if m.ActiveSection < len(m.Sections)-1 {
					m.ActiveSection++
				}
			}
		}

		// Handle shortcuts
		if !m.Loading {
			key := strings.ToLower(msg.String())
			for i, section := range m.Sections {
				shortcut := strings.ToLower(string(section[0]))
				if key == shortcut {
					m.ActiveSection = i
					m.SkillGridFocus = false
					return m, nil
				}
			}
			if key == "g" && m.Sections[m.ActiveSection] == "Contact" {
				m.SignModal = true
				m.SignName = ""
				m.SignDescription = ""
				m.SignField = 0
				m.SignSuccess = false
				m.SignError = ""
				return m, nil
			}
		}
	}
	return m, nil
}

func (m *Model) updateFilter() {
	m.FilteredProjects = nil
	query := strings.ToLower(m.FilterText)
	for _, p := range m.Projects {
		if query == "" || strings.Contains(strings.ToLower(p.Title), query) || strings.Contains(strings.ToLower(p.Stack), query) {
			m.FilteredProjects = append(m.FilteredProjects, p)
		}
	}
	if m.ActiveProject >= len(m.FilteredProjects) {
		m.ActiveProject = 0
	}
}

type ClearCopyURLMsg struct{}

func clearCopyURLAfterDelay() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(5 * time.Second)
		return ClearCopyURLMsg{}
	}
}
