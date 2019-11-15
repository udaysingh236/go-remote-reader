package utils

import (
	"time"

	"golang.org/x/crypto/ssh"
)

// Client is our ssh client struct
type Client struct {
	// the host to connect to
	Host string

	// the client config to use
	ClientConfig *ssh.ClientConfig

	// stores the SSH session while the connection is running
	Session *ssh.Session

	// stores the SSH connection itself in order to close it after transfer
	Conn ssh.Conn

	// the clients waits for the given timeout until given up the connection
	Timeout time.Duration

	// *ssh.client for sftp
	Sshclient *ssh.Client
}

// NewClient Returns a new scp.Client with provided host and ssh.clientConfig
// It has a default timeout of one minute.
func NewClient(host string, config *ssh.ClientConfig) Client {
	return NewConfigurer(host, config).Create()
}

// Connect to the remote SSH server, returns error if it couldn't establish a session to the SSH server
func (a *Client) Connect() error {
	client, err := ssh.Dial("tcp", a.Host, a.ClientConfig)
	if err != nil {
		return err
	}

	a.Sshclient = client
	a.Conn = client.Conn
	a.Session, err = client.NewSession()
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection
func Close(a *Client) {
	a.Session.Close()
	a.Conn.Close()
	a.Sshclient.Close()
}
