package configParsers

import (
	"regexp"
	"strings"
)

type pg_ident_entry struct {
	Mapname        string
	SystemUsername string
	PgUsername     string
}

//Pg_ident is used to store pg_ident.conf values
type Pg_ident struct {
	Invalid_Entries []string
	Entries         []pg_ident_entry
	Server_name     string
	File_contents   string
}

//Set_contents is used to set the File_contents to be used
func (pg *Pg_ident) Set_contents(file_contents string) {

	pg.File_contents = file_contents
}

//Parse parses the file content and populates Pg_ident
func (pg *Pg_ident) Parse() {

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

		if len(t) == 3 {

			entry := pg_ident_entry{
				Mapname:        t[0],
				SystemUsername: t[1],
				PgUsername:     t[2],
			}

			pg.Entries = append(pg.Entries, entry)

		} else {

			pg.Invalid_Entries = append(pg.Invalid_Entries, v)
		}
	}

	return
}
