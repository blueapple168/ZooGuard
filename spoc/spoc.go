package spoc

import (
	"fmt"
	"net"

	"github.com/dminGod/ZooGuard/zg_config"
	"golang.org/x/crypto/ssh"
)

var clients map[string]*ssh.Client
var Conf zg_config.ZgConfig

var CassConnections CassConns
var PostConnections PostConns

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

	for _, v := range Conf.Database {

		if v.DatabaseType == "postgresxl" {

			connectPostgres(v)

		} else if v.DatabaseType == "cassandra" {

			connectCassandra(v)
		}
	}
}
