package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	utils "./utils"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
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
	dstFile, err := os.Create("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// open source file
	srcFile, err := sftpClient.Open("/home/opc/check-yaml/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Read the concents of remote file into bytes
	srcRead, err := ioutil.ReadAll(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", srcRead)

	var body interface{}
	if err := yaml.Unmarshal([]byte(srcRead), &body); err != nil {
		panic(err)
	}

	body = utils.ConvertYamlToJSON(body)

	fmt.Println("body: ", body)

	// mapBody, _ := body.(map[string]interface{})
	// body = mapBody[parsingKey]

	convertedJSON, err := json.Marshal(body)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Output: %s\n", convertedJSON)
	}

	dstFile.Write(convertedJSON)
	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	dstFile.Close()
	// Close all the connection in the end
	defer utils.Close(&sshClient)
	defer sftpClient.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
