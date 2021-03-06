package configParsers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dminGod/ZooGuard/spoc"
)

//PgctlParser is used to store pgxc_ctl.conf values
type PgctlParser struct {
	FileLocation string
	RawConfig    map[string]string
	St           pgctlStagingConfig
	Cluster      pgxcCluster
	PopulateErrs []error
}

//Init is used to initialize RawConfig variable
func (p *PgctlParser) Init() {

	p.RawConfig = make(map[string]string)
}

//Parse parses the file from the FileLocation
func (p *PgctlParser) Parse() {

	// TODO : getting files directly is not correct -- make an abstraction that will return you
	// This guy should not be bothered to get stuff from other systems on the network
	// There's got to be bus layer that will do the work of :

	// 	talk to particular kind of servers and get things for you
	//	send out messages to particular servers or type of servers
	// 	the underlying communication layer could be ssh and command
	//	client running on the side
	// 	running a command on the local server

	// 	1) Get stuff for you
	// 1) Read from a file
	// 2) SSH into a remote server and return the file
	// 3) Get it directly passed as text so the

	// 	2) Send commands over the network for you

	file, err := os.Open(p.FileLocation)
	if err != nil {

		fmt.Println("Error while reading file", err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Fist take all the lines and put them in the RawConfig, so the override happens of the configuration
	// Where its put in multiple times and we have the last most recent value
	for scanner.Scan() { // internally, it advances token based on sperator

		p.interpretLine(scanner.Text())
	}

	// Now that we have a variable to value mapping. Lets start filling our object
	p.Populate()
}

//ParseString is used to parse the content line by line and to format it
func (p *PgctlParser) ParseString(str string) {

	//str = log_collectors.GetPgxcConfig()

	for _, v := range strings.Split(str, "\n") {

		p.interpretLine(v)
		fmt.Println(v)
	}

	p.Populate()
}

// Individual Line parsing, remove the comments, keep the key-value pairs
func (p *PgctlParser) interpretLine(curLine string) {

	matBool, _ := regexp.Match("^( +|\t+)?[#-]", []byte(curLine))
	hasEqualTo := strings.Contains(curLine, "=")

	re := regexp.MustCompile(`\[\d+\]=`)
	re2 := regexp.MustCompile(`^> `)
	re3 := regexp.MustCompile(`"`)

	charsLen := len(strings.Replace(strings.Replace(curLine, " ", "", -1), "\t", "", -1))

	if matBool == false && charsLen > 0 && hasEqualTo {

		remHash := strings.Split(curLine, "#")
		remHash[0] = re.ReplaceAllString(remHash[0], "")
		remHash[0] = re2.ReplaceAllString(remHash[0], "")
		remHash[0] = re3.ReplaceAllString(remHash[0], "")

		fmt.Println("Parsing : ", remHash[0])

		kvPair := strings.Split(remHash[0], "=")

		if len(kvPair) > 1 {

			p.RawConfig[kvPair[0]] = kvPair[1]
		}
	}
}

//Populate module will take the individual lines and start populating objects based on the lines
func (p *PgctlParser) Populate() {

	// First we will just map our values and set them correctly in the flat string:(string/[]string) --
	// Once the mapping on that level is done then we will start putting things in objects.
	// Here we will also do the bulk of validation and give errors related to the cluster

	for k, v := range p.RawConfig {

		switch k {
		case "pgxcOwner":
			p.St.PgxcOwner = retString(v)

		case "pgxcUser":
			p.St.PgxcUser = retString(v)

		case "tmpDir":
			p.St.TmpDir = retString(v)

		case "localTmpDir":
			p.St.LocalTmpDir = retString(v)

		case "dataDirRoot":
			p.St.DataDirRoot = retString(v)

		case "pgxcInstallDir":
			p.St.PgxcInstallDir = retString(v)

		case "configBackup":
			p.St.ConfigBackup = retBool(v)

		case "configBackupFile":
			p.St.ConfigBackupFile = retString(v)

		case "configBackupHost":
			p.St.ConfigBackupHost = retString(v)

		case "configBackupDir":
			p.St.ConfigBackupDir = retString(v)

		case "gtmName":
			p.St.GtmName = retString(v)

		case "gtmMasterServer":
			p.St.GtmMasterServer = retString(v)

		case "gtmMasterPort":
			p.St.GtmMasterPort = retString(v)

		case "gtmMasterDir":
			p.St.GtmMasterDir = retString(v)

		case "gtmExtraConfig":
			p.St.GtmExtraConfig = retString(v)

		case "gtmMasterSpecificExtraConfig":
			p.St.GtmMasterSpecificExtraConfig = retString(v)

		case "gtmSlave":
			p.St.GtmSlave = retBool(v)

		case "gtmSlaveName":
			p.St.GtmSlaveName = retString(v)

		case "gtmSlaveServer":
			p.St.GtmSlaveServer = retString(v)

		case "gtmSlavePort":
			p.St.GtmSlavePort = retString(v)

		case "gtmSlaveDir":
			p.St.GtmSlaveDir = retString(v)

		case "gtmSlaveSpecificExtraConfig":
			p.St.GtmSlaveSpecificExtraConfig = retString(v)

		case "gtmProxyDir":
			p.St.GtmProxyDir = retString(v)

		case "gtmProxy":
			p.St.GtmProxy = retBool(v)

		case "gtmProxyNames":
			p.St.GtmProxyNames = retStringArr(v)

		case "gtmProxyServers":
			p.St.GtmProxyServers = retStringArr(v)

		case "gtmProxyPorts":
			p.St.GtmProxyPorts = retStringArr(v)

		case "gtmProxyDirs":
			p.St.GtmProxyDirs = retStringArr(v)

		case "gtmPxyExtraConfig":
			p.St.GtmPxyExtraConfig = retStringArr(v)

		case "coordMasterDir":
			p.St.CoordMasterDir = retString(v)

		case "coordArchLogDir":
			p.St.CoordArchLogDir = retString(v)

		case "coordMaxWALsender":
			p.St.CoordMaxWALsender = retString(v)

		case "coordNames":
			p.St.CoordNames = retStringArr(v)

		case "coordMasterServers":
			p.St.CoordMasterServers = retStringArr(v)

		case "coordPorts":
			p.St.CoordPorts = retStringArr(v)

		case "poolerPorts":
			p.St.PoolerPorts = retStringArr(v)

		case "coordMasterDirs":
			p.St.CoordMasterDirs = retStringArr(v)

		case "coordMaxWALSenders":
			p.St.CoordMaxWALSenders = retStringArr(v)

		case "coordArchLogDirs":
			p.St.CoordArchLogDirs = retStringArr(v)

		case "coordSpecificExtraPgHba":
			p.St.CoordSpecificExtraPgHba = retStringArr(v)

		case "coordSpecificExtraConfig":
			p.St.CoordSpecificExtraConfig = retStringArr(v)

		case "coordExtraConfig":
			p.St.CoordExtraConfig = retStringArr(v)

		case "coordSlaveDir":
			p.St.CoordSlaveDir = retString(v)

		case "coordSlave":
			p.St.CoordSlave = retBool(v)

		case "coordSlaveServers":
			p.St.CoordSlaveServers = retStringArr(v)

		case "coordSlavePoolerPorts":
			p.St.CoordSlavePoolerPorts = retStringArr(v)

		case "coordSlaveSync":
			p.St.CoordSlaveSync = retBool(v)

		case "coordSlaveDirs":
			p.St.CoordSlaveDirs = retStringArr(v)

		case "coordSlavePorts":
			p.St.CoordSlavePorts = retStringArr(v)

		case "coordPgHbaEntries":
			p.St.CoordPgHbaEntries = retStringArr(v)

		case "datanodeMasterDir":
			p.St.DatanodeMasterDir = retString(v)

		case "datanodeMaxWalSender":
			p.St.DatanodeMaxWalSender = retString(v)

		case "datanodeArchLogDir":
			p.St.DatanodeArchLogDir = retString(v)

		case "primaryDatanode":
			p.St.PrimaryDatanode = retString(v)

		case "datanodeNames":
			p.St.DatanodeNames = retStringArr(v)

		case "datanodeMasterServers":
			p.St.DatanodeMasterServers = retStringArr(v)

		case "datanodePorts":
			p.St.DatanodePorts = retStringArr(v)

		case "datanodePoolerPorts":
			p.St.DatanodePoolerPorts = retStringArr(v)

		case "datanodeMasterDirs":
			p.St.DatanodeMasterDirs = retStringArr(v)

		case "datanodeMasterWALDirs":
			p.St.DatanodeMasterWALDirs = retStringArr(v)

		case "datanodeMaxWALSenders":
			p.St.DatanodeMaxWALSenders = retStringArr(v)

		case "datanodeArchLogDirs":
			p.St.DatanodeArchLogDirs = retStringArr(v)

		case "datanodeExtraConfig":
			p.St.DatanodeExtraConfig = retStringArr(v)

		case "datanodePgHbaEntries":
			p.St.DatanodePgHbaEntries = retStringArr(v)

		case "datanodeSpecificExtraPgHba":
			p.St.DatanodeSpecificExtraPgHba = retStringArr(v)

		case "datanodeSpecificExtraConfig":
			p.St.DatanodeSpecificExtraConfig = retStringArr(v)

		case "datanodeSlaveDir":
			p.St.DatanodeSlaveDir = retString(v)

		case "datanodeSlave":
			p.St.DatanodeSlave = retBool(v)

		case "datanodeSlaveServers":
			p.St.DatanodeSlaveServers = retStringArr(v)

		case "datanodeSlavePorts":
			p.St.DatanodeSlavePorts = retStringArr(v)

		case "datanodeSlaveDirs":
			p.St.DatanodeSlaveDirs = retStringArr(v)

		case "datanodeSlavePoolerPorts":
			p.St.DatanodeSlavePoolerPorts = retStringArr(v)
		}
	}

	// Common Boolean set
	// TODO: Other common variables can be set here as well..
	// TODO: The interpretation of shell set variables!
	p.Cluster.CoordSlave = p.St.CoordSlave
	p.Cluster.HasDatanodeSlaves = p.St.DatanodeSlave
	p.Cluster.HasGtmProxy = p.St.GtmProxy
	p.Cluster.HasGtmSlave = p.St.GtmSlave

	p.MapToObj()
}

//MapToObj maps the values from the staging struct to the
func (p *PgctlParser) MapToObj() {

	allServers := make(map[string]struct{})

	// Check for the counts and make sure they match, else throw out the user...
	coordsOkay := countsMatch(p.St.CoordMasterServers, p.St.CoordMasterDirs, p.St.PoolerPorts, p.St.CoordPorts)
	coordSlOkay := true

	//	dnOkay := countsMatch(p.St.DatanodeMasterServers, p.St.DatanodeMasterDirs, p.St.DatanodeMasterWALDirs)
	dnOkay := countsMatch(p.St.DatanodeMasterServers, p.St.DatanodeMasterDirs)
	dnSlOkay := true

	if p.St.CoordSlave {

		coordSlOkay = countsMatch(p.St.CoordSlaveServers, p.St.CoordSlaveDirs, p.St.CoordSlavePoolerPorts, p.St.CoordSlavePorts)
	}

	if p.St.DatanodeSlave {

		dnSlOkay = countsMatch(p.St.DatanodeSlaveServers, p.St.DatanodeSlavePorts, p.St.DatanodeSlaveDirs, p.St.DatanodeSlavePoolerPorts)
	}

	if coordsOkay == false || coordSlOkay == false || dnOkay == false || dnSlOkay == false {

		p.PopulateErrs = append(p.PopulateErrs, errors.New("There is a mismatch found in pgxc_ctl.conf"))
		fmt.Println("Coord", coordsOkay)
		fmt.Println("Coord Sl", coordSlOkay)
		fmt.Println("DN", dnOkay)
		fmt.Println("DN Sl Coords", dnSlOkay)

		return
	}

	// Map the GTM Slave, if there is one..
	if p.St.GtmSlave {

		p.Cluster.GtmSlave = gtmSlave{
			GtmSlaveServer:              p.St.GtmSlaveServer,
			GtmSlaveDir:                 p.St.GtmSlaveDir,
			GtmSlaveName:                p.St.GtmSlaveName,
			GtmSlavePort:                p.St.GtmSlavePort,
			GtmSlaveSpecificExtraConfig: p.St.GtmSlaveSpecificExtraConfig,
		}

		p.Cluster.GtmSlave.ServerConn = getConnection(p.Cluster.GtmSlave.GtmSlaveServer)
		updateRoles(p.Cluster.GtmSlave.GtmSlaveServer, "gtmSlave")
		allServers[p.St.GtmSlaveServer] = struct{}{}
	}

	// get connection pointer to this server from spoc -- if there then add it
	// to each of the components

	// update the role of the server to add tags
	// spoc.Connectsion.UpdateRoleByIP(GtmMasterServer, 'gtmMaster')

	// Add the GTM Master
	p.Cluster.GtmMaster = gtmMaster{
		GtmName:                      p.St.GtmName,
		GtmMasterServer:              p.St.GtmMasterServer,
		GtmExtraConfig:               p.St.GtmExtraConfig,
		GtmMasterDir:                 p.St.GtmMasterDir,
		GtmMasterPort:                p.St.GtmMasterPort,
		GtmMasterSpecificExtraConfig: p.St.GtmMasterSpecificExtraConfig,
		GtmSlave:                     p.Cluster.GtmSlave,
		HasSlave:                     p.St.GtmSlave,
	}
	updateRoles(p.Cluster.GtmMaster.GtmMasterServer, "gtmMaster")
	p.Cluster.GtmMaster.ServerConn = getConnection(p.Cluster.GtmMaster.GtmMasterServer)

	allServers[p.St.GtmMasterServer] = struct{}{}

	// Put the slaves in the name:object so we can add these as children to the main object
	gp := make(map[string]gtmProxy)
	cs := make(map[int]coordinatorSlave)
	ds := make(map[int]datanodeSlave)

	/*
		First lets add gtm_proxies .. from the pgxc_ctl file how do you know what gtmProxy needs to be configured
		on what server? -- we'll get this info when we start doing the postgresql.conf, but how is this done..?
		TODO: Findout how gtmProxy is assigned to the servers in the pgxc_ctl.conf file

		Lets do the coordinator and datanode slaves first and put them in the map[string]coordinator/datanode


		So when we do the actual coords and datanodes then we will be able to just use its name and see if they
		have a corresponding slave and add that to the server directly.

	*/

	// Has GTM Proxies
	if p.St.GtmProxy == true {

		for i, v := range p.St.GtmProxyServers {

			gp[v] = gtmProxy{
				GtmProxyServer: v,
				GtmProxyName:   p.St.GtmProxyNames[i],
				GtmProxyPort:   p.St.GtmProxyPorts[i],
				GtmProxyDir:    p.St.GtmProxyDirs[i],
				ServerConn:     getConnection(v),
			}

			// We'll add this as an array also, cause not every proxy and slave node may be mapped to something
			p.Cluster.GTMProxies = append(p.Cluster.GTMProxies, gp[v])

			allServers[v] = struct{}{}
		}
	}

	// Has Coord slaves
	if p.St.CoordSlave == true {

		for i, v := range p.St.CoordSlaveServers {

			cs[i] = coordinatorSlave{
				CoordSlaveServer:     v,
				CoordSlavePort:       p.St.CoordSlavePorts[i],
				CoordSlavePoolerPort: p.St.CoordSlavePoolerPorts[i],
				CoordPgHbaEntrie:     p.St.CoordPgHbaEntries[i],
				CoordSlaveDir:        p.St.CoordSlaveDirs[i],
				CoordSlaveSync:       p.St.CoordSlaveSync,
				ServerConn:           getConnection(v),
			}
			//cs[v].ServerConn = getConnection(cs[v].CoordSlaveServer)
			p.Cluster.CoordSlaves = append(p.Cluster.CoordSlaves, cs[i])
			allServers[v] = struct{}{}
		}
	}

	// Has Datanode Slaves
	if p.St.DatanodeSlave == true {

		for i, v := range p.St.DatanodeSlaveServers {

			ds[i] = datanodeSlave{
				DatanodeSlaveServer:     v,
				DatanodeSlavePort:       p.St.DatanodeSlavePorts[i],
				DatanodeSlaveDir:        p.St.DatanodeSlaveDirs[i],
				DatanodeSlavePoolerPort: p.St.DatanodeSlavePoolerPorts[i],
				ServerConn:              getConnection(v),
			}
			//ds[v].ServerConn = getConnection(ds[v].DatanodeSlaveServer)
			p.Cluster.DatanodeSlaves = append(p.Cluster.DatanodeSlaves, ds[i])

			allServers[v] = struct{}{}
		}
	}

	// Loop over Coords
	for i, v := range p.St.CoordMasterServers {

		tmpCrdSlv := coordinatorSlave{}

		// If this coord slaves are set and it exists, it'll get set
		if p.St.CoordSlave {

			if _, ok := cs[i]; ok {

				tmpCrdSlv = cs[i]
			}
		}

		p.Cluster.Coord = append(p.Cluster.Coord, coordinatorMaster{
			CoordName:         p.St.CoordNames[i],
			CoordMasterServer: v,
			//			CoordArchLogDir: p.St.CoordArchLogDirs[i],
			CoordMasterDir:    p.St.CoordMasterDirs[i],
			CoordPort:         p.St.CoordPorts[i],
			PoolerPort:        p.St.PoolerPorts[i],
			CoordMaxWALSender: p.St.CoordMaxWALSenders[i],
			CoordinatorSlave:  &tmpCrdSlv,
			ServerConn:        getConnection(v),
		})

		allServers[v] = struct{}{}
	}

	// Loop over Datanodes
	for i, v := range p.St.DatanodeMasterServers {

		tmpDnSlv := datanodeSlave{}

		if p.St.DatanodeSlave {
			if p.Cluster.HasDatanodeSlaves {
				if _, ok := ds[i]; ok {

					tmpDnSlv = ds[i]
				}
			}
		}

		fmt.Printf("trying to get connection for %v ... %+v \n", v, p.St.DatanodeNames[i])

		p.Cluster.Datanodes = append(p.Cluster.Datanodes, datanodeMaster{
			DatanodeName:         p.St.DatanodeNames[i],
			DatanodeMasterServer: v,
			DatanodeMasterDir:    p.St.DatanodeMasterDirs[i],
			DatanodePort:         p.St.DatanodePorts[i],
			//			DatanodeArchLogDir: p.St.DatanodeArchLogDirs[i],
			DatanodePoolerPort: p.St.DatanodePoolerPorts[i],
			//			DatanodeMasterWALDir: p.St.DatanodeMasterWALDirs[i],
			HasSlave:      p.St.DatanodeSlave,
			DatanodeSlave: &tmpDnSlv,
			ServerConn:    getConnection(v),
		})

		allServers[v] = struct{}{}
	}

	for k := range allServers {

		if len(k) > 0 {
			p.Cluster.ServersList = append(p.Cluster.ServersList, k)
		}
	}

	fmt.Println(allServers)

}

func getConnection(s string) (c *spoc.ConnInfo) {
	c = spoc.ClientConnections.GetServerByIP(s)
	return

}

func updateRoles(ip string, role string) {
	b := spoc.ClientConnections.UpdateRole(ip, role)
	fmt.Println(b)
}
