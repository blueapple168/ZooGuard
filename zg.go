package main

import (
	"fmt"
	"net/http"

	"github.com/dminGod/ZooGuard/spoc"
)

func main() {

	fmt.Println("what is the df of 76")
	fmt.Println(spoc.RunSSHCommand("76", `hostname -i`))

	fmt.Println("ping google.com from 77")
	fmt.Println(spoc.RunSSHCommand("77", `lsblk`))

	fmt.Println("query on postgres 76")
	spoc.SelectDbCommand()

	fmt.Println("query on postgres 80")
	spoc.SelectCassandraCommand()

	fmt.Println("execute query on postgres  ")
	spoc.ExecuteDbCommand()

	fmt.Println("Execute query on cassandra")
	spoc.ExecuteCassandraCommand()

	fmt.Println("collect called, now calling fs webserver")

	fs := http.FileServer(http.Dir("static_content"))
	http.Handle("/", fs)

	fmt.Println("Serving on 3000")
	http.ListenAndServe(":3000", nil)
}
