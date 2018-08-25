package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

const (
	envURL = "GOVC_URL"
)

var (
	user     string
	host     string
	password string
	port     string
)

func init() {
	url2 := os.Getenv(envURL)
	u, err := url.Parse(url2)
	if err != nil {
		panic(err)
	}

	user = u.User.Username()
	if val, ok := u.User.Password(); ok {
		password = val
	}
	host = u.Host
	port = "22"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//"vim-cmd vmsvc/getallvms"
	commands := os.Args[1:len(os.Args)]
	command := strings.Join(commands, " ")

	//var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(KeyboardInteractiveChallenge(password)),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(command)
	// vim-cmd will exit with `1' when there is no arguments given to it
	//if err := ; err != nil {
	//log.Fatal("Failed to run: " + err.Error())
	//}
	fmt.Println(b.String())
}

func KeyboardInteractiveChallenge(password string) func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	return func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {

		// Log all the provided data:
		//log.Println(`User: ` + user)
		//log.Println(`Instruction: ` + instruction)
		//log.Println(`Questions:`)
		//for q := range questions {
		//log.Println(q)
		//}

		// How many questions are asked?
		countQuestions := len(questions)

		if countQuestions == 1 {

			// We expect that in this case (only one question is asked), that the server want to know the password ;-)
			answers = make([]string, countQuestions, countQuestions)
			answers[0] = password

		} else if countQuestions > 1 {

			// After logging, this call will exit the whole program:
			log.Fatalln(`The SSH server is asking multiple questions! This program cannot handle this case.`)
		}

		err = nil
		return
	}
}
