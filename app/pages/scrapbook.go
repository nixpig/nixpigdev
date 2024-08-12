package pages

// import (
// 	"fmt"
// 	"strings"
//
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/glamour"
// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/mmcdole/gofeed"
// 	"github.com/nixpig/nixpigdev/app/theme"
// )
//
// type blogPostsMsg *gofeed.Feed
//
// type blogItem struct {
// 	title string
// 	link  string
// }
//
// type scrapbook struct {
// 	title       string
// 	description string
// 	blogItems   []blogItem
// }
//
// var fp = gofeed.NewParser()
//
// var Scrapbook = scrapbook{
// 	title:       "Scrapbook",
// 	description: "Notes, blogs, gists…",
// }
//
// func (sb *scrapbook) Init() tea.Cmd {
// 	return sb.getBlogPosts
// }
//
// func (sb *scrapbook) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case blogPostsMsg:
// 		sb.blogItems = []blogItem{}
// 		if msg == nil {
// 			fmt.Println("blog posts msg is nil")
// 		} else {
// 			for _, item := range msg.Items {
// 				sb.blogItems = append(sb.blogItems, blogItem{
// 					title: item.Title,
// 					link:  item.GUID,
// 				})
// 			}
// 		}
// 	}
//
// 	return nil, nil
// }
//
// func (sb *scrapbook) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
// 	tr, err := glamour.NewTermRenderer(
// 		glamour.WithStylePath("dracula"),
// 		glamour.WithWordWrap(s.Width/2-2),
// 	)
// 	if err != nil {
// 		return fmt.Sprintf("Failed to create term renderer: %s", err)
// 	}
//
// 	container := renderer.NewStyle()
// 	padded := renderer.NewStyle().
// 		PaddingLeft(2).
// 		PaddingRight(2)
//
// 	blogs := strings.Builder{}
//
// 	for _, item := range sb.blogItems {
// 		blogs.WriteString(fmt.Sprintf("- [%s](%s)", item.title, item.link))
// 	}
//
// 	b, err := tr.Render(fmt.Sprintf("## Recent blogs\n%s", blogs.String()))
// 	if err != nil {
// 		return fmt.Sprintf("failed to render recent blogs: %s", err)
// 	}
//
// 	row := container.Render(b)
//
// 	return strings.Join([]string{
// 		md("# Scrapbook"),
// 		padded.Foreground(lipgloss.Color(theme.Dracula.Foreground)).Render("Just some stuff...\n"),
// 		row,
// 	}, "")
// }
//
// func (sb *scrapbook) Title() string {
// 	return sb.title
// }
//
// func (sb *scrapbook) Description() string {
// 	return sb.description
// }
//
// func (sb *scrapbook) FilterValue() string {
// 	return fmt.Sprintf("%s %s", sb.title, sb.description)
// }
//
// func (sb *scrapbook) getBlogPosts() tea.Msg {
// 	fetched, err := fp.ParseURL("https://medium.com/feed/@nixpig")
// 	if err != nil {
// 		fmt.Println(fmt.Errorf("failed to fetch feed: %w", err))
// 	}
//
// 	return blogPostsMsg(fetched)
// }
