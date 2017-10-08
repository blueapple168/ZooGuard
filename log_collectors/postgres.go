package log_collectors

import (

	"golang.org/x/crypto/ssh"
	"bytes"
)

func init() {

	clients = make(map[string]*ssh.Client)

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
	}

	for _, v := range []string{ "10.5.0.10:22", "10.5.0.11:22", "10.5.0.12:22", "10.5.0.21:22", "10.5.0.22:22"} {

		c, err := ssh.Dial("tcp", v, config)
		if err != nil {

		}

		clients[v] = c
	}
}

var clients map[string]*ssh.Client

func Collect() {

	for _, v := range []string{"cat /home/postgres/pgxc_ctl/pgxc_ctl.conf" } {

		for _, vv := range []string{ "10.5.0.10:22" } {

			runCommand(vv, v)
		}
	}
}

func GetPgxcConfig() (retStr string){

	return runCommand("10.5.0.10:22", "cat /home/postgres/pgxc_ctl/pgxc_ctl.conf")
}


func runCommand(server string, cmd string) (retStr string){

	var stdoutBuf bytes.Buffer

	if _, ok := clients[server]; ok {

		session, _ := (clients[server]).NewSession()

		session.Stdout = &stdoutBuf
		session.Run(cmd)
		retStr = stdoutBuf.String()
	}

	return
}

