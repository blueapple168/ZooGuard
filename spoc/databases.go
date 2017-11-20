package spoc

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dminGod/ZooGuard/zg_config"
	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
)

var dbpool []*sql.DB
var cassandraSession *gocql.Session

func connectPostgres(v zg_config.Database) {

	dbName := v.DatabaseName
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
			fmt.Println(dbUser, dbHost, dbPort, dbName)

			dbInfo = fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s port=%s",
				dbUser, dbName, dbHost, dbPort)
		} else {

			dbInfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
				dbUser, dbPass, dbName, dbHost, dbPort)
		}

		dbpoolConn, err := sql.Open("postgres", dbInfo)
		_ = dbpoolConn

		//dbpoolConn.SetMaxOpenConns(Conf.Postgresxl.MaxOpenConns)
		//dbpoolConn.SetMaxIdleConns(Conf.Postgresxl.MaxIdleConns)
		//dbpoolConn.SetConnMaxLifetime(time.Duration(Conf.Postgresxl.ConnMaxLifetime) * time.Second)

		if err != nil {

			//logger.Error( "ErrorType : INFRA_ERROR, Not able to connect to DB Server:",  dbHost,"Got error", err.Error())
			fmt.Println("Unable to connect", err)
			continue
		} else {

			//logger.Info( "Adding connection to pool, Server : ", dbHost, " Port", dbPort)
			dbpool = append(dbpool, dbpoolConn)
			fmt.Println("Added db connection to dbpool")
		}
	}

}

func getConnection() (retcon *sql.DB, err error) {

	if len(dbpool) > 0 {

		retcon = dbpool[0]
	} else {

		err = errors.New("No Connection found in pool")
	}

	return
}

func connectCassandra(v zg_config.Database) {

	cassandraHost := v.Host
	cassandraUID := v.Username
	cassandraPass := v.Password

	// gocql.NumConnctions = v.Cassandra_NumConnectionsPerHost

	cluster := gocql.NewCluster(cassandraHost...)
	cluster.Keyspace = "system"
	cluster.ProtoVersion = 3
	cluster.Timeout = time.Duration(v.Cassandra_ConnectionTimeout) * time.Second
	cluster.SocketKeepalive = time.Duration(v.Cassandra_SocketKeepAlive) * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: v.Cassandra_NumberOfQueryRetries}

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraUID,
		Password: cassandraPass,
	}

	var err error

	cassandraSession, err = cluster.CreateSession()

	//logger.Info("Cassandra Configuration", cassandraSession)
	fmt.Println(cassandraSession)

	if err != nil {

		//logger.Error( "ErrorType : INFRA_ERROR - Cassandra Connection could not be established, please check!")
		fmt.Println("err")
	} else {
		fmt.Println("Connection succesfull")
	}

}

func SelectDbCommand() {

	var prestoResult []map[string]interface{}

	db, err := getConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query(`select * from local_service_requests_new8 limit 1;`)

	if err != nil {
		fmt.Println("QUERY_ERROR, Postgres query failed", err)
		return
	}

	cols, err := rows.Columns()

	if err != nil {
		fmt.Println("QUERY_ERROR, Postgres query failed", err)

		return
	}

	data := make([]interface{}, len(cols))
	args := make([]interface{}, len(data))

	for i := range data {
		args[i] = &data[i]
	}

	for rows.Next() {

		var rowData = make(map[string]interface{})

		if err := rows.Scan(args...); err != nil {
			//logger.Error("ErrorType : QUERY_ERROR, Postgres query failed,  Request ID : ", dbAbstract.RequestID, ", Error when fetching data. Error: ", err.Error())
			fmt.Println("Query failed")
			return
		}

		for i := range data {

			rowData[cols[i]] = data[i]
		}

		prestoResult = append(prestoResult, rowData)

	}

	rows.Close()

	if err != nil {
		fmt.Println("Error in closing rows")

		return
	} else {
		fmt.Println("Success")
		fmt.Println(prestoResult)

	}

}

func ExecuteDbCommand() {

	//connection, err := getConnection()
	db, err := getConnection()

	if err != nil {
		return
	}

	//var error_messages []string

	_, erro := db.Exec(`INSERT INTO running_format ( length_data,update_datetime_data,prefix_data,create_by_data,create_datetime_data,update_by_data,module_key0,format_data) VALUES ( 'AJc', '2013-07-18 02:27:29+0700', 'FJG', 'kjc', '2014-04-05 14:38:58+0700', 'BKg', 'EFg', 'kKK')`)

	if erro != nil {
		fmt.Println("QUERY_ERROR, Postgres query failed")

	} else {
		fmt.Println("Postgres inserted succesfully")

	}

}

func SelectCassandraCommand() {

	iter := cassandraSession.Query(`SELECT * FROM all_trade.local_service_requests_new8 limit 1;`).Consistency(gocql.LocalOne).Iter()
	result, err := iter.SliceMap()

	if err != nil {
		fmt.Println("Error fetching the details", err)
		return
	} else {
		fmt.Println(result)
	}

}

func ExecuteCassandraCommand() {

	for _, v := range Conf.Database {
		if v.DatabaseType == "cassandra" {
			insertConsistency := v.Cassandra_WriteConsistency

			//insertConsistency := Conf.Database
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

			err := cassandraSession.Query(`INSERT INTO all_trade.local_service_requests_new8 ( local_service_requests_new8_pk,val9_key) VALUES (  1d5568e6-8c77-4401-9d90-641edc7c7cac, {'25102017185122001','25102017185122002','25102017185122003'} )`).Consistency(insertConsistencyCass).Exec()

			if err != nil {

				fmt.Println("Insert Query failed", err)
			} else {
				fmt.Println("Insert successful")
			}
		}
	}

}
