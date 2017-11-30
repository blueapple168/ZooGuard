package main

import (
	"fmt"

	//"github.com/dminGod/ZooGuard/configParsers"
	"github.com/dminGod/ZooGuard/spoc"
	"github.com/dminGod/ZooGuard/targetApplications/postgres_xl"
	//"github.com/dminGod/ZooGuard/zg_config"
)

func main() {
	//fmt.Println("76 server")
	//s := spoc.RunSSHCommand("76", "hostname -i")
	//fmt.Println(s)
	fmt.Println("After ssh")
	//j := spoc.CassConnections.Connections[0]
	k := spoc.PostConnections.Connections[1]
	//r := k.Query(`select * from local_service_requests_new8 limit 1;`)
	//fmt.Println("Postgres select-----")
	//fmt.Println(r)
	k.Execute(`INSERT INTO running_format ( length_data,update_datetime_data,prefix_data,create_by_data,create_datetime_data,update_by_data,module_key0,format_data) VALUES ( 'AJc', '2013-07-18 02:27:29+0700', 'FJG', 'kjc', '2014-04-05 14:38:58+0700', 'BKg', 'EFg', 'kKK')`)
	//rr := j.Query(`select * from all_trade.local_service_requests_new8 limit 1;`)
	//fmt.Println(rr)
	//j.Execute(`INSERT INTO cassandra_test.local_service_requests_new8 ( local_service_requests_new8_pk,val9_key) VALUES (  1d5568e6-8c77-4401-9d90-641edc7c7cac, {'25102017185122001','25102017185122002','25102017185122003'} ) `)


	/*appc := spoc.AppConnections.Connections[1]
	server := appc.Server
	fmt.Println("server is", server)
	con := fmt.Sprintf("cat %s", appc.ConfigFile)
	fmt.Println(con)
	str := spoc.RunCommand(server, con)
	//_ = str

	var p configParsers.PgctlParser
	fmt.Println("Parse string")
	p.Init()
	fmt.Println("Parse string")
	p.ParseString(str)
	//fmt.Println(str)
	fmt.Println(p.Cluster.GTMProxies)*/
	targetApplications.LoadCluster()

}
