package tui

import (
	"os/exec"
	"runtime"
	"strings"

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

	case tea.KeyPressMsg:
		if m.Filtering {
			switch msg.String() {
			case "esc", "enter":
				m.Filtering = false
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
			if m.Sections[m.ActiveSection] == "Skills" && !m.SkillGridFocus {
				if m.ActiveSkillLog < len(m.OSSContributionLog) {
					m.openBrowser(m.OSSContributionLog[m.ActiveSkillLog].URL)
				}
			}
			if m.Sections[m.ActiveSection] == "Contact" {
				if m.ActiveContactLink < len(m.ContactLinks) {
					m.openBrowser(m.ContactLinks[m.ActiveContactLink].URL)
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
					m.SkillGridFocus = false // Reset focus when switching tabs
					return m, nil
				}
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

func (m Model) openBrowser(url string) {
	if url == "" {
		return
	}
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux and others
		cmd = "xdg-open"
		args = []string{url}
	}
	_ = exec.Command(cmd, args...).Start()
}
