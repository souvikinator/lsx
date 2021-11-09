package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/souvikinator/lsx/lsx"
	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

var App lsx.Lsx
var currentPath string

/*UI*/

const listHeight = 20
const defaultWidth = 50

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(1).Italic(true).Foreground(lipgloss.Color("#61D1C2")).Bold(true).MaxWidth(100)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#7E6CFA")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(0)
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
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
		// TODO: can't go before /home/souvikinator
		if msg.String() == "enter" {
			i, ok := m.list.SelectedItem().(item)
			home, _ := os.UserHomeDir()
			if ok {
				m.choice = filepath.Join(m.choice, string(i))
				var items []list.Item
				m.list.Title = fmt.Sprintf("📌 %s", strings.Replace(m.choice, home, "~", -1))
				App.ClearDirs()
				App.GetPathContent(m.choice)
				dirs := App.GetDirs()
				if !App.AllMode {
					dirs = utils.GetNonDotDirs(dirs)
				}

				items = append(items, item(".."))
				for _, f := range dirs {
					items = append(items, item(f))
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

/*UI End*/
//FIXME: lsx error: some error occured stat /home/souvikinator/.test: no such file or directory

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lsx",
	Version: App.Version,
	Args:    cobra.ArbitraryArgs,
	Short:   " A command line utility, let's you navigate across the terminal like butter",
	Long:    `lsx stands for "ls Xtended". It let's you navigate across the terminal using cursors along with search. One can also set aliases for paths.`,
	Run: func(cmd *cobra.Command, args []string) {

		home := utils.HomeDir()
		var items []list.Item

		App.AllMode, _ = cmd.Flags().GetBool("all")

		// if alias is passed
		if len(args) > 0 {
			pathAlias := args[0]
			ChdirToAlias(pathAlias)
			os.Exit(0)
		}

		currentPath, _ = os.Getwd()

		// get all the directories from the current path
		// App.ClearDirs() TODO: in Update()
		App.GetPathContent(currentPath)

		dirs := App.GetDirs()
		// remove all directories starting with .
		// if -a/--all is not used
		if !App.AllMode {
			dirs = utils.GetNonDotDirs(dirs)
		}

		items = append(items, item(".."))
		for _, f := range dirs {
			items = append(items, item(f))
		}

		// currentPath = filepath.Join(currentPath, selectedDir)
		l := list.NewModel(items, itemDelegate{}, defaultWidth, listHeight)
		l.SetShowStatusBar(true)
		l.SetFilteringEnabled(true)
		//TODO: add style for status bar
		l.Styles.Title = titleStyle
		l.Styles.PaginationStyle = paginationStyle
		l.Styles.HelpStyle = helpStyle

		m := Model{list: l, choice: currentPath}
		m.list.Title = fmt.Sprintf("📌 %s", strings.Replace(currentPath, home, "~", -1))

		p := tea.NewProgram(m)
		p.EnterAltScreen()

		if err := p.Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	App.Init()
	// commands
	rootCmd.AddCommand(aliasCmd)
	rootCmd.AddCommand(removeAliasCmd)
	rootCmd.AddCommand(setAliasCmd)

	//flags
	rootCmd.Flags().BoolP("all", "a", false, "display hidden (dotdirs) directories as well")
}

func ChdirToAlias(pathAlias string) {
	if len(App.Alias) == 0 {
		utils.Warn("no alias found in records")
		utils.Warn("get started by using 'set-alias' command")
		os.Exit(0)
	}

	moveToPath, ok := App.Alias[pathAlias]
	if !ok {
		utils.Err("alias '" + pathAlias + "' not found")
	}

	// path corresponding to alias exists?
	if !utils.PathExists(moveToPath) {
		utils.Err("'" + dirPath + "' not found, make sure that the path exists and is a directory")
	}

	utils.WriteToFile(App.TempFile, moveToPath)

}
