package lsx

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

const listHeight = 20
const defaultWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(1).Italic(true).Foreground(lipgloss.Color("#61D1C2")).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#7E6CFA")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(0)
)

type Item string

func (i Item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := string(i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type Model struct {
	list   list.Model
	choice string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "ctrl+c" {
			// TODO: write to file
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = filepath.Join(m.choice, string(i))
				var items []list.Item
				m.list.Title = fmt.Sprintf("📌 %s", m.choice)
				files, err := ioutil.ReadDir(m.choice)
				if err != nil {
					log.Fatal(err)
				}
				for _, f := range files {
					items = append(items, Item(f.Name()))
				}
				m.list.SetItems(items)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.choice != "" {
		return m.list.View()
	}
	return "\n" + m.list.View()
}

// func main() {

// 	l := list.NewModel(items, itemDelegate{}, defaultWidth, listHeight)
// 	l.SetShowStatusBar(true)
// 	l.SetFilteringEnabled(true)
// 	l.Styles.Title = titleStyle
// 	l.Styles.PaginationStyle = paginationStyle
// 	l.Styles.HelpStyle = helpStyle

// 	m := Model{list: l}
// 	m.list.Title = "My Fave Things"

// 	p := tea.NewProgram(m)
// 	p.EnterAltScreen()

// 	if err := p.Start(); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }
