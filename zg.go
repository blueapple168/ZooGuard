package main

import
(
	"github.com/dminGod/ZooGuard/spoc"
	"net/http"
	"fmt"
)




func main() {

	fmt.Println("what is the df of 76")
	fmt.Println(spoc.RunSSHCommand("76", `hostname -i`))
	
	fmt.Println("ping google.com from 77")
	fmt.Println(spoc.RunSSHCommand("77", `lsblk`))

	fmt.Println("collect called, now calling fs webserver")

  	fs := http.FileServer(http.Dir("static_content"))
  	http.Handle("/", fs)

  	fmt.Println("Serving on 3000")
  	http.ListenAndServe(":3000", nil)
}
