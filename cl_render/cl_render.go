package cl_render

import (
	"github.com/dminGod/ZooGuard/pgctl_parser"
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
)

// This file will show stuff on the command line

func RenderStatusTable(k pgctl_parser.Pgctl_parser) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)

	data := [][]string{}

	fmt.Println("-- Server Details --")
	table.SetHeader([]string{"Server", "Component", "Component Name", "State", "Issues Count", "Port", "Data Dir"})


	// Add GTM Details
	data = append(data, []string{
		k.Cluster.GtmMaster.GtmMasterServer, "GTM", k.Cluster.GtmMaster.GtmName, "", "", k.Cluster.GtmMaster.GtmMasterPort, k.Cluster.GtmMaster.GtmMasterDir},
	)

	// If GTM has a slave
	if k.Cluster.HasGtmSlave {
		// Add GTM Details
		data = append(data, []string{
			k.Cluster.GtmSlave.GtmSlaveServer, "GTM SL", "", "-", "", k.Cluster.GtmSlave.GtmSlavePort, k.Cluster.GtmSlave.GtmSlaveDir},
		)
	}

	// Coordinators
	for _, v := range k.Cluster.Coord {

		data = append(data, []string{
			v.CoordMasterServer, "Coord", v.CoordName, "", "", v.CoordPort, v.CoordMasterDir},
		)

		// Cooord Slave
		if v.HasSlave {
			data = append(data, []string{
				v.CoordinatorSlave.CoordSlaveServer, "Coord SL", v.CoordName, "", "", v.CoordinatorSlave.CoordSlavePort, v.CoordinatorSlave.CoordSlaveDir},
			)
		}
	}

	// Datanodes
	for _, v := range k.Cluster.Datanodes {

		data = append(data, []string{
			v.DatanodeMasterServer, "Data", v.DatanodeName, "", "", v.DatanodePort, v.DatanodeMasterDir},
		)

		// Datanode Slave
		if v.HasSlave {
			data = append(data, []string{
				v.DatanodeSlave.DatanodeSlaveServer, "Data SL", v.DatanodeName, "", "", v.DatanodeSlave.DatanodeSlavePort, v.DatanodeSlave.DatanodeSlaveDir},
			)
		}
	}

	table.AppendBulk(data)
	table.Render()
}

func RenderIssuesTable(k pgctl_parser.Pgctl_parser) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)

	data := [][]string{}

	fmt.Println("-- Issues Details --")
	table.SetHeader([]string{"Server", "Component", "Component Name", "Issue Code", "Issue Details"})


	// Add GTM Details
	data = append(data, []string{
		k.Cluster.GtmMaster.GtmMasterServer, "GTM", k.Cluster.GtmMaster.GtmName, "Low HDD", "Check the HDD is very lowwwww asd  asdf as s sdf sdf"},
	)

	// If GTM has a slave
	if k.Cluster.HasGtmSlave {
		// Add GTM Details
		data = append(data, []string{
			k.Cluster.GtmSlave.GtmSlaveServer, "GTM SL", "", "-", ""},
		)
	}

	// Coordinators
	for _, v := range k.Cluster.Coord {

		data = append(data, []string{
			v.CoordMasterServer, "Coord", v.CoordName, "", ""},
		)

		// Cooord Slave
		if v.HasSlave {
			data = append(data, []string{
				v.CoordinatorSlave.CoordSlaveServer, "Coord SL", v.CoordName, "", ""},
			)
		}
	}

	// Datanodes
	for _, v := range k.Cluster.Datanodes {

		data = append(data, []string{
			v.DatanodeMasterServer, "Data", v.DatanodeName, "", ""},
		)

		// Datanode Slave
		if v.HasSlave {
			data = append(data, []string{
				v.DatanodeSlave.DatanodeSlaveServer, "Data SL", v.DatanodeName, "", ""},
			)
		}
	}

	table.AppendBulk(data)
	table.Render()
}






