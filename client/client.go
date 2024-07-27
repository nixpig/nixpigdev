package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/charmbracelet/x/term"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
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

func (s *SSHClient) Run(cmd string, out io.Writer) (int, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return 0, fmt.Errorf("failed to start ssh session: %w", err)
	}

	defer session.Close()
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.ECHOCTL:       0,
		ssh.TTY_OP_OSPEED: 14400,
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	w, h, err := term.GetSize(0)
	if err != nil {
		return 0, fmt.Errorf("failed to get terminal size: %w", err)
	}

	fd := int(os.Stdin.Fd())

	originalState, err := terminal.MakeRaw(fd)
	if err != nil {
		return 0, fmt.Errorf("failed to put terminal into raw mode: %w", err)
	}
	defer terminal.Restore(fd, originalState)

	if err := session.RequestPty("xterm-256color", h, w, modes); err != nil {
		return 0, fmt.Errorf("failed to get remote pty: %w", err)
	}

	if err := session.Shell(); err != nil {
		return 0, fmt.Errorf("failed to get shell: %w", err)
	}

	if err := session.Wait(); err != nil {
		return 0, fmt.Errorf("failed to wait and exit: %w", err)
	}

	// output, err := session.CombinedOutput(cmd)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to get combined output: %w", err)
	// }
	//
	// fmt.Println("output: ", output)

	// p, err := out.Write(output)
	// if err != nil {
	// 	return p, fmt.Errorf("failed to write to out: %w", err)
	// }

	return 0, nil
}

func main() {

	sshConfig := ssh.ClientConfig{
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	conn, err := ssh.Dial("tcp", "localhost:23234", &sshConfig)
	if err != nil {
		fmt.Printf("failed to dial connection: %s\n", err)
		os.Exit(1)
	}

	sshClient := &SSHClient{conn}

	if _, err := sshClient.Run("", os.Stdout); err != nil {
		fmt.Printf("failed to run command: %s\n", err)
		os.Exit(1)
	}

}
