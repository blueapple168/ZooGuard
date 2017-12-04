package targetApplications

import (
	"fmt"

	"github.com/dminGod/ZooGuard/configParsers"
	"github.com/dminGod/ZooGuard/spoc"
)

// Here we are first pulling the cluster configuration details from pgxc_ctl
// and then we are populating the cluster struct configParsers.PgctlParser -- PgCluster

//PgCluster variable is used to call the *PgctlParser.ParseString method
// to parse the pgxc_ctl.conf file
var PgCluster configParsers.PgctlParser

var PgConf configParsers.PgConf

//LoadCluster gets the information from the pgxc_ctl.conf file
//and loads the clusters of the application
func LoadCluster() {

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

			//fmt.Println("response from server::::", str)
			PgCluster.ParseString(str)

			for i := range PgCluster.Cluster.Datanodes {

				PgNodeDetails(&(PgCluster.Cluster.Datanodes[i]))
				fmt.Printf("Datanode Server Configuration: %+v \n \n Datanode Ident configuration\n %+v \n Datanode HBAConfiguration\n %+v\n", PgCluster.Cluster.Datanodes[i].ServerConfiguration, PgCluster.Cluster.Datanodes[i].IdentConfiguration, PgCluster.Cluster.Datanodes[i].HbaConfiguration)
			}

			for i := range PgCluster.Cluster.Coord {
				PgNodeDetails(&(PgCluster.Cluster.Coord[i]))
				fmt.Printf("Coord server config:\n %+v, Coord Ident Config \n %+v \n Cood HBA config:\n %+v \n", PgCluster.Cluster.Coord[i].ServerConfiguration, PgCluster.Cluster.Coord[i].IdentConfiguration, PgCluster.Cluster.Coord[i].HbaConfiguration)

			}

			for i := range PgCluster.Cluster.DatanodeSlaves {

				PgNodeDetails(&(PgCluster.Cluster.DatanodeSlaves[i]))
				fmt.Println("Datanode slave configuration", PgCluster.Cluster.DatanodeSlaves[i].ServerConfiguration)
			}

			for i := range PgCluster.Cluster.CoordSlaves {

				PgNodeDetails(&(PgCluster.Cluster.CoordSlaves[i]))

			}

			for i := range PgCluster.Cluster.GTMProxies {

				GtNodeDetails(&PgCluster.Cluster.GTMProxies[i])
				fmt.Printf("Printing gtmconf for proxies %+v \n", PgCluster.Cluster.GTMProxies[i].GtmConfiguration)
			}

			GtNodeDetails(&PgCluster.Cluster.GtmMaster)
			fmt.Printf("Printing gtmconf for master %+v", PgCluster.Cluster.GtmMaster.GtmConfiguration)

			/*GtNodeDetails(&PgCluster.Cluster.GtmSlave)
			fmt.Printf("Printing gtmconf for slave %+v", PgCluster.Cluster.GtmSlave.GtmConfiguration)
			*/
		}
	}
}

//PgNodeDetails gets the configuration of all the nodes in the postgres-xl database
func PgNodeDetails(p configParsers.PgNode) {

	k := p.GetPgConfig()

	//postgresql.conf
	var pp configParsers.PgConf
	//fmt.Println("Directory is:", k.PgDir, k.ServerIP)
	cmd := fmt.Sprintf("cat %v/postgresql.conf", k.PgDir)
	s := spoc.RunCommand(k.Conn, cmd)
	pp.SetContents(s)
	pp.Parse()
	p.SetPgConfig(pp)

	//pg_ident.conf
	var pg configParsers.PgIdent
	cmdi := fmt.Sprintf("cat %v/pg_ident.conf", k.PgDir)
	si := spoc.RunCommand(k.Conn, cmdi)
	pg.SetContents(si)
	pg.Parse()
	p.SetIdentConfig(pg)

	//pg_hba.conf
	var ph configParsers.PgHba
	cmdh := fmt.Sprintf("cat %v/pg_hba.conf", k.PgDir)
	sh := spoc.RunCommand(k.Conn, cmdh)
	ph.SetContents(sh)
	ph.Parse()
	p.SetHbaConfig(ph)

}

//GtNodeDetails gets the configuration details of gtm master, slave and proxies
func GtNodeDetails(gt configParsers.GtNode) {

	k := gt.GetGtConfig()
	var gc configParsers.GTMConfig
	var cmd string

	if k.IsProxy {
		cmd = fmt.Sprintf("cat %v/gtm_proxy.conf", k.Dir)
	} else {
		cmd = fmt.Sprintf("cat %v/gtm.conf", k.Dir)
	}

	f := spoc.RunCommand(k.Conn, cmd)
	gc.SetContents(f)
	gc.Parse()
	gt.SetGtConfig(gc)
}
