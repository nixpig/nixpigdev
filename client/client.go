package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

const (
	terminal = "xterm-256color"
)

type SSHClient struct {
	client *ssh.Client
}

func (s *SSHClient) Close() error {
	if err := s.client.Close(); err != nil {
		return fmt.Errorf("failed to close ssh client: %w", err)
	}

	return nil
}

func (s *SSHClient) Connect(
	in io.Reader,
	out io.Writer,
	errOut io.Writer,
) error {
	session, err := s.client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to start ssh session: %w", err)
	}

	defer session.Close()

	session.Stdin = in
	session.Stdout = out
	session.Stderr = errOut

	var fd int
	f, ok := session.Stdin.(*os.File)
	if !ok {
		fd = 0
	} else {
		fd = int(f.Fd())
	}

	w, h, err := term.GetSize(fd)
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	originalState, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("failed to put terminal into raw mode: %w", err)
	}

	defer term.Restore(fd, originalState)

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.ECHOCTL:       0,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty(terminal, h, w, modes); err != nil {
		return fmt.Errorf("failed to get remote pty: %w", err)
	}

	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to get shell: %w", err)
	}

	if err := session.Wait(); err != nil {
		return fmt.Errorf("failed to wait and exit: %w", err)
	}

	return nil
}

func main() {
	sshConfig := ssh.ClientConfig{
		HostKeyCallback: func(
			hostname string,
			remote net.Addr,
			key ssh.PublicKey,
		) error {
			return nil
		},
	}

	conn, err := ssh.Dial("tcp", "localhost:23234", &sshConfig)
	if err != nil {
		fmt.Printf("failed to dial connection: %s\n", err)
		os.Exit(1)
	}

	sshClient := &SSHClient{conn}

	// file, err := os.Create("tmp.txt")
	// if err != nil {
	// 	fmt.Printf("failed to open file: %s\n", err)
	// 	os.Exit(1)
	// }

	if err := sshClient.Connect(
		os.Stdin,
		os.Stdout,
		// file,
		os.Stderr,
	); err != nil {
		fmt.Printf("failed to run ssh client: %s\n", err)
		os.Exit(1)
	}

}
