package configParsers

import (
	"regexp"
	"strings"
)

type pgHbaEntry struct {
	Type     string
	Database string
	User     string
	Address  string
	Method   string
}

// PgHba is used to store pg_hba.conf values and to be able to
// use the methods from other packages
type PgHba struct {
	Invalid_Entries []string
	Entries         []pgHbaEntry
	Server_name     string
	FileContents    string
}

//SetContents is used to set the FileContents to be used
func (pg *PgHba) SetContents(fileContents string) {

	pg.FileContents = fileContents
}

//Parse parses the file content and populates PgHba
func (pg *PgHba) Parse() {

	if len(pg.FileContents) == 0 {

		return
	}

	reBlank := regexp.MustCompile(`^[ \t]*$`)
	reComment := regexp.MustCompile(`^[ \t]*#`)

	splitRe := regexp.MustCompile(`[ \t]+`)

	for _, v := range strings.Split(pg.FileContents, "\n") {

		v = strings.TrimSpace(v)

		if reBlank.MatchString(v) || reComment.MatchString(v) {

			continue
		}

		t := splitRe.Split(v, -1)

		if len(t) == 5 {

			entry := pgHbaEntry{
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
