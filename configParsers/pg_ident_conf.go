package configParsers

import (
	"regexp"
	"strings"
)

type pgIdentEntry struct {
	Mapname        string
	SystemUsername string
	PgUsername     string
}

//PgIdent is used to store pg_ident.conf values
type PgIdent struct {
	InvalidEntries []string
	Entries        []pgIdentEntry
	ServerName     string
	FileContents   string
}

//SetContents is used to set the FileContents to be used
func (pg *PgIdent) SetContents(fileContents string) {

	pg.FileContents = fileContents
}

//Parse parses the file content and populates PgIdent
func (pg *PgIdent) Parse() {

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

		if len(t) == 3 {

			entry := pgIdentEntry{
				Mapname:        t[0],
				SystemUsername: t[1],
				PgUsername:     t[2],
			}

			pg.Entries = append(pg.Entries, entry)

		} else {

			pg.InvalidEntries = append(pg.InvalidEntries, v)
		}
	}

	return
}
