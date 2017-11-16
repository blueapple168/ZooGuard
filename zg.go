package main

import
(
//	"github.com/dminGod/ZooGuard/parsers/pgctl"
//	"github.com/dminGod/ZooGuard/cli"
	 "github.com/dminGod/ZooGuard/zg_config"
	"github.com/dminGod/ZooGuard/spoc"
	"net/http"
	"fmt"
// 	_ "github.com/dminGod/ZooGuard/log_collectors"
//	"github.com/kr/pretty"
)


func init(){

	Config = zg_config.GetConfig()
}



var Config zg_config.ZgConfig

func main() {



//	k := pgctl_parser.Pgctl_parser{}
//	k.Init()


//	k.Parse_string(log_collectors.GetPgxcConfig())

//	fmt.Printf("%# v",pretty.Formatter(k))


	// If we have some errors here in parsing, send the user out




//	cli.RenderStatusTable(k)
//	cli.RenderIssuesTable(k)

//	fmt.Println("zg init, before collect")
//	 spoc.Collect()

	fmt.Println("what is the df of 76")
	fmt.Println(spoc.RunCommand("76", `df -h`))
	
	fmt.Println("ping google.com from 77")
	fmt.Println(spoc.RunCommand("77", `lsblk`))



	fmt.Println("collect called, now calling fs webserver")

  	fs := http.FileServer(http.Dir("static_content"))
  	http.Handle("/", fs)

  	fmt.Println("Serving on 3000")
  	http.ListenAndServe(":3000", nil)


}
