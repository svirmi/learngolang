package main

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := "lopoi-popoi"
	addr := "lopoi-popoi.myjino.ru"
	password := "Zbf87587f3dd13e4"
	command := "ls -la"

	var output string

	output, err := remoteRun(user, addr, password, command)

	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(output)
}

// https://stackoverflow.com/questions/37679939/how-do-i-execute-a-command-on-a-remote-machine-in-a-golang-cli
// e.g. output, err := remoteRun("root", "MY_IP", "PRIVATE_KEY", "ls")
func remoteRun(user string, addr string, password string, cmd string) (string, error) {
	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring

	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		// Auth: []ssh.AuthMethod{
		// 	ssh.PublicKeys(key),
		// },
		//alternatively, you could use a password
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(cmd)
	return b.String(), err
}
