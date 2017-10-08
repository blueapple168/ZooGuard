package main

import
(
	"github.com/dminGod/ZooGuard/parsers/pgctl"
	"github.com/dminGod/ZooGuard/cli"
	"net/http"
	"fmt"
	"github.com/dminGod/ZooGuard/log_collectors"
	"github.com/kr/pretty"
)


func main() {


	// AppConf := zg_config.GetConfig()

	k := pgctl_parser.Pgctl_parser{}
	k.Init()


	k.Parse_string(log_collectors.GetPgxcConfig())

	fmt.Printf("%# v",pretty.Formatter(k))


	// If we have some errors here in parsing, send the user out




	cli.RenderStatusTable(k)
	cli.RenderIssuesTable(k)

//	log_collectors.Collect()


  fs := http.FileServer(http.Dir("static_content"))
  http.Handle("/", fs)

  fmt.Println("Serving on 3000")
  http.ListenAndServe(":3000", nil)


}