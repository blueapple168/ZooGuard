package spoc

import (
	"github.com/dminGod/ZooGuard/zg_config"
	"time"
	"strings"
	"fmt"
	"database/sql"
	"math/rand"
)

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
			continue
		} else {

			//logger.Info( "Adding connection to pool, Server : ", dbHost, " Port", dbPort)
			//dbpool = append(dbpool, dbpoolConn)
		}
	}

}


func connectCassandra(v zg_config.Database) {



}
