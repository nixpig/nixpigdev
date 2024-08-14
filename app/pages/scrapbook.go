package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/mmcdole/gofeed"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/theme"
)

type blogPostsMsg *gofeed.Feed

type blogItem struct {
	title string
	link  string
}

type scrapbookModel struct {
	title        string
	description  string
	blogItems    []blogItem
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
}

var fp = gofeed.NewParser()

func NewScrapbook(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) scrapbookModel {
	return scrapbookModel{
		title:       "Scrapbook",
		description: "Notes, blogs, gists…",
		renderer:    renderer,
		md:          md,
	}
}

func (s scrapbookModel) Init() tea.Cmd {
	fmt.Println("initialising scrapbook")
	return s.getBlogPosts
}

func (s scrapbookModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case blogPostsMsg:
		fmt.Println("received blogs")
		s.blogItems = []blogItem{}
		if msg == nil {
			fmt.Println("blog posts msg is nil")
		} else {
			for _, item := range msg.Items {
				s.blogItems = append(s.blogItems, blogItem{
					title: item.Title,
					link:  item.GUID,
				})
			}
		}

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
		glamour.WithWordWrap(s.contentWidth/2-2),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create term renderer: %s", err)
	}

	container := s.renderer.NewStyle()
	padded := s.renderer.NewStyle().
		PaddingLeft(2).
		PaddingRight(2)

	blogs := strings.Builder{}

	for _, item := range s.blogItems {
		blogs.WriteString(fmt.Sprintf("- [%s](%s)", item.title, item.link))
	}

	b, err := tr.Render(fmt.Sprintf("## Recent blogs\n%s", blogs.String()))
	if err != nil {
		return fmt.Sprintf("failed to render recent blogs: %s", err)
	}

	row := container.Render(b)

	return strings.Join([]string{
		s.md("# Scrapbook", s.contentWidth),
		padded.Foreground(lipgloss.Color(theme.Dracula.Foreground)).Render("Just some stuff...\n"),
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

func (s scrapbookModel) getBlogPosts() tea.Msg {
	fmt.Println("fetching blogs...")
	fetched, err := fp.ParseURL("https://medium.com/feed/@nixpig")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to fetch feed: %w", err))
	}

	fmt.Println("fetched blogs.")
	return blogPostsMsg(fetched)
}
