package zg_config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type http struct {
	HttpEnabled bool
	HttpPort    int
}

type Remote_server struct {
	Ip_or_host        string
	Connection_method string // This will be password / key

	Username string
	Password string // Will be used if the way to connect is with password

	Key_file_location string //

}

type zg_configuration struct {

	// How to fetch the pgxc_ctl file?
	Ctl_file_is_remote bool // if

	// Remote pgxc_ctl file settings

	Ctl_file_local_path string
}

type Server struct {
	Server_name  string
	Server_ip    string
	Environment  string
	Ssh_user     string
	Ssh_password string
}

type App struct {
	App_type       string
	App_unique_key string

	Config_folder string
	Config_file   string

	Component_name   string
	Installed_server string
	App_port         string
	Grpc_port        string
	Http_port        string

	Cluster_servers []string
	Seed_servers    []string

	Pgxc_ctl_server string
}

type Database struct {
	Username                        string
	Password                        string
	DatabaseName                    string
	DatabaseType                    string
	Host                            []string
	Parent_App_Name                 string
	Linked_Component                string
	Parent_Type                     string
	Component_Role                  string
	Db_Identity                     string
	Cassandra_NumConnectionsPerHost int
	Cassandra_ConnectionTimeout     int
	Cassandra_SocketKeepAlive       int
	Cassandra_NumberOfQueryRetries  int
	Cassandra_ReadConsistency       int
	Cassandra_WriteConsistency      int
}

type ZgConfig struct {
	Zgconf   zg_configuration
	Http     http
	Servers  []Server
	Apps     []App
	Database []Database
}

// Dont need to add the .toml in the name here
var configFile string = "zg.toml"

// With trailing slash
var linuxConfigFolders []string = []string{"/etc/"}
var windowsConfigFolders []string = []string{"\\zg\\"}
var ViConfig *viper.Viper
var configLoaded bool

var Config ZgConfig

/*
	This module handles the application configuration
	configuration can be passed from the configuration file
	also can be set using flags.
	location of the configuration file will either be a standard location or will using command line flags.
*/

func init() {

	GetConfiguration()

}

func GetConfiguration() {

	configFile, err := getConfigFile()

	if err != nil {

		fmt.Println("There were errors during fetching application config")
		for _, e := range err {

			fmt.Println(e.Error())
		}

		os.Exit(1)
	}

	ViConfig = viper.New()

	ViConfig.SetConfigFile(configFile)
	ViConfig.AutomaticEnv()

	verr := ViConfig.ReadInConfig()

	e2 := ViConfig.Unmarshal(&Config)

	if e2 != nil {

		fmt.Print("Error marshaling config ", e2)
	}

	if verr != nil {

		fmt.Println("There was an error reading in configuration. Error : ", verr.Error())
	}

	configLoaded = true

}

func GetConfig() ZgConfig {

	if configLoaded == true {

		return Config
	} else {
		GetConfiguration()
	}

	return Config

	/*
		configFile, err := getConfigFile()

		if err != nil {

			fmt.Println("There were errors during fetching application config")
			for _, e := range err {

				fmt.Println(e.Error())
			}

			os.Exit(1)
		}

		ViConfig = viper.New()

		ViConfig.SetConfigFile(configFile)
		ViConfig.AutomaticEnv()

		verr := ViConfig.ReadInConfig()

		e2 := ViConfig.Unmarshal(&Config)

		if e2 != nil {

			fmt.Print("Error marshaling config ", e2)
		}

		if verr != nil {

			fmt.Println("There was an error reading in configuration. Error : ", verr.Error())
		}

		configLoaded = true
	*/
}

func ShowConfig() {

	fmt.Println(ViConfig.AllSettings())

}

/*
	Check in common default locations if the config file is available
	Check if there is an environment variable set
	Check if there is a config file passed
	Command line flags will override environment variables and common loctions
		-- If passed and file not found then application will exit with 1
	Environment variables will override common locations
		-- If environment variables set and file not found app will exit with 1
	If flags and env variables not found then common locations will be used
	If config file not found in any of the places then app will exit with  code 1
*/
func getConfigFile() (retFilePath string, retErrors []error) {

	var flagConfigFile string
	flag.StringVar(&flagConfigFile, "config_file", "", "Customize the path of the configuration file using this flag.")
	flag.Parse()

	envConfigFile := os.Getenv("ZG_CONFIG_FILE")

	// If we are in windows, check the folders we generally put stuff in
	if runtime.GOOS == "windows" {

		winFolders := getAllWindowsDrives()

		if len(winFolders) == 0 {

			retErrors = append(retErrors, errors.New("Windows, no drives found when searching for config."))
		}

	FileChecking:
		// Get the drive letters, loop over them and check which folder seems to be correct
		for _, curDrive := range getAllWindowsDrives() {

			// Loop over all the common windows folders
			for _, curFolder := range windowsConfigFolders {

				curFile := curDrive + ":" + curFolder + configFile

				if fileExists(curFile) {

					// Config file found
					retFilePath = curFile
					break FileChecking

				}
			}
		}
	}

	// If we are in linux, check the folders for that
	if runtime.GOOS == "linux" {

		for _, curFolder := range linuxConfigFolders {

			curFile := curFolder + configFile

			if fileExists(curFile) {

				retFilePath = curFile
				break
			}
		}
	}

	if len(flagConfigFile) > 0 {

		if fileExists(flagConfigFile) {

			retFilePath = flagConfigFile
		} else {

			retErrors = append(retErrors, errors.New("Unable to locate the config_file file '"+flagConfigFile+"' exiting."))
		}
	} else {

		if len(envConfigFile) > 0 {

			// In windows if you use set DWARA_CONFIG_FILE="C:\..." it adds the " to the string also, removing that
			envConfigFile = strings.Trim(envConfigFile, `"`)

			if fileExists(envConfigFile) {

				retFilePath = envConfigFile
			} else {

				retErrors = append(retErrors, errors.New("Unable to locate the config file that is set on environment variable DWARA_CONFIG_FILE "+envConfigFile+" exiting"))
			}
		}
	}

	// In the end if you still don't have anything you have an error
	if retFilePath == "" {

		retErrors = append(retErrors, errors.New("No configuration files named "+configFile+" found."))
	}

	return
}

func getAllWindowsDrives() (availDrives []string) {

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {

		_, err := os.Open(string(drive) + ":\\")

		if err == nil {
			availDrives = append(availDrives, string(drive))
		}
	}

	return
}

func fileExists(curFile string) (retBool bool) {

	_, err := os.Stat(curFile)

	if err == nil {

		retBool = true
	} else {

		fmt.Println("Error locating file, "+curFile, err.Error())
	}

	return
}
