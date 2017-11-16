package log_collectors

import (

	"golang.org/x/crypto/ssh"
	"fmt"
	"bytes"
	"net"
)

func init() {

	clients = make(map[string]*ssh.Client)

	config := &ssh.ClientConfig{
		User: "serveradm",
		Auth: []ssh.AuthMethod{
			ssh.Password("R3dh@t!@#"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	fmt.Println("Init lc")

	for _, v := range []string{  "10.138.32.76:22"} {

		c, err := ssh.Dial("tcp", v, config)
		
		if err != nil {
			fmt.Println("Error during establishing connection : ", err)
		} else {
			fmt.Println("Added ip to config")
		}

		clients[v] = c
	}
}

var clients map[string]*ssh.Client

func Collect() {

	for _, v := range []string{"df -h" } {

		fmt.Println("calling the command now..")
		for _, vv := range []string{ "10.138.32.76:22" } {

			s := runCommand(vv, v)
			fmt.Println(s)
		}
	}
}

func GetPgxcConfig() (retStr string){

	return runCommand("127.0.0.1:22", "cat /tmp/hello.txt")
}


func runCommand(server string, cmd string) (retStr string){

	var stdoutBuf bytes.Buffer

	if _, ok := clients[server]; ok {

		fmt.Println("Getting session info and ssh")
		session, err := (clients[server]).NewSession()

		if err != nil {
			fmt.Println("Error in running command", err)
		}

		session.Stdout = &stdoutBuf
		session.Run(cmd)
		retStr = stdoutBuf.String()
		
	} else {

		fmt.Println("Unable to find server")
	}

	return
}

