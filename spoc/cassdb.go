package spoc

import (
	//"errors"
	"fmt"
	"time"

	"github.com/dminGod/ZooGuard/zgConfig"
	"github.com/gocql/gocql"
)

//CassDB is used to store configuration details of cassandra database
type CassDB struct {
	Host                  []string
	UID                   string
	Pass                  string
	DBName                string
	DBType                string
	NumConnectionsPerHost int
	ConnectionTimeOut     int
	SocketKeepAlive       int
	NumberOfQueryRetries  int
	ReadConsistency       int
	WriteConsistency      int

	Session *gocql.Session
}

//CassConns has information regarding the various cassabdra databases connected
type CassConns struct {
	Connections []*CassDB
}

func connectCassandra(v zgConfig.Database) {
	var cassdb CassDB
	cassdb.Host = v.Host
	cassdb.UID = v.Username
	cassdb.Pass = v.Password
	cassdb.WriteConsistency = v.CassandraWriteConsistency

	// gocql.NumConnctions = v.CassandraNumConnectionsPerHost

	cluster := gocql.NewCluster(cassdb.Host...)
	cluster.Keyspace = "system"
	cluster.ProtoVersion = 3
	cluster.Timeout = time.Duration(v.CassandraConnectionTimeout) * time.Second
	cluster.SocketKeepalive = time.Duration(v.CassandraSocketKeepAlive) * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: v.CassandraNumberOfQueryRetries}

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassdb.UID,
		Password: cassdb.Pass,
	}

	var err error

	cassdb.Session, err = cluster.CreateSession()

	//logger.Info("Cassandra Configuration", cassandraSession)
	//fmt.Println(cassdb.Session)

	if err != nil {

		//logger.Error( "ErrorType : INFRA_ERROR - Cassandra Connection could not be established, please check!")
		fmt.Println("err")
	} else {
		fmt.Println("Connection succesfull")
		CassConnections.Connections = append(CassConnections.Connections, &cassdb)
	}

}

//Query is used to run Select query on Cassandra database
func (c_db *CassDB) Query(s string) (retVal []map[string]interface{}) {

	iter := c_db.Session.Query(s).Consistency(gocql.LocalOne).Iter()
	result, err := iter.SliceMap()

	if err != nil {
		fmt.Println("Error fetching the details", err)

	} else {
		retVal = result
		fmt.Println(result)

	}
	return

}

//Execute is used to run Insert query on Cassandra database
func (c_db *CassDB) Execute(s string) {

	insertConsistency := c_db.WriteConsistency
	insertConsistencyCass := gocql.Two

	if insertConsistency == 1 {
		insertConsistencyCass = gocql.One
	} else if insertConsistency == 2 {
		insertConsistencyCass = gocql.Two
	} else if insertConsistency == 3 {
		insertConsistencyCass = gocql.Three
	} else {
		insertConsistencyCass = gocql.One
	}

	err := c_db.Session.Query(s).Consistency(insertConsistencyCass).Exec()

	if err != nil {

		fmt.Println("Insert Query failed", err)
	} else {
		fmt.Println("Insert successful")
	}

}
