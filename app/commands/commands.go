package commands

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"

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

type SendEmailSuccessMsg string
type SendEmailErrMsg error

func FetchFeed(fp *gofeed.Parser) tea.Cmd {
	return func() tea.Msg {
		fetched, err := fp.ParseURL("https://medium.com/feed/@nixpig")
		if err != nil {
			return FetchFeedErrMsg(fmt.Errorf("failed to fetch blog feed: %w", err))
		}

		return FetchFeedSuccessMsg(fetched)
	}
}

func SendEmail(name, email, message string) tea.Cmd {
	return func() tea.Msg {
		identity := os.Getenv("SMTP_AUTH_IDENTITY")
		username := os.Getenv("SMTP_AUTH_USERNAME")
		password := os.Getenv("SMTP_AUTH_PASSWORD")
		host := os.Getenv("SMTP_AUTH_HOST")
		to := os.Getenv("MAIL_RECIPIENT")
		send := os.Getenv("SMTP_SEND_ADDR")

		auth := smtp.PlainAuth(
			identity,
			username,
			password,
			host,
		)

		header := make(map[string]string)
		header["From"] = to
		header["To"] = to
		header["Subject"] = "nixpig.dev contact form"
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"

		msg := strings.Builder{}

		for k, v := range header {
			msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
		}

		msg.WriteString(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\n\n", email))))
		msg.WriteString(base64.StdEncoding.EncodeToString([]byte(message)))

		err := smtp.SendMail(
			send,
			auth,
			to,
			[]string{to},
			[]byte(msg.String()),
		)
		if err != nil {
			fmt.Println("error: ", err)
			return SendEmailErrMsg(errors.New("error sending message"))
		}

		return SendEmailSuccessMsg("email sent")
	}
}
