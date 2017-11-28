package config_parsers

import (
	"github.com/dminGod/ZooGuard/spoc"
)

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

	HasGtmSlave       bool
	GtmSlave          gtm_slave
	GtmPxyExtraConfig string

	HasGtmProxy bool
	GTMProxies  []gtm_proxy

	Coord []coordinator_master

	HasCoordinatorSlaves bool
	CoordSlaves          []coordinator_slave

	Datanodes []datanode_master

	HasDatanodeSlaves bool
	DatanodeSlaves    []datanode_slave

	ServersList []string
}

type gtm_master struct {

	// GTM Master
	GtmName                      string
	GtmMasterServer              string
	GtmMasterPort                string
	GtmMasterDir                 string
	GtmExtraConfig               string
	GtmMasterSpecificExtraConfig string

	HasSlave            bool
	GtmSlave            gtm_slave
	ServerConn          *spoc.ConnInfo
	ServerConfiguration Pg_conf
}

type gtm_slave struct {

	// GTM Slave
	GtmSlave bool

	GtmSlaveName                string
	GtmSlaveServer              string
	GtmSlavePort                string
	GtmSlaveDir                 string
	GtmSlaveSpecificExtraConfig string
	ServerConn                  *spoc.ConnInfo
	ServerConfiguration         Pg_conf
}

type gtm_proxy struct {
	GtmProxyName   string
	GtmProxyServer string
	GtmProxyPort   string
	GtmProxyDir    string

	ConnectedCoords          []*coordinator_master
	ConnectedCoordSlaves     []*coordinator_slave
	ConnectedDatanodeMasters []*datanode_master
	ConnectedDatanodeSlaves  []*datanode_slave
	ServerConfiguration      Pg_conf
	ServerConn               *spoc.ConnInfo
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
	ServerConfiguration      Pg_conf
	IdentConfiguration       Pg_ident
	HbaConfiguration         Pg_hba

	ConfiguredGtmIp   string
	ConfiguredGtmPort string
	ViaGTMProxy       bool

	Gtm_proxy        *gtm_proxy
	HasSlave         bool
	CoordinatorSlave *coordinator_slave
	ServerConn       *spoc.ConnInfo
}

type coordinator_slave struct {

	// Coordinator Slaves
	CoordSlaveServer     string
	CoordSlavePoolerPort string
	CoordSlaveSync       bool
	CoordSlaveDir        string
	CoordSlavePort       string
	CoordPgHbaEntrie     string

	ConfiguredGtmIp     string
	ConfiguredGtmPort   string
	Gtm_proxy           *gtm_proxy
	ServerConfiguration Pg_conf
	IdentConfiguration  Pg_ident
	HbaConfiguration    Pg_hba
	ServerConn          *spoc.ConnInfo
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
	DatanodeSlave *datanode_slave

	ConfiguredGtmIp   string
	ConfiguredGtmPort string

	Gtm_proxy           *gtm_proxy
	ServerConfiguration Pg_conf
	IdentConfiguration  Pg_ident
	HbaConfiguration    Pg_hba
	ServerConn          *spoc.ConnInfo
	Abc                 string
}

type datanode_slave struct {
	DatanodeSlave bool
	PgDetails     PgConf

	DatanodeSlaveServer string
	DatanodeSlaveDir    string
	DatanodeSlavePort   string
	ServerConn          *spoc.ConnInfo
	ConfiguredGtmIp     string
	ConfiguredGtmPort   string
	Gtm_proxy           *gtm_proxy
	ServerConfiguration Pg_conf
	IdentConfiguration  Pg_ident
	HbaConfiguration    Pg_hba

	DatanodeSlavePoolerPort string
}

type PgConf struct {
	ServerIp            string
	PgDir               string
	PgPort              string
	Conn                *spoc.ConnInfo
	GtmIp               string
	GtmPort             string
	Gtm_proxy           *gtm_proxy
	ServerConfiguration Pg_conf
	IdentConfiguration  Pg_ident
	Role                string
}

type PgNode interface {
	GetPgConfig() PgConf
	SetPgConfig(Pg_conf) bool
	SetIdentConfig(Pg_ident) bool
	SetHbaConfig(Pg_hba) bool
}

func (d *datanode_slave) GetPgConfig() (k PgConf) {

	k = PgConf{
		ServerIp:  d.DatanodeSlaveServer,
		PgDir:     d.DatanodeSlaveDir,
		PgPort:    d.DatanodeSlavePort,
		Conn:      d.ServerConn,
		GtmIp:     d.ConfiguredGtmIp,
		GtmPort:   d.ConfiguredGtmPort,
		Gtm_proxy: d.Gtm_proxy,
		Role:      "Datanode_Slave",
	}

	return

}

func (d *datanode_master) GetPgConfig() (k PgConf) {

	k = PgConf{
		ServerIp:  d.DatanodeMasterServer,
		PgDir:     d.DatanodeMasterDir,
		PgPort:    d.DatanodePort,
		Conn:      d.ServerConn,
		GtmIp:     d.ConfiguredGtmIp,
		GtmPort:   d.ConfiguredGtmPort,
		Gtm_proxy: d.Gtm_proxy,
		Role:      "Datanode_Master",
	}
	return
}

func (c *coordinator_master) GetPgConfig() (k PgConf) {

	k = PgConf{
		ServerIp:  c.CoordMasterServer,
		PgDir:     c.CoordMasterDir,
		PgPort:    c.CoordPort,
		Conn:      c.ServerConn,
		GtmIp:     c.ConfiguredGtmIp,
		GtmPort:   c.ConfiguredGtmPort,
		Gtm_proxy: c.Gtm_proxy,
		Role:      "Coordinator_Master",
	}
	return
}

func (c *coordinator_slave) GetPgConfig() (k PgConf) {

	k = PgConf{
		ServerIp:  c.CoordSlaveServer,
		PgDir:     c.CoordSlaveDir,
		PgPort:    c.CoordSlavePort,
		Conn:      c.ServerConn,
		GtmIp:     c.ConfiguredGtmIp,
		GtmPort:   c.ConfiguredGtmPort,
		Gtm_proxy: c.Gtm_proxy,
		Role:      "Coordinator_Slave",
	}
	return
}

func (d *datanode_slave) SetPgConfig(p Pg_conf) (retVal bool) {

	d.ServerConfiguration = p
	return
}

func (d *datanode_master) SetPgConfig(p Pg_conf) (retVal bool) {

	d.ServerConfiguration = p
	retVal = false
	return
}

func (c *coordinator_slave) SetPgConfig(p Pg_conf) (retVal bool) {

	c.ServerConfiguration = p
	return
}

func (c *coordinator_master) SetPgConfig(p Pg_conf) (retVal bool) {

	c.ServerConfiguration = p
	return
}

func (d *datanode_master) SetIdentConfig(pi Pg_ident) (retVal bool) {

	d.IdentConfiguration = pi
	return
}

func (d *datanode_slave) SetIdentConfig(pi Pg_ident) (retVal bool) {

	d.IdentConfiguration = pi
	return
}

func (c *coordinator_master) SetIdentConfig(pi Pg_ident) (retVal bool) {

	c.IdentConfiguration = pi
	return
}

func (c *coordinator_slave) SetIdentConfig(pi Pg_ident) (retVal bool) {

	c.IdentConfiguration = pi
	return
}

func (d *datanode_master) SetHbaConfig(ph Pg_hba) (retVal bool) {

	d.HbaConfiguration = ph
	return
}

func (d *datanode_slave) SetHbaConfig(ph Pg_hba) (retVal bool) {

	d.HbaConfiguration = ph
	return
}

func (c *coordinator_master) SetHbaConfig(ph Pg_hba) (retVal bool) {

	c.HbaConfiguration = ph
	return
}

func (c *coordinator_slave) SetHbaConfig(ph Pg_hba) (retVal bool) {

	c.HbaConfiguration = ph
	return
}
