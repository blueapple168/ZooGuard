package config_prasers

import (
	"github.com/mitchellh/mapstructure"
	"regexp"
	"strings"
)

type gtm_config struct {
	File_contents string

	Nodename                   string
	Listen_addresses           string
	Port                       string
	Worker_threads             string
	Gtm_host                   string
	Gtm_port                   string
	Gtm_connect_retry_interval string
	Keepalives_idle            string
	Keepalives_interval        string
	Keepalives_count           string
	Log_file                   string
	Log_min_messages           string
	Kv_pairs                   map[string]string
}

func (gc *gtm_config) parse() {

	re_blank := regexp.MustCompile(`^[ \t]*$`)
	re_comment := regexp.MustCompile(`^[ \t]*#`)

	kv_re := regexp.MustCompile(`[ \t]?(?P<key>\S+)[ |\t]*=[ |\t]?(?P<value>\S+)[ |\t]*#?(?:.*)?`)

	retMap := make(map[string]string)

	for _, v := range strings.Split(gc.File_contents, "\n") {

		v = strings.TrimSpace(v)

		if re_blank.MatchString(v) || re_comment.MatchString(v) {

			continue
		}

		var ik, iv string

		for _, v := range kv_re.FindAllStringSubmatch(v, -1) {
			for i, vv := range v {
				if i != 0 {

					if kv_re.SubexpNames()[i] == "key" {
						ik = strings.ToLower(vv)
					}

					if kv_re.SubexpNames()[i] == "value" {
						iv = vv
					}
				}
			}

			retMap[ik] = iv
		}
	}

	gc.Kv_pairs = retMap
	mapstructure.Decode(retMap, &gc)
}
