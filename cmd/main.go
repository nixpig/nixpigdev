package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	app "github.com/nixpig/nixpigdev/internal"
	"github.com/nixpig/nixpigdev/pkg/logging"
	"github.com/nixpig/nixpigdev/pkg/markdown"
	"github.com/rs/zerolog"
)

const (
	hostname = "0.0.0.0"
	ssh_port = "23234"
	web_port = "8080"
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

	sshServer, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(hostname, ssh_port)),
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
		logger.Info().Str("host", hostname).Str("port", ssh_port).Msg("starting ssh server")
		if err := sshServer.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			logger.Error().Err(err).Msg("failed to start ssh server")
			done <- nil
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		index, err := os.ReadFile("web/index.html")
		if err != nil {
			fmt.Println(fmt.Errorf("read file: %w", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(index)
	})

	webServer := http.Server{
		Addr:         fmt.Sprintf(":%v", web_port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	go func() {
		logger.Info().Str("host", hostname).Str("port", web_port).Msg("starting web server")
		if err := webServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Msg("failed to start web server")
			done <- nil
		}
	}()

	<-done
	logger.Info().Msg("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := sshServer.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		logger.Error().Err(err).Msg("failed to shutdown server gracefully")
	}
}

func teaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	termRenderer := bubbletea.MakeRenderer(sess)

	pty, _, active := sess.Pty()
	if !active {
		wish.Fatalln(sess, "no active pty")

	}

	return app.New(
			pty,
			termRenderer,
			markdown.Render,
		), []tea.ProgramOption{
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		}
}
