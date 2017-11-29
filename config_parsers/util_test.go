package configParsers

import (
	"testing"
)

func TestRetString(t *testing.T) {

	retstr := []struct {
		str    string
		result string
	}{{"(helloo  ()", "helloo"}}

	for _, v := range retstr {

		res := retString(v.str)
		if v.result != res {
			t.Errorf("Unexpexcted string returned for %v, expected result: %v, resultant string: %v", v.str, v.result, res)

		}
	}

}

func TestRetInt(t *testing.T) {

	reti := []struct {
		str string
		res int
	}{{"100", 100}, {}}

	for _, v := range reti {

		i := retInt(v.str)
		if i != v.res {
			t.Errorf("Unexpected output for %v, expected int instead of string", v.str)
		}
	}

}

func TestRetStringArr(t *testing.T) {

	retarr := []struct {
		str    string
		result []string
	}{{"( ( hi hello how    ( ", []string{"hi", "hello", "how"}}, {" (Welcome to    my ) ", []string{"Welcome", "to", "my"}}}

	for _, v := range retarr {

		arr := retStringArr(v.str)

		for ii, vv := range arr {

			if vv != v.result[ii] {
				t.Errorf("Unexpected string array returned")
			}
		}
	}

}

func TestRetBool(t *testing.T) {

	retb := []struct {
		str    string
		result bool
	}{{"  YES", true}, {"yEs", true}, {"(y  )", true}, {"Y", true}}

	for _, v := range retb {

		b := retBool(v.str)
		if b != v.result {
			t.Errorf("Unexpexted result for %v, expected boolean: %v, resultant boolean: %v", v.str, v.result, b)
		}
	}

}

func TestCleanSpaces(t *testing.T) {

	retstr := []struct {
		str    string
		result string
	}{{"(helloo  ()", "helloo"}}

	for _, v := range retstr {

		res := cleanSpaces(v.str)
		if v.result != res {
			t.Errorf("Unexpexcted string returned for %v, expected result: %v, resultant string: %v", v.str, v.result, res)

		}
	}

}

func TestCleanSpacesArr(t *testing.T) {

	retarr := []struct {
		str    string
		result []string
	}{{"( ( hi hello how    ( ", []string{"hi", "hello", "how"}}, {" (Welcome to    my ) ", []string{"Welcome", "to", "my"}}}

	for _, v := range retarr {

		arr := cleanSpacesArr(v.str)

		for ii, vv := range arr {

			if vv != v.result[ii] {
				t.Errorf("Unexpected string array returned")
			}
		}
	}

}

func TestCleanSpacesBool(t *testing.T) {

	retb := []struct {
		str    string
		result bool
	}{{"  YES", true}, {"yEs", true}, {"(y  )", true}, {"Y", true}}

	for _, v := range retb {

		b := retBool(v.str)
		if b != v.result {
			t.Errorf("Unexpexted result for %v, expected boolean: %v, resultant boolean: %v", v.str, v.result, b)
		}
	}

}

func TestCleanSpacesInt(t *testing.T) {

	reti := []struct {
		str string
		res int
	}{{"100", 100}, {}}

	for _, v := range reti {

		i := retInt(v.str)
		if i != v.res {
			t.Errorf("Unexpected output for %v, expected int instead of string", v.str)
		}
	}

}

func TestCountsMatch(t *testing.T) {

}
