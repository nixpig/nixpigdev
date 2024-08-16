package commands

import "github.com/mmcdole/gofeed"

type PageNavigationMsg int

type SectionSizeMsg struct {
	Width  int
	Height int
}

type FeedFetchMsg *gofeed.Feed
