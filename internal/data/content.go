package data

// Project represents a portfolio project
type Project struct {
	Title       string
	Description string
	Stack       string
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
			"About":   "Welcome! I'm a passionate Go developer who loves building **beautiful terminal user interfaces** and high-performance backend systems.\nI enjoy solving complex problems and exploring new technologies. Currently focusing on distributed systems and terminal-based developer tools.",
			"Contact": "GitHub: github.com/username\nTwitter: @twitterhandle\nEmail: mail@example.com\nLinkedIn: linkedin.com/in/username",
		},
		Dependencies: map[string][]string{
			"Go":     {"Bubble Tea", "Lip Gloss", "Gin", "Gorm"},
			"Python": {"NumPy", "Pandas", "Scikit-Learn", "FastAPI"},
			"React":  {"Next.js", "Tailwind", "Redux", "Zustand"},
			"Rust":   {"Tokio", "Serde", "Actix", "Rustfmt"},
		},
		ContactName:   "Go Developer",
		ContactTitle:  "Terminal Enthusiast",
		ContactStatus: "available for freelance",
		ContactLinks: []ContactLink{
			{Label: "GitHub", Value: "github.com/username", URL: "https://github.com/username"},
			{Label: "Twitter", Value: "@twitterhandle", URL: "https://twitter.com/twitterhandle"},
			{Label: "Email", Value: "mail@example.com", URL: "mailto:mail@example.com"},
			{Label: "LinkedIn", Value: "linkedin.com/in/username", URL: "https://linkedin.com/in/username"},
		},
		Projects: []Project{
			{
				Title:       "🚀 Gopher-TUI",
				Description: "A powerful toolkit for building **beautiful terminal user interfaces** with reactive components.\nFeatures: **Hot-reloading styles**, custom widgets, and **mouse support**.",
				Stack:       "Go, Bubble Tea, Lip Gloss",
			},
			{
				Title:       "⚡️ FastRoute",
				Description: "A **high-performance HTTP router** with zero allocation and middleware support.\nOptimized for **low-latency microservices** using a modified Radix Tree algorithm.",
				Stack:       "Go, net/http, Radix Tree",
			},
			{
				Title:       "🛡️ Sentinel-Auth",
				Description: "Distributed **authentication service** with JWT and OAuth2 integration.\nDesigned for **high availability** with multi-region session synchronization.",
				Stack:       "Go, Redis, PostgreSQL",
			},
			{
				Title:       "📦 Kube-Watch",
				Description: "Real-time **Kubernetes event monitoring** and notification tool.\nProvides **instant alerts** for deployment failures and resource exhaustion.",
				Stack:       "Go, client-go, Slack SDK",
			},
			{
				Title:       "🌐 NetMesh",
				Description: "A lightweight **service mesh** implementation for resource-constrained environments.\nProvides **mutual TLS** and automatic service discovery with minimal overhead.",
				Stack:       "Go, gRPC, eBPF",
			},
		},
		SkillCategories: []SkillCategory{
			{
				Name:   "Languages",
				Skills: []string{"Go", "Rust", "TypeScript", "C++", "Python"},
			},
			{
				Name:   "Backend",
				Skills: []string{"gRPC", "GraphQL", "Postgres", "Redis", "NATS"},
			},
			{
				Name:   "Frontend",
				Skills: []string{"React", "Tailwind", "Lip Gloss", "Bubble Tea"},
			},
			{
				Name:   "Infra",
				Skills: []string{"K8s", "Docker", "Terraform", "AWS", "CI/CD"},
			},
		},
		SkillAcquisitionLog: []LogEntry{
			{
				Date:        "2025 Q4",
				Title:       "Advanced Rust Mastery",
				Description: "Mastered memory safety and concurrency patterns in high-load systems.",
			},
			{
				Date:        "2024 Q2",
				Title:       "K8s Operator Development",
				Description: "Built custom controllers for automated infrastructure management.",
			},
			{
				Date:        "2023 Q1",
				Title:       "Distributed Systems Architecture",
				Description: "Deep dive into consensus algorithms and eventual consistency.",
			},
		},
		OSSContributionLog: []LogEntry{
			{
				Date:        "Feb 2026",
				Title:       "Bubble Tea #1283",
				Description: "Implemented advanced mouse scroll event handling for grid components.",
				URL:         "https://github.com/charmbracelet/bubbletea/pull/1283",
			},
			{
				Date:        "Dec 2025",
				Title:       "Lip Gloss #455",
				Description: "Fixed color profile detection on rare terminal emulators.",
				URL:         "https://github.com/charmbracelet/lipgloss/pull/455",
			},
			{
				Date:        "Oct 2025",
				Title:       "Go-Redis #2102",
				Description: "Improved connection pooling efficiency for high-latency networks.",
				URL:         "https://github.com/go-redis/redis/pull/2102",
			},
		},
	}
}
