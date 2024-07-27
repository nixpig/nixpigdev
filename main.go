package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/nixpig/nixpigdev/app"
	"github.com/nixpig/nixpigdev/logging"
	"github.com/rs/zerolog"
)

const (
	hostname = "localhost"
	port     = "23234"
)

// TODO: recover from panics
// TODO: rate limiting

var pages = []app.Page{
	{
		PageTitle: "🏡 Home",
		Desc:      "The home page",
		Filepath:  "pages/home.md",
	},
	{
		PageTitle: "🏗️ Projects",
		Desc:      "Open source stuff",
		Filepath:  "pages/projects.md",
	},
	{
		PageTitle: "💻️ Uses",
		Desc:      "The stuff I use",
		Filepath:  "pages/uses.md",
	},
	{
		PageTitle: "📬️ Contact",
		Desc:      "Come say hi!",
		Filepath:  "pages/contact.md",
	},
}

func main() {
	var logger = zerolog.
		New(os.Stdout).
		Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02T15:04:05.999Z07:00",
		}).With().Timestamp().Logger()

	wg := &sync.WaitGroup{}
	for i, page := range pages {
		logger.Info().Str("page", page.Filepath).Msg("loading file")
		wg.Add(1)
		go func(i int, filepath string) {
			defer wg.Done()
			if filepath != "" {
				content, err := os.ReadFile(filepath)
				if err != nil {
					logger.Error().Err(err).Str("filepath", filepath).Msg("failed to load file content")
					return
				}

				pages[i].Content = string(content)
			}
		}(i, page.Filepath)
	}
	wg.Wait()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(hostname, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(&logger),
		),
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create server")
		os.Exit(1)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info().Str("host", hostname).Str("port", port).Msg("starting server")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			logger.Error().Err(err).Msg("failed to start server")
			done <- nil
		}
	}()

	<-done
	logger.Info().Msg("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		logger.Error().Err(err).Msg("failed to shutdown server gracefully")
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	renderer := bubbletea.MakeRenderer(s)

	m := app.Model{
		Content: app.NewContent(renderer, pages),
		Nav:     app.NewNav(renderer, pages),
		Footer:  app.NewFooter(renderer, app.InputKeys),
	}

	return m, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
}
