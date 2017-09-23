package main

import
(
	"github.com/dminGod/ZooGuard/config"
	"github.com/dminGod/ZooGuard/pgctl_parser"
	"github.com/dminGod/ZooGuard/cl_render"
	"net/http"
	"fmt"
)


func main() {


	AppConf := config.GetConfig()

	k := pgctl_parser.Pgctl_parser{  FileLocation: AppConf.Zgconf.Pgxcctl_conf_file }

	k.Prase()

	// If we have some errors here in parsing, send the user out

	cl_render.RenderStatusTable(k)
	cl_render.RenderIssuesTable(k)


 
  fs := http.FileServer(http.Dir("static_content"))
  http.Handle("/", fs)

  fmt.Println("Serving on 3000")
  http.ListenAndServe(":3000", nil)


}