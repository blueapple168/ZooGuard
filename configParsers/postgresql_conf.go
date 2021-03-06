package configParsers

import (
	"bufio"
	"log"
	"os"
)

func ParseTestFile() {

	ff, err := os.Open(`zz_reference_and_scripts/SampleFiles/postgresql.conf`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(ff)

	var p PgConf

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var s string

	for scanner.Scan() {
		s += scanner.Text() + "\n"
	}

	p.SetContents(s)
	p.Parse()

}
