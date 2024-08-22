package commands

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mmcdole/gofeed"
)

type PageNavigationMsg int

type SectionSizeMsg struct {
	Width  int
	Height int
}

type FetchFeedSuccessMsg *gofeed.Feed
type FetchFeedErrMsg error

func FetchFeed(fp *gofeed.Parser) tea.Cmd {
	return func() tea.Msg {
		fetched, err := fp.ParseURL("https://medium.com/feed/@nixpig")
		if err != nil {
			return FetchFeedErrMsg(fmt.Errorf("failed to fetch blog feed: %w", err))
		}

		return FetchFeedSuccessMsg(fetched)
	}
}

type SendEmailSuccessMsg string
type SendEmailErrMsg error

func SendEmail(name, email, message string) tea.Cmd {
	return func() tea.Msg {
		auth := smtp.PlainAuth(
			os.Getenv("SMTP_AUTH_IDENTITY"),
			os.Getenv("SMTP_AUTH_USERNAME"),
			os.Getenv("SMTP_AUTH_PASSWORD"),
			os.Getenv("SMTP_AUTH_HOST"),
		)

		to := []string{os.Getenv("MAIL_RECIPIENT")}
		msg := []byte(message)
		err := smtp.SendMail(os.Getenv("SMTP_SEND_ADDR"), auth, email, to, msg)
		if err != nil {
			fmt.Println("error: ", err)
			return SendEmailErrMsg(errors.New("error sending message"))
		}

		return SendEmailSuccessMsg("email sent")
	}
}
