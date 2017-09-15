package pgctl_parser

type pgxc_cluster struct {

	// Overall
	PgxcOwner      string
	PgxcUser       string
	TmpDir         string
	LocalTmpDir    string
	DataDirRoot    string
	PgxcInstallDir []string

	// Backup
	ConfigBackup     bool
	ConfigBackupFile string
	ConfigBackupHost string
	ConfigBackupDir  string
	HasGtmProxy      bool

	// GTM Proxy
	GtmProxyDir string

	// Coordinator Master
	CoordMasterDir    string
	CoordArchLogDir   string
	CoordMaxWALsender string

	CoordSlave bool

	// Datanode Master
	DatanodeMasterDir    string
	DatanodeMaxWalSender string
	DatanodeArchLogDir   string
	PrimaryDatanode      string

	// Datanode Slave Settings
	DatanodeSlaveDir string

	GtmMaster gtm_master

	HasGTMSlave bool
	GtmSlave    gtm_slave

	HasGTMProxy bool
	GTMProxies  []gtm_proxy

	Coord []coordinator_master

	HasCoordinatorSlaves bool
	CoordSlaves          []coordinator_slave

	Datanodes []datanode_master

	HasDatanodeSlaves bool
	DatanodeSlaves    []datanode_slave
}

type gtm_master struct {

	// GTM Master
	GtmName                      string
	GtmMasterServer              string
	GtmMasterPort                int
	GtmMasterDir                 string
	GtmExtraConfig               string
	GtmMasterSpecificExtraConfig string

	HasSlave bool
	GtmSlave gtm_slave
}

type gtm_slave struct {

	// GTM Slave
	GtmSlave bool

	GtmSlaveName                string
	GtmSlaveServer              string
	GtmSlavePort                string
	GtmSlaveDir                 string
	GtmSlaveSpecificExtraConfig string
}

type gtm_proxy struct {
	GtmProxyName      string
	GtmProxyServer    string
	GtmProxyPort      string
	GtmProxyDir       string
	GtmPxyExtraConfig string
}

type coordinator_master struct {
	CoordName         string
	CoordMasterServer string
	CoordPort         string
	PoolerPort        string
	CoordMasterDir    string
	CoordMaxWALSender string

	CoordArchLogDir string

	CoordSpecificExtraPgHba  string
	CoordSpecificExtraConfig string
	CoordExtraConfig         string

	HasSlave         bool
	CoordinatorSlave coordinator_slave
}

type coordinator_slave struct {

	// Coordinator Slaves
	CoordSlaveDir         string
	CoordSlaveServers     string
	CoordSlavePoolerPorts int
	CoordSlaveSync        bool
	CoordSlaveDirs        string
	CoordSlavePorts       int
	CoordPgHbaEntries     string
}

type datanode_master struct {
	DatanodeNames         string
	DatanodeMasterServers string
	DatanodePorts         int
	DatanodePoolerPorts   int
	DatanodeMasterDirs    string
	DatanodeMasterWALDirs string
	DatanodeMaxWALSenders string
	DatanodeArchLogDirs   string

	DatanodeExtraConfig         string
	DatanodePgHbaEntries        string
	DatanodeSpecificExtraPgHba  string
	DatanodeSpecificExtraConfig string

	HasSlave      bool
	DatanodeSlave datanode_slave
}

type datanode_slave struct {
	DatanodeSlave           bool
	DatanodeSlaveServer     string
	DatanodeSlaveDir        string
	DatanodeSlavePort       int
	DatanodeSlavePoolerPort int
}
