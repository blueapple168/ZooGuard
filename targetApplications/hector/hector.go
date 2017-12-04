package targetApplications

import (
	"fmt"
	"os"

	"github.com/dminGod/ZooGuard/spoc"
	"github.com/spf13/viper"
)

type cassandra struct {
	Host                  []string
	Username              string
	Password              string
	NumConnectionsPerHost int
	ConnectionTimeout     int
	SocketKeepAlive       int
	NumberOfQueryRetries  int
	ReadConsistency       int
	WriteConsistency      int
}

type etcd_configuration struct {
	Enable_etcd                     bool
	Etcd_servers                    []string
	Etcd_connection_url             string
	Etcd_key                        string
	Etcd_heartbeat_key              string
	Etcd_heartbeat_ttl              int
	Etcd_heartbeat_message_interval int
	Etcd_fetch_config_interval      int
}

// hector struct represents the configuration parameters for the hector server
type hector struct {
	ConnectionType         string
	Version                string
	Host                   string
	Port                   string
	Log                    string
	EnableLogTrace         bool
	ListenToOSReloadSignal bool
	LogDirectory           string
	LogFilePrefix          string
	LogFileMaxSize         int
	GrpcHostListen         string
	ApplicationJsonFile    string
	ApplicationJsonPath    string
	StartServersOfType     []string
	RequestMetrics         bool
	QueryMetrics           bool

	// Query Register Feature
	QueryRegisterEnable                   bool
	QueryRegisterNewRequest               bool
	QueryRegisterAllowBumping             bool
	QueryRegisterKeepRegisterAliveDefault int
	QueryRegisterKeepRegisterAliveBumped  int
	QueryRegisterCleanupTicker            int // These tables will not be used for the query register feature

	PortHTTP                        string
	HttpUrlList                     []string
	DefaultRecordsLimit             string
	MaxLimitAllowedByAPI            string
	AsyncProcessRequests            bool
	ManipulateData                  bool
	RunCountQueryForSelect          bool
	ProcessCassandraQueriesFromNode bool
	ProcessCassandraQueriesInterval int
	HistoricalPostgresLimitUpper    int
}

type postgresxl struct {
	Username              string
	Password              string
	Database              string
	Port                  string
	Host                  []string
	MaxOpenConns          int
	MaxIdleConns          int
	ConnMaxLifetime       int
	ThrottleFirstPG       int
	ThrottleFirstPGEnable bool
}

// presto struct represents the configuration parameters for the Presto endpoint
type presto struct {
	ConnectionURL string
}

// Config struct represents the overall configuration comprising of nested cassandra, presto and hector information
type Config struct {
	Etcd       etcd_configuration
	Cassandra  cassandra
	Presto     presto
	Hector     hector
	Postgresxl postgresxl
	loaded     bool
	InstanceID string
	ConfigFile string
}

var Conf Config

func LoadHector() {

	for _, v := range spoc.AppConnections.Connections {
		if v.ApplicationType == "hector" {

			//useFolder := getUseFolder() // "/etc/hector"
			useFolder := v.ConfigFile

			if len(Conf.ConfigFile) == 0 {

				fmt.Println("Using configuration folder "+useFolder, " Searching for file", Conf.InstanceID+"config.toml")
				viper.SetConfigName(Conf.InstanceID + "config") // path to look for the config file in
				viper.AddConfigPath(useFolder)
			} else {

				fmt.Println("Using custom configuration file :", Conf.ConfigFile)
				viper.SetConfigFile(Conf.ConfigFile)
			}

			viper.SetConfigType("toml")
			err := viper.ReadInConfig()

			if err != nil {

				fmt.Println(err.Error())
				os.Exit(1)
			}

			viper.Unmarshal(&Conf)
			fmt.Println("Fetched Configuration from local. Current Config is : ", Conf)

			Conf.loaded = true

		}
	}
}
