package spoc

import (
	"fmt"

	"github.com/dminGod/ZooGuard/zg_config"
)

type AppConfigure struct {
	ApplicationType      string
	ApplicationUniqueKey string
	Server               string
	AppPort              string
	GRPCPort             string
	HTTPPort             string
	ConfigFile           string
}

type AppConns struct {
	Connections []AppConfigure
}

func connectApps(v zg_config.App) {

	var app AppConfigure
	app.ApplicationType = v.Application_type
	app.ApplicationUniqueKey = v.Application_unique_key
	app.Server = v.Server
	app.AppPort = v.Native_port
	app.GRPCPort = v.Hector_Grpc_port
	app.HTTPPort = v.Http_port
	app.ConfigFile = v.Config_file

	AppConnections.Connections = append(AppConnections.Connections, app)
	fmt.Println("App configured")

}
