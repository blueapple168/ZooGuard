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

	HasGtmSlave bool
	GtmSlave    gtm_slave
	GtmPxyExtraConfig string

	HasGtmProxy      bool
	GTMProxies  []gtm_proxy

	Coord []coordinator_master

	HasCoordinatorSlaves bool
	CoordSlaves          []coordinator_slave

	Datanodes []datanode_master

	HasDatanodeSlaves bool
	DatanodeSlaves    []datanode_slave

	ServersList		[]string
}

type gtm_master struct {

	// GTM Master
	GtmName                      string
	GtmMasterServer              string
	GtmMasterPort                string
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
	CoordSlaveServer     string
	CoordSlavePoolerPort string
	CoordSlaveSync       bool
	CoordSlaveDir        string
	CoordSlavePort       string
	CoordPgHbaEntrie     string
}

type datanode_master struct {
	DatanodeName         string
	DatanodeMasterServer string
	DatanodePort         string
	DatanodePoolerPort   string
	DatanodeMasterDir    string
	DatanodeMasterWALDir string
	DatanodeMaxWALSender string
	DatanodeArchLogDir   string

	DatanodeExtraConfig         string
	DatanodePgHbaEntry          string
	DatanodeSpecificExtraPgHba  string
	DatanodeSpecificExtraConfig string

	HasSlave      bool
	DatanodeSlave datanode_slave
}

type datanode_slave struct {
	DatanodeSlave           bool
	DatanodeSlaveServer     string
	DatanodeSlaveDir        string
	DatanodeSlavePort       string
	DatanodeSlavePoolerPort string
}
