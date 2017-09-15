package pgctl_parser

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)


type Pgctl_parser struct {
	FileLocation string
	RawConfig map[string]string
}


func (p *Pgctl_parser)Prase() {

	file, err := os.Open(p.FileLocation)

	if err != nil {

		fmt.Println("Error while reading file", err)

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	p.RawConfig = make(map[string]string)

	for scanner.Scan() {             // internally, it advances token based on sperator

		p.interpret_line(scanner.Text())
	}
}


func (p *Pgctl_parser) interpret_line(curLine string){

	matBool, _ := regexp.Match("^( +|\t+)?[#-]", []byte(curLine))
	hasEqualTo := strings.Contains(curLine, "=")

	charsLen := len(strings.Replace(strings.Replace(curLine, " ", "", -1), "\t", "", -1))


	if matBool == false && charsLen > 0 && hasEqualTo{

		remHash := strings.Split(curLine, "#")

		kvPair := strings.Split(remHash[0], "=")

		if len(kvPair) > 1 {

			p.RawConfig[kvPair[0]] = kvPair[1]
		}
	}
}





