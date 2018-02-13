package spoc

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/dminGod/ZooGuard/zgConfig"
	"golang.org/x/crypto/ssh"
)

//ConnInfo contains details regarding SSH connection of the servers
type ConnInfo struct {
	ServerName   string
	ServerIP     string
	Environment	 string
	AppsInServer []string // D30, D40, Datanode Master, Datanode Slave, GTM, Postgresxl,
	SSHConn      *ssh.Client

	OSVersion    string
	TimeDelta    int64
	ServerIssues []ServerIss

	UlimitDetails map[string]string //  ulimit -a | awk -F '[[:space:]][[:space:]]+|) ' ' { print "\""$1","$3 }  '

	// df -h | awk -F '[[:space:][:space:]]+' ' { print "{size:"$2",\"used:\":"$3",\"available\":"$4 "}" } '
	HddDriveUtilization    []HddUtil
	HddLastChecked         time.Time
	CpuCount               int
	RamAvailable           int // Save the RAM in mb
	LoadAverage            []float64
	CpuAdjustedLoadAverage []float64
}

//ClientConns has information regarding the SSH connection of all the servers
type ClientConns struct {
	Connections []*ConnInfo
}

type HddUtil struct {
	Partition      string
	SpaceAllocated string
	SpaceUsed      string
	PercentageUsed string

	HasIssues bool
}
type ServerIss struct {
	IssueType       string
	AffectedDetails string
	Message         string
	IssueCode       string
}

//UpdateTag adds the Role of the server as a tag if not already present in the list
func (c *ConnInfo) UpdateTag(tag string) (found bool) {
	fmt.Println(c.ServerName, c.SSHConn)
	for _, v := range c.AppsInServer {

		if tag == v {
			found = true
		}
	}
	fmt.Println(tag)
	if !found {

		c.AppsInServer = append(c.AppsInServer, tag)
		//fmt.Println(c.AppsInServer, c.ServerIP, c.SSHConn)
		return
	}

	return
}

//UpdateRole is used to update the role of a server
func (c *ClientConns) UpdateRole(ip string, role string) (retBool bool) {
	// connection := ClientConnections.GetServerByIP(ip)
	conn := c.GetServerByIP(ip)

	retBool = conn.UpdateTag(role)
	// connection.UpdateTag(string)

	return
}

//GetServerByIP gets the connection information of the server by it's IP
func (c *ClientConns) GetServerByIP(ip string) (con *ConnInfo) {
	// Loop over all the servers
	// Return the matching server
	ip += ":22"
	fmt.Println("serverbyip", ip)
	for _, v := range c.Connections {

		if ip == v.ServerIP {
			con = v
			fmt.Println("assigned server", con.ServerIP)
			return
		}
	}
	return
}

//GetServerByName gets the connection information of the server by it's server name
func (c *ClientConns) GetServerByName(name string) (con *ConnInfo) {

	for _, v := range c.Connections {

		k := *v

		//fmt.Println(k.ServerName)

		if name == k.ServerName {
			con = v
		}
	}
	return

}

//Conf variable is used to get the configuration details from the toml file
var Conf zgConfig.ZgConfig

//ClientConnections varibale is used to store configuration details of all servers
var ClientConnections ClientConns

//CassConnections variable is used to store configuration details of all cassandra databases
var CassConnections CassConns

//PostConnections variable is used to store configuration details of all postgres-xl databases
var PostConnections PostConns

//AppConnections variable is used to store configuration details of all applications
var AppConnections AppConns

func init() {

	Conf = zgConfig.GetConfig()

	for _, v := range Conf.Servers {

		var conninfo ConnInfo

		config := &ssh.ClientConfig{
			User: v.SSHUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(v.SSHPassword),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}

		c, err := ssh.Dial("tcp", v.ServerIP, config)

		if err != nil {
			fmt.Println("Error during establishing connection : ", err)

		} else {
			fmt.Println("Added ip to config", v.ServerIP, v.ServerName)
			conninfo.ServerIP = v.ServerIP
			conninfo.ServerName = v.ServerName
			conninfo.SSHConn = c
			conninfo.Environment = v.Environment

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

	//fmt.Printf("%+v", c)

	if c.SSHConn == nil {

		fmt.Println("SSH is nil")
	} else {

		fmt.Println("SSH is not nil")
	}

	k := c.SSHConn

	if k == nil {

		fmt.Printf("K is nil %v", k)
		return
	}

	fmt.Printf("K is not nil %v", k)

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
