package spoc

import (
	"bytes"
	"fmt"
	"net"

	"github.com/dminGod/ZooGuard/zg_config"
	"golang.org/x/crypto/ssh"
)

//ConnInfo contains details regarding SSH connection of the servers
type ConnInfo struct {
	Server_name  string
	Server_ip    string
	AppsInServer []string // D30, D40, Datanode Master, Datanode Slave, GTM, Postgresxl,
	SSH_conn     *ssh.Client
}

//ClientConns has information regarding the SSH connection of all the servers
type ClientConns struct {
	Connections []*ConnInfo
}

//UpdateTag adds the Role of the server as a tag if not already present in the list
func (c *ConnInfo) UpdateTag(tag string) (found bool) {
	fmt.Println(c.Server_name, c.SSH_conn)
	for _, v := range c.AppsInServer {

		if tag == v {
			found = true
		}
	}
	fmt.Println(tag)
	if !found {

		c.AppsInServer = append(c.AppsInServer, tag)
		fmt.Println(c.AppsInServer, c.Server_ip, c.SSH_conn)
		return
	}

	return
}

//UpdateRole is used to update the role of a server
func (c *ClientConns) UpdateRole(ip string, role string) (retBool bool) {
	// connection := ClientConnections.GetServerByIp(ip)
	conn := c.GetServerByIp(ip)

	retBool = conn.UpdateTag(role)
	// connection.UpdateTag(string)

	return
}

//GetServerByIp gets the connection information of the server by it's IP
func (c *ClientConns) GetServerByIp(ip string) (con *ConnInfo) {
	// Loop over all the servers
	// Return the matching server
	ip += ":22"
	fmt.Println("serverbyip", ip)
	for _, v := range c.Connections {

		if ip == v.Server_ip {
			con = v
			fmt.Println("assigned server", con.Server_ip)
			return
		}
	}
	return
}

//GetServerByName gets the connection information of the server by it's server name
func (c *ClientConns) GetServerByName(name string) (con *ConnInfo) {

	for _, v := range c.Connections {

		k := *v

		fmt.Println(k.Server_name)

		if name == k.Server_name {
			con = v
		}
	}
	return

}

var Conf zg_config.ZgConfig

var ClientConnections ClientConns
var CassConnections CassConns
var PostConnections PostConns
var AppConnections AppConns

func init() {

	Conf = zg_config.GetConfig()

	for _, v := range Conf.Servers {

		var conninfo ConnInfo

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
			fmt.Println("Added ip to config", v.Server_ip, v.Server_name)
			conninfo.Server_ip = v.Server_ip
			conninfo.Server_name = v.Server_name
			conninfo.SSH_conn = c

			ClientConnections.Connections = append(ClientConnections.Connections, &conninfo)
		}
	}

	fmt.Printf("Servers %+v", ClientConnections.Connections)

	for _, v := range Conf.Database {

		if v.DatabaseType == "postgresxl" {

			connectPostgres(v)

		} else if v.DatabaseType == "cassandra" {

			connectCassandra(v)
		}
	}

	for _, v := range Conf.Apps {

		connectApps(v)
	}
}

//RunCommand is used to run any command on any  requested server
func (c *ConnInfo) RunCommand(s string) (retStr string) {

	var stdoutBuf bytes.Buffer

	var err error

	fmt.Printf("%+v", c)

	if c.SSH_conn == nil {

		fmt.Println("SSH is nil")
	} else {

		fmt.Println("SSH is not nil")
	}

	k := c.SSH_conn

	if k == nil {

		fmt.Printf("K is nil %v", k)
		return
	} else {

		fmt.Printf("K is not nil %v", k)
	}

	sess, err := k.NewSession()

	if err != nil {
		fmt.Println("Error in running command", err)
	}

	sess.Stdout = &stdoutBuf
	sess.Run(s)
	retStr = stdoutBuf.String()

	return
}

func RunCommand(c *ConnInfo, cmd string) string {

	return c.RunCommand(cmd)
}
