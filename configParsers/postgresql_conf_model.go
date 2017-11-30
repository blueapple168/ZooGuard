package configParsers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type PgConf struct {
	FileContents string

	DataDirectory                string
	HbaFile                      string
	IdentFile                    string
	ExternalPidFile              string
	ListenAddresses              string
	Port                         string
	MaxConnections               string
	SuperuserReservedConnections string
	UnixSocketDirectories      string
	UnixSocketGroup            string
	UnixSocketPermissions      string
	Bonjour                    string
	BonjourName                string
	AuthenticationTimeout      string
	Ssl                        string
	SslCiphers                 string
	SslPreferServerCiphers     string
	SslEcdhCurve               string
	SslCertFile                string
	SslKeyFile                 string
	SslCaFile                  string
	SslCrlFile                 string
	PasswordEncryption         string
	DbUserNamespace          string
	RowSecurity              string
	KrbServerKeyfile         string
	KrbCaseinsUsers          string
	TcpKeepalivesIdle        string
	TcpKeepalivesInterval    string
	TcpKeepalivesCount       string
	SharedBuffers            string
	HugePages                string
	TempBuffers              string
	MaxPreparedTransactions  string
	WorkMem                 string
	MaintenanceWorkMem      string
	AutovacuumWorkMem            string
	MaxStackDepth                string
	DynamicSharedMemoryType      string
	TempFileLimit                string
	MaxFilesPerProcess           string
	SharedPreloadLibraries       string
	VacuumCostDelay              string
	VacuumCostPageHit            string
	VacuumCostPageMiss           string
	VacuumCostPageDirty          string
	VacuumCostLimit              string
	BgwriterDelay                string
	BgwriterLruMaxpages          string
	BgwriterLruMultiplier        string
	EffectiveIoConcurrency       string
	MaxWorkerProcesses           string
	SharedQueues                 string
	SharedQueueSize              string
	WalLevel                    string
	Fsync                        string
	SynchronousCommit            string
	WalSyncMethod                string
	FullPageWrites               string
	WalCompression               string
	WalLogHints                  string
	WalBuffers                   string
	WalWriterDelay               string
	CommitDelay                  string
	CommitSiblings               string
	CheckpointTimeout            string
	MaxWalSize                   string
	MinWalSize                   string
	CheckpointCompletionTarget   string
	CheckpointWarning           string
	ArchiveMode                 string
	ArchiveCommand              string
	ArchiveTimeout            string
	MaxWalSenders             string
	WalKeepSegments           string
	WalSenderTimeout          string
	MaxReplicationSlots       string
	TrackCommitTimestamp      string
	SynchronousStandbyNames   string
	VacuumDeferCleanupAge     string
	HotStandby                string
	MaxStandbyArchiveDelay    string
	MaxStandbyStreamingDelay  string
	WalReceiverStatusInterval string
	HotStandbyFeedback        string
	WalReceiverTimeout        string
	WalRetrieveRetryInterval  string
	EnableBitmapscan          string
	EnableHashagg             string
	EnableHashjoin            string
	EnableIndexscan           string
	EnableIndexonlyscan       string
	EnableMaterial            string
	EnableMergejoin           string
	EnableNestloop            string
	EnableSeqscan             string
	EnableSort                string
	EnableTidscan             string
	SeqPageCost              string
	RandomPageCost           string
	CpuTupleCost             string
	CpuIndexTupleCost          string
	CpuOperatorCost            string
	NetworkByteCost            string
	RemoteQueryCost            string
	EffectiveCacheSize         string
	Geqo                       string
	GeqoThreshold              string
	GeqoEffort                 string
	GeqoPoolSize               string
	GeqoGenerations            string
	GeqoSelectionBias          string
	GeqoSeed                   string
	DefaultStatisticsTarget    string
	ConstraintExclusion        string
	CursorTupleFraction        string
	FromCollapseLimit          string
	JoinCollapseLimit        string
	LogDestination           string
	LoggingCollector         string
	LogDirectory             string
	LogFilename              string
	LogFileMode              string
	LogTruncateOnRotation    string
	LogRotationAge            string
	LogRotationSize           string
	SyslogFacility            string
	SyslogIdent               string
	EventSource               string
	ClientMinMessages         string
	LogMinMessages            string
	LogMinErrorStatement      string
	LogMinDurationStatement     string
	DebugPrintParse             string
	DebugPrintRewritten         string
	DebugPrintPlan              string
	DebugPrettyPrint            string
	LogCheckpoints              string
	LogConnections               string
	LogDisconnections              string
	LogDuration                    string
	LogErrorVerbosity              string
	LogHostname                     string
	LogLinePrefix                   string
	LogLockWaits                        string
	LogStatement                        string
	LogReplicationCommands              string
	LogTempFiles                        string
	LogTimezone                         string
	ClusterName                         string
	UpdateProcessTitle                  string
	TrackActivities                     string
	TrackCounts                         string
	TrackIoTiming                       string
	TrackFunctions                      string
	TrackActivityQuerySize              string
	StatsTempDirectory                  string
	LogParserStats                      string
	LogPlannerStats                 string
	LogExecutorStats                string
	LogStatementStats                 string
	Autovacuum                        string
	LogAutovacuumMinDuration          string
	AutovacuumMaxWorkers              string
	AutovacuumNaptime                 string
	AutovacuumVacuumThreshold         string
	AutovacuumAnalyzeThreshold        string
	AutovacuumVacuumScaleFactor       string
	AutovacuumAnalyzeScaleFactor      string
	AutovacuumFreezeMaxAge            string
	AutovacuumMultixactFreezeMaxAge   string
	AutovacuumVacuumCostDelay         string
	AutovacuumVacuumCostLimit         string
	SearchPath                    string
	DefaultTablespace             string
	TempTablespaces               string
	CheckFunctionBodies           string
	DefaultTransactionIsolation   string
	DefaultTransactionReadOnly    string
	DefaultTransactionDeferrable  string
	SessionReplicationRole        string
	StatementTimeout               string
	LockTimeout                    string
	VacuumFreezeMinAge             string
	VacuumFreezeTableAge           string
	VacuumMultixactFreezeMinAge    string
	VacuumMultixactFreezeTableAge  string
	ByteaOutput                    string
	Xmlbinary                      string
	Xmloption                      string
	GinFuzzySearchLimit            string
	GinPendingListLimit            string
	Datestyle                      string
	Intervalstyle                  string
	Timezone                       string
	TimezoneAbbreviations           string
	ExtraFloatDigits                string
	ClientEncoding                  string
	LcMessages                      string
	LcMonetary                      string
	LcNumeric                       string
	LcTime                          string
	DefaultTextSearchConfig         string
	DynamicLibraryPath              string
	LocalPreloadLibraries           string
	SessionPreloadLibraries         string
	DeadlockTimeout                 string
	MaxLocksPerTransaction          string
	MaxPredLocksPerTransaction      string
	ArrayNulls                    string
	BackslashQuote                string
	DefaultWithOids               string
	EscapeStringWarning           string
	LoCompatPrivileges            string
	OperatorPrecedenceWarning     string
	QuoteAllIdentifiers           string
	SqlInheritance                string
	StandardConformingStrings     string
	SynchronizeSeqscans           string
	TransformNullEquals           string
	ExitOnError                   string
	RestartAfterCrash             string
	PersistentDatanodeConnections string
	MaxCoordinators               string
	MaxDatanodes                  string
	PgxcNodeName                  string
	GtmBackupBarrier              string
	IncludeDir                    string
	IncludeIfExists               string
	Include                       string
	MaxPoolSize                   string
	PoolConnKeepalive             string
	PoolMaintenanceTimeout        string
	PoolerPort                    string
	GtmHost                       string
	GtmPort                       string
	KvPairs                       map[string]string
}

//SetContents is used to set the FileContents to be used
func (pc *PgConf) SetContents(fileContents string) {

	pc.FileContents = fileContents
}

//Parse parses the file content and populates PGConfig
func (pc *PgConf) Parse() (errs error) {

	if len(pc.FileContents) == 0 {
		return errors.New("No contents set for config")
	}

	reBlank := regexp.MustCompile(`^[ \t]*$`)
	reComment := regexp.MustCompile(`^[ \t]*#`)
	kvRe := regexp.MustCompile(`[ \t]?(?P<key>\S+)[ |\t]*=[ |\t]?(?P<value>\S+)[ |\t]*#?(?:.*)?`)

	retMap := make(map[string]string)

	for _, v := range strings.Split(pc.FileContents, "\n") {

		v = strings.TrimSpace(v)

		if reBlank.MatchString(v) || reComment.MatchString(v) {

			continue
		}

		var ik, iv string

		for _, v := range kvRe.FindAllStringSubmatch(v, -1) {
			for i, vv := range v {
				if i != 0 {

					if kvRe.SubexpNames()[i] == "key" {
						ik = strings.ToLower(vv)
					}

					if kvRe.SubexpNames()[i] == "value" {
						iv = vv
					}
				}
			}

			retMap[ik] = iv
		}
	}

	pc.KvPairs = retMap

	mapstructure.Decode(retMap, &pc)
	return
}
