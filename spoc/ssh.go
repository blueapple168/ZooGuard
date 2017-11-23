package spoc

//"bytes"
//"fmt"

/*func RunSSHCommand(server string, cmd string) (retStr string) {

	var stdoutBuf bytes.Buffer

	if _, ok := Clients[server]; ok {

		fmt.Println("Getting session info and ssh")
		session, err := (Clients[server]).NewSession()

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
}*/
