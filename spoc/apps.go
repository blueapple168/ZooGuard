package spoc

import (
	"fmt"

	"github.com/dminGod/ZooGuard/zgConfig"
)

//AppConfigure is used to store configuration details of the application
type AppConfigure struct {
	ApplicationType      string
	ApplicationUniqueKey string
	Server               string
	AppPort              string
	GRPCPort             string
	HTTPPort             string
	ConfigFile           string
}

//AppConns stores an array of AppConfigure containing details of various applications
type AppConns struct {
	Connections []*AppConfigure
}

func connectApps(v zgConfig.App) {

	var app AppConfigure
	app.ApplicationType = v.ApplicationType
	app.ApplicationUniqueKey = v.ApplicationUniqueKey
	app.Server = v.Server
	app.AppPort = v.NativePort
	app.GRPCPort = v.HectorGrpcPort
	app.HTTPPort = v.HTTPPort
	app.ConfigFile = v.ConfigFile

	AppConnections.Connections = append(AppConnections.Connections, &app)
	fmt.Println("App configured")

}
