package commands

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mmcdole/gofeed"
)

type NavigatePageMsg int

type SectionSizeMsg struct {
	Width  int
	Height int
}

type FetchFeedSuccessMsg *gofeed.Feed
type FetchFeedErrMsg error

type SendEmailSuccessMsg string
type SendEmailErrMsg error

type FetchProjectsSuccessMsg []struct {
	Name        string   `json:"name"`
	HTMLURL     string   `json:"html_url"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}
type FetchProjectErrMsg error

func NavigatePageCmd(i int) tea.Cmd {
	return func() tea.Msg {
		return NavigatePageMsg(i)
	}
}

func FetchFeedCmd(fp *gofeed.Parser) tea.Cmd {
	return func() tea.Msg {
		fetched, err := fp.ParseURL("https://medium.com/feed/@nixpig")
		if err != nil {
			return FetchFeedErrMsg(fmt.Errorf("failed to fetch blog feed: %w", err))
		}

		return FetchFeedSuccessMsg(fetched)
	}
}

func SendEmailCmd(name, email, message string) tea.Cmd {
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

func FetchProjects() tea.Cmd {
	return func() tea.Msg {
		res, err := http.Get("https://api.github.com/users/nixpig/repos?sort=updated")
		if err != nil || res.StatusCode != http.StatusOK {
			return FetchProjectErrMsg(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return FetchProjectErrMsg(err)
		}

		var projects FetchProjectsSuccessMsg
		var filteredProjects FetchProjectsSuccessMsg

		if err := json.Unmarshal(body, &projects); err != nil {
			return FetchProjectErrMsg(err)
		}

		for _, p := range projects {
			if slices.IndexFunc(p.Topics, func(t string) bool {
				return t == "featured"
			}) != -1 {
				filteredProjects = append(filteredProjects, p)
			}
		}

		return FetchProjectsSuccessMsg(filteredProjects)
	}
}
