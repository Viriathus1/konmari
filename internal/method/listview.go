package method

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type fileItem struct {
	path string
	info os.FileInfo
}

func (fi fileItem) Title() string       { return fi.path }
func (fi fileItem) FilterValue() string { return fi.path }
func (fi fileItem) Description() string { return fi.info.ModTime().Format("2006-01-02 15:04") }

type ListViewModel struct {
	list     list.Model
	selected map[string]bool
}

func NewListView(paths []string) tea.Model {
	items := []list.Item{}
	selected := map[string]bool{}

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil || fileInfo.IsDir() {
			return nil
		}
		selected[path] = true
		items = append(items, fileItem{path: path, info: fileInfo})
	}

	delegate := toggleDelegate{selected: selected}
	l := list.New(items, delegate, 80, 20)
	l.Title = "Do they spark joy? (space to toggle and enter to confirm)"
	return ListViewModel{
		list:     l,
		selected: selected,
	}
}

func (lv ListViewModel) Init() tea.Cmd {
	return nil
}

func (lv ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "enter":
			return lv, tea.Quit
		case " ":
			if file, ok := lv.list.SelectedItem().(fileItem); ok {
				lv.selected[file.path] = !lv.selected[file.path]
			}
		}
	}
	var cmd tea.Cmd
	lv.list, cmd = lv.list.Update(msg)
	return lv, cmd
}

func (lv ListViewModel) View() string {
	return lv.list.View()
}

func (lv ListViewModel) SelectedPaths() []string {
	var paths []string

	for path, ok := range lv.selected {
		if ok {
			paths = append(paths, path)
		}

	}

	return paths
}
