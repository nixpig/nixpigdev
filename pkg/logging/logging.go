package logging

import (
	"github.com/charmbracelet/ssh"
	"github.com/rs/zerolog"
)

func Middleware(logger *zerolog.Logger) func(next ssh.Handler) ssh.Handler {
	return func(next ssh.Handler) ssh.Handler {
		return func(sess ssh.Session) {
			logger.Info().
				Str("session", sess.Context().SessionID()).
				Str("user", sess.User()).
				Str("address", sess.RemoteAddr().String()).
				Bool("publickey", sess.PublicKey() != nil).
				Str("client", sess.Context().ClientVersion()).
				Msg("connect")

			next(sess)

			logger.Info().
				Str("session", sess.Context().SessionID()).
				Msg("disconnect")

		}
	}
}
