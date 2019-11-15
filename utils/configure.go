package utils

import (
	"time"

	"golang.org/x/crypto/ssh"
)

// ClientConfigurer struct containing all the configuration options
type ClientConfigurer struct {
	host         string
	clientConfig *ssh.ClientConfig
	timeout      time.Duration
	remoteBinary string
}

// NewConfigurer Creates a new client configurer.
// It takes the required parameters: the host and the ssh.ClientConfig and
// returns a configurer populated with the default values for the optional
// parameters.
//
// These optional parameters can be set by using the methods provided on the
// ClientConfigurer struct.
func NewConfigurer(host string, config *ssh.ClientConfig) *ClientConfigurer {
	return &ClientConfigurer{
		host:         host,
		clientConfig: config,
		timeout:      time.Minute,
	}
}

// Host Alters the host of the client connects to
func (c *ClientConfigurer) Host(host string) *ClientConfigurer {
	c.host = host
	return c
}

// Timeout Changes the connection timeout.
// Defaults to one minute
func (c *ClientConfigurer) Timeout(timeout time.Duration) *ClientConfigurer {
	c.timeout = timeout
	return c
}

// ClientConfig Alters the ssh.ClientConfig
func (c *ClientConfigurer) ClientConfig(config *ssh.ClientConfig) *ClientConfigurer {
	c.clientConfig = config
	return c
}

// Create Builds a client with the configuration stored within the ClientConfigurer
func (c *ClientConfigurer) Create() Client {
	return Client{
		Host:         c.host,
		ClientConfig: c.clientConfig,
		Timeout:      c.timeout,
	}
}
