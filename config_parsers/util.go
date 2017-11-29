package configParsers

import (
	"strconv"
	"strings"
)

func retString(str string) (retStr string) {

	// Take the string, first remove the space around it
	// Remove any ( and ) brackets around it
	// Anything that remains, remove the space and tabs around it, take only the inside
	// string around it.
	retStr = cleanSpaces(str)

	return
}

func retInt(str string) (retInt int) {

	// Take the input as string, see if this is a valid integer that can be returned
	retInt = cleanSpacesInt(str)

	return
}

func retStringArr(str string) (retStrArr []string) {

	// Take the string, remove the brackets, remove any spaces around it
	// Split it by space or tabs
	// For each element trim the sapces and tabs around it, add it to the string array
	// and return that
	retStrArr = cleanSpacesArr(str)

	return
}

func retBool(str string) (retBool bool) {

	// Take the string if clean the spaces around, brackets around if there are
	// lower(str) if the string is yes -- or y it is true -- if the string is n or no then false
	retBool = cleanSpacesBool(str)

	return
}

func cleanSpaces(str string) string {

	strRet := strings.Replace(str, "(", "", -1)
	strRet = strings.Replace(strRet, ")", "", -1)

	strRet = strings.TrimSpace(strRet)

	return strRet
}

func cleanSpacesArr(str string) (strArrRet []string) {

	str = cleanSpaces(str)
	strArrRet = strings.Fields(str)

	return
}

func cleanSpacesBool(str string) (retBool bool) {

	str = strings.ToLower(cleanSpaces(str))

	if str == "y" || str == "yes" {

		retBool = true
	}

	return
}

func cleanSpacesInt(str string) (retInt int) {

	str = strings.ToLower(cleanSpaces(str))

	retInt, _ = strconv.Atoi(str)

	return
}

func countsMatch(kk ...[]string) (retBool bool) {

	retBool = true
	cnt := len(kk[0])

	for _, v := range kk {

		if len(v) != cnt {

			retBool = false
		}
	}

	return
}
