package config_prasers

import (
	"regexp"
	"strings"
)

type pg_hba_entry struct {
	Type     string
	Database string
	User     string
	Address  string
	Method   string
}

type pg_hba struct {
	Invalid_Entries []string
	Entries         []pg_hba_entry
	Server_name     string
	File_contents   string
}

func (pg *pg_hba) parse() {

	if len(pg.File_contents) == 0 {

		return
	}

	re_blank := regexp.MustCompile(`^[ \t]*$`)
	re_comment := regexp.MustCompile(`^[ \t]*#`)

	split_re := regexp.MustCompile(`[ \t]+`)

	for _, v := range strings.Split(pg.File_contents, "\n") {

		v = strings.TrimSpace(v)

		if re_blank.MatchString(v) || re_comment.MatchString(v) {

			continue
		}

		t := split_re.Split(v, -1)

		if len(t) == 5 {

			entry := pg_hba_entry{
				Type:     t[0],
				Database: t[1],
				User:     t[2],
				Address:  t[3],
				Method:   t[4],
			}

			pg.Entries = append(pg.Entries, entry)

		} else {

			pg.Invalid_Entries = append(pg.Invalid_Entries, v)
		}
	}

	return
}
