package targetApplications

import (
	"fmt"
	"regexp"
	"strconv"
	//"time"
	//"strconv"
)

// This is a file that is used to init servers related health and spec information that may be relevant for the application.

// Initial Recce:
// CPU(nproc), Memory, HDD details as shown on df -h
// Uptime of the server
// Time delta from current server (where this binary is running) -- Config level threshold for time difference check (by default we can take it as 1.5 seconds
// Populate the ulimit -a for the logged in user

// Run the HDD usage check()
// Save the load average values for this server

// HDD Usage check:
// df -h -- update the server level object with the latest details.
// Check if there are any values higher than expected for any of the settings.

// - On the configuration level you can set the max disk percentage -- This will then be used when checking for disk
// space issues. Any drives higher than the configured percentage need to be maintained on the server level as a simple
// struct server_level_issues: { issue_type : 'string, affected_details : 'string', message : '' }

// Get load average details for this server
// Actual Load average of 1, 5 and 15 minutes (uptime)  -- CPU Adjusted Load average : the 3 values divided by the number of CPU cores

import (
	"strings"

	"github.com/dminGod/ZooGuard/spoc"
)

func InitializeServer() {

	for _, v := range spoc.ClientConnections.Connections {

		var hd spoc.HddUtil

		cmdHdd := `df -h | awk -F '[[:space:][:space:]]+' ' { print $2","$3","$4","$5 } '`
		s := v.RunCommand(cmdHdd)
		fmt.Printf("df -h:\n %+v", s)

		for _, vv := range strings.Split(s, "\n") {
			//fmt.Printf("string: %+v, len %+v", vv, len(vv))
			/*if i == 0 {
				break
			} else {*/
			if len(vv) > 3 {
				str := strings.Split(vv, ",")
				if str[0] == "Size" {
					continue
				} else {
					if len(str) == 4 {
						hd.Partition = str[0]
						hd.SpaceUsed = str[1]
						hd.SpaceAllocated = str[2]
						hd.PercentageUsed = str[3]

						v.HddDriveUtilization = append(v.HddDriveUtilization, hd)
					}
				}
			}
		}

		fmt.Printf("\nDF- h %+v\n", v.HddDriveUtilization)

		cmdUlimit := `ulimit -a | awk -F '[[:space:]][[:space:]]+|) ' ' { print $1","$3 }'`
		ss := v.RunCommand(cmdUlimit)
		fmt.Printf("ulimit %+v", ss)
		v.UlimitDetails = make(map[string]string)

		for _, kk := range strings.Split(ss, "\n") {
			strr := strings.Split(kk, ",")
			//fmt.Printf("Split string %+v len %v\n", strr, len(strr))
			if len(strr) > 1 {
				v.UlimitDetails[strr[0]] = strr[1]
			}

		}
		fmt.Printf("\nUlimitDetails %v\n", v.UlimitDetails)

		cmdLoadAvg := `uptime`
		sLA := v.RunCommand(cmdLoadAvg)
		fmt.Printf("Load average %v\n", sLA)

		str1 := strings.Split(sLA, ":")
		n := len(str1)
		//fmt.Printf("split string %v len:%v\n", str1, n)

		strrr := strings.Split(str1[(n-1)], ",")
		for _, ff := range strrr {
			fmt.Println(ff)

			re := regexp.MustCompile(`\n`)
			f := strings.TrimSpace(ff)
			fstr := re.ReplaceAllString(f, "")
			fff, err := strconv.ParseFloat(fstr, 64)
			if err == nil {
				v.LoadAverage = append(v.LoadAverage, fff)
			} else {
				fmt.Println(err)
			}

		}

		fmt.Printf("Load Averages: %v", v.LoadAverage)

		cmdCPU := `nproc`
		sc := v.RunCommand(cmdCPU)
		ad := &v.CpuCount
		re2 := regexp.MustCompile(`\n`)
		scc := strings.TrimSpace(sc)
		sc2 := re2.ReplaceAllString(scc, "")
		a, er := strconv.Atoi(sc2)
		if er == nil {
			fmt.Println(a)
			*ad = a
		} else {
			fmt.Println("error", er)
		}

		cmdRAM := `free -m | grep "Mem" | awk '{print $3","$4}'`
		cram := v.RunCommand(cmdRAM)
		sr := strings.Split(cram, ",")
		ramAvail := &v.RamAvailable
		re3 := regexp.MustCompile(`\n`)
		srr := strings.TrimSpace(sr[1])
		sr2 := re3.ReplaceAllString(srr, "")
		srInt, e := strconv.Atoi(sr2)
		if e == nil {
			*ramAvail = srInt
		} else {
			fmt.Println(e)
		}

	}

	for _, b := range spoc.ClientConnections.Connections {
		fmt.Printf("\nLoadavg %v nproc %v ramavailabe %v\n", b.LoadAverage, b.CpuCount, b.RamAvailable)
	}

}

func HDDUtilization() (retVal []spoc.HddUtil) {

	for _, v := range spoc.ClientConnections.Connections {

		var hd spoc.HddUtil

		cmdHdd := `df -h | awk -F '[[:space:][:space:]]+' ' { print $2","$3","$4","$5 } '`
		s := v.RunCommand(cmdHdd)
		fmt.Printf("df -h:\n %+v", s)

		for _, vv := range strings.Split(s, "\n") {
			//fmt.Printf("string: %+v, len %+v", vv, len(vv))
			/*if i == 0 {
				break
			} else {*/
			if len(vv) > 3 {
				str := strings.Split(vv, ",")
				if str[0] == "Size" {
					continue
				} else {
					if len(str) == 4 {
						hd.Partition = str[0]
						hd.SpaceUsed = str[1]
						hd.SpaceAllocated = str[2]
						hd.PercentageUsed = str[3]

						v.HddDriveUtilization = append(v.HddDriveUtilization, hd)
					}
				}
			}

		}
		return v.HddDriveUtilization
	}
	return
}
