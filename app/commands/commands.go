package commands

import (
	"errors"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mmcdole/gofeed"
)

type PageNavigationMsg int

type SectionSizeMsg struct {
	Width  int
	Height int
}

type FeedFetchSuccessMsg *gofeed.Feed
type FeedFetchErrMsg error

type SendEmailSuccessMsg string
type SendEmailErrMsg error

func SendEmail(name, email, message string) tea.Cmd {
	return func() tea.Msg {
		// TODO: send email
		fmt.Println("send email: ", name, email, message)
		time.Sleep(time.Duration(time.Second * 2))
		return SendEmailErrMsg(errors.New("error sending message"))
	}
}
