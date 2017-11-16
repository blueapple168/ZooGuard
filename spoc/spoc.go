package spoc

import(
        "golang.org/x/crypto/ssh"
        "fmt"
        "bytes"
        "net"
	"github.com/dminGod/ZooGuard/zg_config"
)

func init() {

        clients = make(map[string]*ssh.Client)

	Conf = zg_config.GetConfig()


	for _, v := range Conf.Servers {


        config := &ssh.ClientConfig{
                User: v.Ssh_user,
                Auth: []ssh.AuthMethod{
                        ssh.Password(v.Ssh_password),
                },
                HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
                        return nil
                },
        }

                c, err := ssh.Dial("tcp", v.Server_ip, config)
                
                if err != nil {
                        fmt.Println("Error during establishing connection : ", err)
                } else {
                        fmt.Println("Added ip to config")
                }

                clients[v.Server_name] = c
        }

}

var clients map[string]*ssh.Client

var Conf zg_config.ZgConfig

func Collect() {


        for _, v := range []string{`echo "hellllllooo world"` } {

                fmt.Println("calling the command now..")
                for kk, _ := range clients {

                        s := RunCommand(kk, v)
                        fmt.Println(s)
                }
        }
}



func RunCommand(server string, cmd string) (retStr string){

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

    
