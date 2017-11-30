package configParsers

import (
	"github.com/dminGod/ZooGuard/spoc"
)

type pgctlStagingConfig struct {

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

type pgxcCluster struct {

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

	GtmMaster gtmMaster

	HasGtmSlave       bool
	GtmSlave          gtmSlave
	GtmPxyExtraConfig string

	HasGtmProxy bool
	GTMProxies  []gtmProxy

	Coord []coordinatorMaster

	HasCoordinatorSlaves bool
	CoordSlaves          []coordinatorSlave

	Datanodes []datanodeMaster

	HasDatanodeSlaves bool
	DatanodeSlaves    []datanodeSlave

	ServersList []string
}

type gtmMaster struct {

	// GTM Master
	GtmName                      string
	GtmMasterServer              string
	GtmMasterPort                string
	GtmMasterDir                 string
	GtmExtraConfig               string
	GtmMasterSpecificExtraConfig string

	HasSlave            bool
	GtmSlave            gtmSlave
	ServerConn          *spoc.ConnInfo
	ServerConfiguration PgConf
}

type gtmSlave struct {

	// GTM Slave
	GtmSlave bool

	GtmSlaveName                string
	GtmSlaveServer              string
	GtmSlavePort                string
	GtmSlaveDir                 string
	GtmSlaveSpecificExtraConfig string
	ServerConn                  *spoc.ConnInfo
	ServerConfiguration         PgConf
}

type gtmProxy struct {
	GtmProxyName   string
	GtmProxyServer string
	GtmProxyPort   string
	GtmProxyDir    string

	ConnectedCoords          []*coordinatorMaster
	ConnectedCoordSlaves     []*coordinatorSlave
	ConnectedDatanodeMasters []*datanodeMaster
	ConnectedDatanodeSlaves  []*datanodeSlave
	ServerConfiguration      PgConf
	ServerConn               *spoc.ConnInfo
}

type coordinatorMaster struct {
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
	ServerConfiguration      PgConf
	IdentConfiguration       PgIdent
	HbaConfiguration         PgHba

	ConfiguredGtmIP   string
	ConfiguredGtmPort string
	ViaGTMProxy       bool

	GtmProxy         *gtmProxy
	HasSlave         bool
	CoordinatorSlave *coordinatorSlave
	ServerConn       *spoc.ConnInfo
}

type coordinatorSlave struct {

	// Coordinator Slaves
	CoordSlaveServer     string
	CoordSlavePoolerPort string
	CoordSlaveSync       bool
	CoordSlaveDir        string
	CoordSlavePort       string
	CoordPgHbaEntrie     string

	ConfiguredGtmIP     string
	ConfiguredGtmPort   string
	GtmProxy            *gtmProxy
	ServerConfiguration PgConf
	IdentConfiguration  PgIdent
	HbaConfiguration    PgHba
	ServerConn          *spoc.ConnInfo
}

type datanodeMaster struct {
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
	DatanodeSlave *datanodeSlave

	ConfiguredGtmIP   string
	ConfiguredGtmPort string

	GtmProxy            *gtmProxy
	ServerConfiguration PgConf
	IdentConfiguration  PgIdent
	HbaConfiguration    PgHba
	ServerConn          *spoc.ConnInfo
	Abc                 string
}

type datanodeSlave struct {
	DatanodeSlave bool
	PgDetails     PGConfig

	DatanodeSlaveServer string
	DatanodeSlaveDir    string
	DatanodeSlavePort   string
	ServerConn          *spoc.ConnInfo
	ConfiguredGtmIP     string
	ConfiguredGtmPort   string
	GtmProxy            *gtmProxy
	ServerConfiguration PgConf
	IdentConfiguration  PgIdent
	HbaConfiguration    PgHba

	DatanodeSlavePoolerPort string
}

//PGConfig is a struct used to store values temporarily, irrespective of the node type
type PGConfig struct {
	ServerIP            string
	PgDir               string
	PgPort              string
	Conn                *spoc.ConnInfo
	GtmIP               string
	GtmPort             string
	GtmProxy            *gtmProxy
	ServerConfiguration PgConf
	IdentConfiguration  PgIdent
	Role                string
}

//PgNode interface provides with the methods to get and set the configuration of all the nodes
type PgNode interface {
	GetPgConfig() PGConfig
	SetPgConfig(PgConf) bool
	SetIdentConfig(PgIdent) bool
	SetHbaConfig(PgHba) bool
}

func (d *datanodeSlave) GetPgConfig() (k PGConfig) {

	k = PGConfig{
		ServerIP: d.DatanodeSlaveServer,
		PgDir:    d.DatanodeSlaveDir,
		PgPort:   d.DatanodeSlavePort,
		Conn:     d.ServerConn,
		GtmIP:    d.ConfiguredGtmIP,
		GtmPort:  d.ConfiguredGtmPort,
		GtmProxy: d.GtmProxy,
		Role:     "Datanode_Slave",
	}

	return

}

func (d *datanodeMaster) GetPgConfig() (k PGConfig) {

	k = PGConfig{
		ServerIP: d.DatanodeMasterServer,
		PgDir:    d.DatanodeMasterDir,
		PgPort:   d.DatanodePort,
		Conn:     d.ServerConn,
		GtmIP:    d.ConfiguredGtmIP,
		GtmPort:  d.ConfiguredGtmPort,
		GtmProxy: d.GtmProxy,
		Role:     "Datanode_Master",
	}
	return
}

func (c *coordinatorMaster) GetPgConfig() (k PGConfig) {

	k = PGConfig{
		ServerIP: c.CoordMasterServer,
		PgDir:    c.CoordMasterDir,
		PgPort:   c.CoordPort,
		Conn:     c.ServerConn,
		GtmIP:    c.ConfiguredGtmIP,
		GtmPort:  c.ConfiguredGtmPort,
		GtmProxy: c.GtmProxy,
		Role:     "Coordinator_Master",
	}
	return
}

func (c *coordinatorSlave) GetPgConfig() (k PGConfig) {

	k = PGConfig{
		ServerIP: c.CoordSlaveServer,
		PgDir:    c.CoordSlaveDir,
		PgPort:   c.CoordSlavePort,
		Conn:     c.ServerConn,
		GtmIP:    c.ConfiguredGtmIP,
		GtmPort:  c.ConfiguredGtmPort,
		GtmProxy: c.GtmProxy,
		Role:     "Coordinator_Slave",
	}
	return
}

func (d *datanodeSlave) SetPgConfig(p PgConf) (retVal bool) {

	d.ServerConfiguration = p
	return
}

func (d *datanodeMaster) SetPgConfig(p PgConf) (retVal bool) {

	d.ServerConfiguration = p
	retVal = false
	return
}

func (c *coordinatorSlave) SetPgConfig(p PgConf) (retVal bool) {

	c.ServerConfiguration = p
	return
}

func (c *coordinatorMaster) SetPgConfig(p PgConf) (retVal bool) {

	c.ServerConfiguration = p
	return
}

func (d *datanodeMaster) SetIdentConfig(pi PgIdent) (retVal bool) {

	d.IdentConfiguration = pi
	return
}

func (d *datanodeSlave) SetIdentConfig(pi PgIdent) (retVal bool) {

	d.IdentConfiguration = pi
	return
}

func (c *coordinatorMaster) SetIdentConfig(pi PgIdent) (retVal bool) {

	c.IdentConfiguration = pi
	return
}

func (c *coordinatorSlave) SetIdentConfig(pi PgIdent) (retVal bool) {

	c.IdentConfiguration = pi
	return
}

func (d *datanodeMaster) SetHbaConfig(ph PgHba) (retVal bool) {

	d.HbaConfiguration = ph
	return
}

func (d *datanodeSlave) SetHbaConfig(ph PgHba) (retVal bool) {

	d.HbaConfiguration = ph
	return
}

func (c *coordinatorMaster) SetHbaConfig(ph PgHba) (retVal bool) {

	c.HbaConfiguration = ph
	return
}

func (c *coordinatorSlave) SetHbaConfig(ph PgHba) (retVal bool) {

	c.HbaConfiguration = ph
	return
}
