package main

import
(
	"github.com/dminGod/ZooGuard/config"
	"github.com/dminGod/ZooGuard/pgctl_parser"
	"github.com/dminGod/ZooGuard/cl_render"
)


func main() {


	AppConf := config.GetConfig()

	k := pgctl_parser.Pgctl_parser{FileLocation: AppConf.Zgconf.Pgxcctl_conf_file}

	k.Prase()

	// If we have some errors here in parsing, send the user out

	cl_render.RenderStatusTable(k)
	cl_render.RenderIssuesTable(k)


}

