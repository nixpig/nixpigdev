package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
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
	hostname = "0.0.0.0"
	port     = "23234"
)

// TODO: recover from panics
// TODO: rate limiting

func main() {
	var logger = zerolog.
		New(os.Stdout).
		Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02T15:04:05.999Z07:00",
		}).With().Timestamp().Logger()

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

func teaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	renderer := bubbletea.MakeRenderer(sess)

	pty, _, active := sess.Pty()
	if !active {
		wish.Fatalln(sess, "no active pty")

	}

	return app.New(pty, renderer), []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
}
