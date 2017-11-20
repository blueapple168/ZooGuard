package spoc

import (
	//"errors"
	"fmt"
	"time"

	"github.com/dminGod/ZooGuard/zg_config"
	"github.com/gocql/gocql"
)

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

type CassConns struct {
	Connections []CassDB
}

func connectCassandra(v zg_config.Database) {
	var cassdb CassDB
	cassdb.Host = v.Host
	cassdb.UID = v.Username
	cassdb.Pass = v.Password
	cassdb.WriteConsistency = v.Cassandra_WriteConsistency

	// gocql.NumConnctions = v.Cassandra_NumConnectionsPerHost

	cluster := gocql.NewCluster(cassdb.Host...)
	cluster.Keyspace = "system"
	cluster.ProtoVersion = 3
	cluster.Timeout = time.Duration(v.Cassandra_ConnectionTimeout) * time.Second
	cluster.SocketKeepalive = time.Duration(v.Cassandra_SocketKeepAlive) * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: v.Cassandra_NumberOfQueryRetries}

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
		CassConnections.Connections = append(CassConnections.Connections, cassdb)
	}

}

func (c_db *CassDB) Select(s string) {

	iter := c_db.Session.Query(s).Consistency(gocql.LocalOne).Iter()
	result, err := iter.SliceMap()

	if err != nil {
		fmt.Println("Error fetching the details", err)
		return
	} else {
		fmt.Println(result)
	}

}

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
