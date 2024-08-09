package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mmcdole/gofeed"
)

func Notes(renderer *lipgloss.Renderer) Page {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://medium.com/@nixpig/feed")

	blogs := make([]string, len(feed.Items))

	for i, item := range feed.Items {
		blogs[i] = fmt.Sprintf("- _%s_\n[%s](%s)\n&nbsp;", item.Published, item.Title, item.Link)
	}

	notes := Page{
		title:       "Notes",
		description: "Blogs and stuff",
		renderer:    renderer,
		content: strings.Join([]string{
			"# Notes",
			strings.Join(blogs, "\n"),
		}, "\n\n"),
	}

	return notes
}
