package configParsers

import (
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
)

//GTMConfig stores the configuration details parsed from the gtm config file
type GTMConfig struct {
	FileContents string

	Nodename                string
	ListenAddresses         string
	Port                    string
	WorkerThreads           string
	GtmHost                 string
	GtmPort                 string
	GtmConnectRetryInterval string
	KeepalivesIdle          string
	KeepalivesInterval      string
	KeepalivesCount         string
	LogFile                 string
	LogMinMessages          string
	KvPairs                 map[string]string
}

//SetContents sets the parsed configuration file(string) as the file contents
func (gc *GTMConfig) SetContents(fileContents string) {

	gc.FileContents = fileContents
}

//Parse parses the file contents and prepares the GTMConfig object
func (gc *GTMConfig) Parse() {

	reBlank := regexp.MustCompile(`^[ \t]*$`)
	reComment := regexp.MustCompile(`^[ \t]*#`)

	kvRe := regexp.MustCompile(`[ \t]?(?P<key>\S+)[ |\t]*=[ |\t]?(?P<value>\S+)[ |\t]*#?(?:.*)?`)

	retMap := make(map[string]string)

	for _, v := range strings.Split(gc.FileContents, "\n") {

		v = strings.TrimSpace(v)

		if reBlank.MatchString(v) || reComment.MatchString(v) {

			continue
		}

		var ik, iv string

		for _, v := range kvRe.FindAllStringSubmatch(v, -1) {
			for i, vv := range v {
				if i != 0 {

					if kvRe.SubexpNames()[i] == "key" {
						ik = strings.ToLower(vv)
					}

					if kvRe.SubexpNames()[i] == "value" {
						iv = vv
					}
				}
			}

			retMap[ik] = iv
		}
	}

	gc.KvPairs = retMap
	mapstructure.Decode(retMap, &gc)
}
