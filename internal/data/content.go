package data

// Project represents a portfolio project
type Project struct {
	Title       string
	Description string
	Stack       string
	URL         string
}

// SkillCategory represents a group of related skills
type SkillCategory struct {
	Name   string
	Skills []string
}

// LogEntry represents an entry in a timeline or changelog
type LogEntry struct {
	Date        string
	Title       string
	Description string
	URL         string
}

// ContactLink represents a clickable contact link
type ContactLink struct {
	Label string
	Value string
	URL   string
}

// Portfolio represents the content of the portfolio
type Portfolio struct {
	Sections            []string
	Content             map[string]string
	Projects            []Project
	SkillCategories     []SkillCategory
	SkillAcquisitionLog []LogEntry
	OSSContributionLog  []LogEntry
	Dependencies        map[string][]string
	ContactName         string
	ContactTitle        string
	ContactStatus       string
	ContactLinks        []ContactLink
}

// GetDefaultPortfolio returns the default portfolio content
func GetDefaultPortfolio() Portfolio {
	return Portfolio{
		Sections: []string{"About", "Projects", "Skills", "Contact"},
		Content: map[string]string{
			// "About":   "Welcome! I'm a passionate Go developer who loves building **beautiful terminal user interfaces** and high-performance backend systems.\nI enjoy solving complex problems and exploring new technologies. Currently focusing on distributed systems and terminal-based developer tools.",
			"About":   "** This Application is still under construction **\n Thank you For your understanding",
			"Contact": "GitHub: github.com/username\nTwitter: @twitterhandle\nEmail: mail@example.com\nLinkedIn: linkedin.com/in/username",
		},
		Dependencies: map[string][]string{
			"Python": {"Flask", "Scikit-Learn", "PyTorch", "Matplotlib"},
			"Go":     {"Bubble Tea", "Lip Gloss", "Fiber", "Gin"},
			"Rust":   {"Tokio", "Serde", "Actix", "Clap"},
			"Bash":   {"Shell", "Scripting"},
		},
		ContactName:   "Supun Gimres",
		ContactTitle:  "Terminal Enthusiast",
		ContactStatus: "Let’s get to work.",
		ContactLinks: []ContactLink{
			{Label: "GitHub", Value: "github.com/sgimres", URL: "https://github.com/sgimres"},
			{Label: "X", Value: "Who the hell use this..", URL: "https://x.com"},
			{Label: "Email", Value: "mail@example.com", URL: "mailto:mail@example.com"},
			{Label: "LinkedIn", Value: "linkedin.com/in/sgimres", URL: "https://linkedin.com/in/sgimres"},
		},
		Projects: []Project{
			{
				Title:       "🔐 Wardenpy",
				Description: "**Wardenpy**\nA Python-based **CLI password manager** for the security-conscious.\n Uses**Argon2** and **ChaCha20-Poly1305** because your secrets deserve military-grade encryption without the cloud drama.",
				Stack:       "Python, Argon2, ChaCha20",
				URL:         "https://github.com/sgimres/wardenpy",
			},
			{
				Title:       "✍️ Blog CMS",
				Description: "**Blog CMS**\nA custom **PHP engine** built for speed and simplicity.\nLeveraged **AI-assisted 'vibe coding'** workflows to go from concept to a functional architecture in record time.",
				Stack:       "PHP, AI-Assisted Development",
				URL:         "https://github.com/sgimres/blog-cms",
			},
			{
				Title:       "🌱 Agribase",
				Description: "**Agribase**\nA comprehensive **agricultural management system** built as a year-end capstone.\nPairs a snappy **Svelte** frontend with a **PocketBase** backend for real-time data that actually feels lightweight.",
				Stack:       "Svelte, JavaScript, PocketBase",
				URL:         "https://github.com/sgimres/agribase",
			},
		},
		SkillCategories: []SkillCategory{
			{
				Name:   "Languages",
				Skills: []string{"Python", "Rust", "Go", "SQL", "Bash"},
			},
			{
				Name:   "Data",
				Skills: []string{"Pandas", "Scikit-Learn", "Matplotlib", "PyTorch"},
			},
			{
				Name:   "Infra",
				Skills: []string{"Docker", "Podman", "Nix", "Git", "CI/CD"},
			},
			{
				Name:   "Mix",
				Skills: []string{"Neovim", "Linux", "Svelte", "SysAdmin"},
			},
		},
		SkillAcquisitionLog: []LogEntry{
			{
				Date:        "Right Now",
				Title:       "Machine Learning & Deep Learning",
				Description: "Currently teaching computers how to think so I don't have to. Diving deep into neural networks and ML frameworks.",
			},
			{
				Date:        "2025 Q4",
				Title:       "Web Programming II (PHP & Laravel)",
				Description: "Survived the University Web II module by mastering Laravel. Building clean, robust backends—no spaghetti code allowed.",
			},
			{
				Date:        "2024 Q2",
				Title:       "NixOS & Declarative Chaos",
				Description: "Obsessively configuring my self-hosted servers with NixOS. If it's not declarative and reproducible, I don't want it.",
			},
		},
		OSSContributionLog: []LogEntry{
			{
				Date:        "Sep 2025",
				Title:       "auto-cpufreq #883",
				Description: "Fix: support **charging thresholds for Asus ExpertBook laptops**",
				URL:         "https://github.com/AdnanHodzic/auto-cpufreq/pull/883",
			},
			{
				Date:        "Sep 2025",
				Title:       "auto-cpufreq #884",
				Description: "Fix: add additional file paths for edge case",
				URL:         "https://github.com/AdnanHodzic/auto-cpufreq/pull/884",
			},
			{
				Date:        "Jun 2024",
				Title:       "gruvbox-material #206",
				Description: "Integrated native color palette support for the Neovim rainbow-delimiters plugin directly into the overarching theme",
				URL:         "https://github.com/sainnhe/gruvbox-material/pull/206",
			},
		},
	}
}
