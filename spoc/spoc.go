package spoc

import(
        "golang.org/x/crypto/ssh"
        "fmt"
        "bytes"
        "net"
	"time"
	"string"
	"database/sql"
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

	for _, v := range Conf.Database{

	time.Sleep(2 * time.Second)
	//Conf := config.Get()

	dbName := v.Database
	dbUser := v.Username
	dbPass := v.Password

	rand.Seed(time.Now().UTC().UnixNano())

	for _, curDB := range v.Host {


		// Get the host and port
		curDBSplit := strings.Split(curDB, ":")

		if len(curDBSplit) == 1 {

			//logger.Error( "ErrorType : CONFIG_ERROR, Got server configured without port information skipping server, Server:", curDB)
			continue
		}

		dbHost := curDBSplit[0]
		dbPort := curDBSplit[1]

		var dbInfo string

		if len(dbPass) == 0 {

			dbInfo = fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s port=%s",
				dbUser, dbName, dbHost, dbPort)
		} else {

			dbInfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
				dbUser, dbPass, dbName, dbHost, dbPort)
		}

		dbpoolConn, err := sql.Open("postgres", dbInfo)

		//dbpoolConn.SetMaxOpenConns(Conf.Postgresxl.MaxOpenConns)
		//dbpoolConn.SetMaxIdleConns(Conf.Postgresxl.MaxIdleConns)
		//dbpoolConn.SetConnMaxLifetime(time.Duration(Conf.Postgresxl.ConnMaxLifetime) * time.Second)

		if err != nil {

			//logger.Error( "ErrorType : INFRA_ERROR, Not able to connect to DB Server:",  dbHost,"Got error", err.Error())
			continue
		} else {

			//logger.Info( "Adding connection to pool, Server : ", dbHost, " Port", dbPort)
			//dbpool = append(dbpool, dbpoolConn)
		}
	}
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

    
