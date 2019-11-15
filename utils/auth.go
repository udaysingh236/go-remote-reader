package utils

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

//PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(username string, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	signer, err := ssh.ParsePrivateKey(privateKey)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}

//PrivateKeyWithPassphrase Loads a private and public key from "path" having passphrase and returns a SSH ClientConfig to authenticate with the server
func PrivateKeyWithPassphrase(username string, passpharase []byte, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, passpharase)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}

//PasswordKey uses username and password to authenticate with the server
func PasswordKey(username string, password string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}
