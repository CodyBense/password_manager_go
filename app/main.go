package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var Columns = []table.Column{
		{Title: "Website", Width: 25},
		{Title: "Username", Width: 25},
		{Title: "Password", Width: 25},
	}
type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
        case "d":
            return m.DeleteCurrent()
        case "a":
            return m.AddLogin()
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
}

func Main() {
    m := CreateTable()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func CreateTable() model{
    rows := SqlList()

	t := table.New(
		table.WithColumns(Columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}

    return m
}

func (m model) DeleteCurrent() (tea.Model, tea.Cmd) {
    SqlRemove(m.table.SelectedRow()[0])
    m = CreateTable()

    var cmd tea.Cmd
    m.table, cmd = m.table.Update(nil)
    return m, cmd
}

func (m model) AddLogin() (tea.Model, tea.Cmd) {
    SqlAdd("test", "test", "test")
    m = CreateTable()

    var cmd tea.Cmd
    m.table, cmd = m.table.Update(nil)
    return m, cmd
}
