package pgctl_parser

type pgctl_staging_config struct {

	// Overall
	PgxcOwner      string
	PgxcUser       string
	TmpDir         string
	LocalTmpDir    string
	DataDirRoot    string
	PgxcInstallDir string

	// Backup
	ConfigBackup     bool
	ConfigBackupFile string
	ConfigBackupHost string
	ConfigBackupDir  string


	// GTM Master
	GtmName                      string
	GtmMasterServer              string
	GtmMasterPort                string
	GtmMasterDir                 string
	GtmExtraConfig               string
	GtmMasterSpecificExtraConfig string


	// GTM Slave
	GtmSlave                    bool
	GtmSlaveName                string
	GtmSlaveServer              string
	GtmSlavePort                string
	GtmSlaveDir                 string
	GtmSlaveSpecificExtraConfig string



	// GTM Proxy
	GtmProxyDir string

	GtmProxy        bool
	GtmProxyNames   []string
	GtmProxyServers []string
	GtmProxyPorts   []string
	GtmProxyDirs    []string

	GtmPxyExtraConfig []string



	// Coordinator Master
	CoordMasterDir    string
	CoordArchLogDir   string
	CoordMaxWALsender string

	CoordNames         []string
	CoordMasterServers []string
	CoordPorts         []string
	PoolerPorts        []string
	CoordMasterDirs    []string
	CoordMaxWALSenders []string

	CoordArchLogDirs         []string
	CoordSpecificExtraPgHba  []string
	CoordSpecificExtraConfig []string
	CoordExtraConfig         []string



	// Coordinator Slaves
	CoordSlaveDir string

	CoordSlave            bool
	CoordSlaveServers     []string
	CoordSlavePoolerPorts []string
	CoordSlaveSync        bool
	CoordSlaveDirs        []string
	CoordSlavePorts       []string
	CoordPgHbaEntries     []string



	// Datanode Master
	DatanodeMasterDir    string
	DatanodeMaxWalSender string
	DatanodeArchLogDir   string
	PrimaryDatanode      string

	DatanodeNames         []string
	DatanodeMasterServers []string
	DatanodePorts         []string
	DatanodePoolerPorts   []string
	DatanodeMasterDirs    []string
	DatanodeMasterWALDirs []string
	DatanodeMaxWALSenders []string
	DatanodeArchLogDirs   []string

	DatanodeExtraConfig         []string
	DatanodePgHbaEntries        []string
	DatanodeSpecificExtraPgHba  []string
	DatanodeSpecificExtraConfig []string



	// Datanode Slave Settings
	DatanodeSlaveDir string

	DatanodeSlave            bool
	DatanodeSlaveServers     []string
	DatanodeSlavePorts       []string
	DatanodeSlaveDirs        []string
	DatanodeSlavePoolerPorts []string
}
