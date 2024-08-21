package commands

import "github.com/mmcdole/gofeed"

type PageNavigationMsg int

type SectionSizeMsg struct {
	Width  int
	Height int
}

type FeedFetchSuccessMsg *gofeed.Feed
type FeedFetchErrMsg error
