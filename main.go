package main

import (
	"fmt"
	"io"
	"log"
	"os"

	utils "./utils"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	sshConfig, err := utils.PrivateKey("opc", "/home/uday/.ssh/id_rsa", ssh.InsecureIgnoreHostKey())
	checkError(err)

	sshClient := utils.NewClient("129.213.203.59:22", &sshConfig)

	err = sshClient.Connect()
	checkError(err)

	// create new SFTP client
	sftpClient, err := sftp.NewClient(sshClient.Sshclient)
	checkError(err)

	// create destination file
	dstFile, err := os.Create("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// open source file
	srcFile, err := sftpClient.Open("/home/opc/check-yaml/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// Close all the connection in the end
	defer utils.Close(&sshClient)
	defer sftpClient.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
