package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/mmcdole/gofeed"
	"github.com/nixpig/nixpigdev/internal/commands"
	"github.com/nixpig/nixpigdev/pkg/markdown"
	"github.com/nixpig/nixpigdev/pkg/theme"
)

type blogItem struct {
	title string
	link  string
	date  string
}

type scrapbookModel struct {
	title        string
	description  string
	blogItems    []blogItem
	termRenderer *lipgloss.Renderer
	mdRenderer   markdown.Renderer
	contentWidth int
}

func NewScrapbook(
	termRenderer *lipgloss.Renderer,
	mdRenderer markdown.Renderer,
) scrapbookModel {
	return scrapbookModel{
		title:        "Scrapbook",
		description:  "Notes, blogs, gists…",
		termRenderer: termRenderer,
		mdRenderer:   mdRenderer,
	}
}

func (s scrapbookModel) Init() tea.Cmd {
	fp := gofeed.NewParser()
	return commands.FetchFeedCmd(fp)
}

func (s scrapbookModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.FetchFeedSuccessMsg:
		s.blogItems = []blogItem{}
		if msg == nil {
			fmt.Println("failed to load blog posts")
		} else {
			for _, item := range msg.Items {
				s.blogItems = append(s.blogItems, blogItem{
					title: item.Title,
					link:  item.GUID,
					date:  item.Published,
				})
			}
		}

		return s, nil

	case commands.FetchFeedErrMsg:
		fmt.Println(msg)
		return s, nil

	case commands.SectionSizeMsg:
		s.contentWidth = msg.Width
		return s, nil
	}

	return s, nil
}

func (s scrapbookModel) View() string {
	tr, err := glamour.NewTermRenderer(
		glamour.WithStylePath("dracula"),
		glamour.WithWordWrap(s.contentWidth-2),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create term renderer: %s", err)
	}

	container := s.termRenderer.NewStyle()
	padded := s.termRenderer.NewStyle().
		PaddingLeft(2).
		PaddingRight(2)

	blogs := strings.Builder{}

	for _, item := range s.blogItems {
		blogs.WriteString(fmt.Sprintf(
			"- 🗓 %s \n   [%s](%s)",
			item.date,
			item.title,
			item.link,
		))
	}

	b, err := tr.Render(fmt.Sprintf("## Recent blogs\n%s", blogs.String()))
	if err != nil {
		return fmt.Sprintf("failed to render recent blogs: %s", err)
	}

	row := container.Render(b)

	return strings.Join([]string{
		s.mdRenderer("# Scrapbook", s.contentWidth),
		padded.
			Foreground(lipgloss.Color(theme.Dracula.Foreground)).
			Render("Just some stuff...\n"),
		row,
	}, "")
}

func (s scrapbookModel) Title() string {
	return s.title
}

func (s scrapbookModel) Description() string {
	return s.description
}

func (s scrapbookModel) FilterValue() string {
	return fmt.Sprintf("%s %s", s.title, s.description)
}
