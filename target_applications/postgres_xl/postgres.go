package target_applications

import (
	"fmt"

	"github.com/dminGod/ZooGuard/config_parsers"
	"github.com/dminGod/ZooGuard/spoc"
)

var PgCluster config_parsers.Pgctl_parser
var PgConf config_parsers.PgConf

func Load_cluster() {

	PgCluster.Init()

	for _, v := range spoc.AppConnections.Connections {

		if v.ApplicationType == "postgresxl" {

			t := `( set -o posix ; set) >/tmp/variables.before
				source %v
				(set -o posix ; set) >/tmp/variables.after
				diff /tmp/variables.before /tmp/variables.after
			`

			command := fmt.Sprintf(t, v.ConfigFile)

			c := spoc.ClientConnections.GetServerByName(v.Server)

			if c == nil {
				fmt.Println("Skipping nil connectionn")
				continue
			} else {

				fmt.Println("Going forward")
			}

			str := spoc.RunCommand(c, command)

			fmt.Println("response from server::::", str)
			PgCluster.ParseString(str)

			for i, _ := range PgCluster.Cluster.Datanodes {

				PgNode_details(&(PgCluster.Cluster.Datanodes[i]))
				fmt.Printf("Datanode Server Configuration:\n %+v \n Datanode Ident configuration\n %+v \n Datanode HBAConfiguration\n %+v\n", PgCluster.Cluster.Datanodes[i].ServerConfiguration, PgCluster.Cluster.Datanodes[i].IdentConfiguration, PgCluster.Cluster.Datanodes[i].HbaConfiguration)
			}

			for i, _ := range PgCluster.Cluster.Coord {
				PgNode_details(&(PgCluster.Cluster.Coord[i]))
				//fmt.Println("Coord server config:", PgCluster.Cluster.Coord[i].ServerConfiguration)

			}

			for i, _ := range PgCluster.Cluster.DatanodeSlaves {

				PgNode_details(&(PgCluster.Cluster.DatanodeSlaves[i]))
				//fmt.Println("Datanode slave configuration", PgCluster.Cluster.DatanodeSlaves[i].ServerConfiguration)
			}

			for i, _ := range PgCluster.Cluster.CoordSlaves {

				PgNode_details(&(PgCluster.Cluster.CoordSlaves[i]))

			}

			/*for i, v := range PgCluster.Cluster.GTMProxies {

				cmd := fmt.Sprintf("cat %v/gtm_proxy.conf", v.GtmProxyDir)

				if v.ServerConn != nil {
					fmt.Println("proxy dir:", v.GtmProxyDir)
					fmt.Println(v.ServerConn.Server_ip)
					kk := spoc.RunCommand(v.ServerConn, cmd)
					fmt.Println("printing kk of gtmproxy", kk)

					var pp config_parsers.Pg_conf

					pp.Set_contents(kk)
					pp.Parse()

					PgCluster.Cluster.GTMProxies[i].ServerConfiguration = pp
					//fmt.Printf("Printing pp for proxies %+v \n", pp)
				} else {
					fmt.Println("Server configuration unavailable")
				}

			}

			v := PgCluster.Cluster.GtmMaster

			cmd := fmt.Sprintf("cat %v/gtm.conf", v.GtmMasterDir)

			if v.ServerConn != nil {
				kk := spoc.RunCommand(v.ServerConn, cmd)

				var pp config_parsers.Pg_conf

				pp.Set_contents(kk)
				pp.Parse()

				PgCluster.Cluster.GtmMaster.ServerConfiguration = pp
				fmt.Printf("Printing pp for master %+v \n", PgCluster.Cluster.GtmMaster.ServerConfiguration)
			} else {
				fmt.Println("Server configuration unavailable")
			}

			vv := PgCluster.Cluster.GtmSlave

			cmdi := fmt.Sprintf("cat %v/gtm.conf", vv.GtmSlaveDir)

			if vv.ServerConn != nil {

				kkk := spoc.RunCommand(vv.ServerConn, cmdi)

				var ppp config_parsers.Pg_conf

				ppp.Set_contents(kkk)
				ppp.Parse()

				PgCluster.Cluster.GtmSlave.ServerConfiguration = ppp
				fmt.Printf("printing pp for gtm slave %+v \n", ppp)
			} else {
				fmt.Println("Server configuration not available")
			}*/

		}
	}
}

func PgNode_details(p config_parsers.PgNode) {

	k := p.GetPgConfig()

	//postgresql.conf
	var pp config_parsers.Pg_conf
	//fmt.Println("Directory is:", k.PgDir, k.ServerIp)
	cmd := fmt.Sprintf("cat %v/postgresql.conf", k.PgDir)
	s := spoc.RunCommand(k.Conn, cmd)
	pp.Set_contents(s)
	pp.Parse()
	p.SetPgConfig(pp)

	//pg_ident.conf
	var pg config_parsers.Pg_ident
	cmdi := fmt.Sprintf("cat %v/pg_ident.conf", k.PgDir)
	si := spoc.RunCommand(k.Conn, cmdi)
	pg.Set_contents(si)
	pg.Parse()
	p.SetIdentConfig(pg)

	//pg_hba.conf
	var ph config_parsers.Pg_hba
	cmdh := fmt.Sprintf("cat %v/pg_hba.conf", k.PgDir)
	sh := spoc.RunCommand(k.Conn, cmdh)
	ph.Set_contents(sh)
	ph.Parse()
	p.SetHbaConfig(ph)

}
