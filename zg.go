package main

import (
	"github.com/dminGod/ZooGuard/config"
	"github.com/dminGod/ZooGuard/pgctl_parser"
	"fmt"
)

var AppConf config.ZgConfig

func main() {


	AppConf := config.GetConfig()

	// fmt.Println(AppConf.Zgconf.Pgxcctl_conf_file);

	k := pgctl_parser.Pgctl_parser{FileLocation: AppConf.Zgconf.Pgxcctl_conf_file}

	k.Prase()


//	fmt.Println(k.RawConfig)

	for k, _ := range k.RawConfig {

		fmt.Println( k )
	}




}
