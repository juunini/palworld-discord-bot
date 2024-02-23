package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	quitTextStyle := lipgloss.NewStyle().Margin(1, 0, 2, 4)

	if m.choice != "" {
		return quitTextStyle.Render(m.choice)
	}
	if m.quitting {
		return "\n"
	}
	return "\n" + m.list.View()
}

func Choose(title string, items []string) (string, error) {
	l := list.New(stringsToItems(items), itemDelegate{}, 20, 14)
	l.Title = "Select language"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = lipgloss.NewStyle().MarginLeft(2)
	l.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)

	m := model{list: l}

	userSelect, err := tea.NewProgram(m).Run()
	if err != nil {
		return "", err
	}

	selectedModel, ok := userSelect.(model)
	if !ok {
		return "", fmt.Errorf("unexpected model type: %T", userSelect)
	}
	if selectedModel.choice == "" {
		return "", fmt.Errorf("no language selected")
	}

	return selectedModel.choice, nil
}

func stringsToItems(items []string) []list.Item {
	_items := []list.Item{}
	for _, i := range items {
		_items = append(_items, item(i))
	}
	return _items
}
