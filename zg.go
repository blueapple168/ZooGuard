package main

import (
	"github.com/dminGod/ZooGuard/config"
	"github.com/dminGod/ZooGuard/pgctl_parser"
	"fmt"
)


func main() {


	AppConf := config.GetConfig()


	k := pgctl_parser.Pgctl_parser{FileLocation: AppConf.Zgconf.Pgxcctl_conf_file}

	k.Prase()

	for k, _ := range k.RawConfig {

		fmt.Println( k )
	}
}

