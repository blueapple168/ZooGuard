package targetApplications

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
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

		//hddutilization
		HDDUtilization(v)
		fmt.Printf("\nDF- h %+v\n", v.HddDriveUtilization)

		//ulimit
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

		//uptime
		cmdLoadAvg := `uptime`
		sLA := v.RunCommand(cmdLoadAvg)
		fmt.Printf("Load average %v\n", sLA)
		str1 := strings.Split(sLA, ":")
		n := len(str1)
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

		//nproc
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
		for _, load := range v.LoadAverage {
			cpuload := load / (float64(v.CpuCount))
			v.CpuAdjustedLoadAverage = append(v.CpuAdjustedLoadAverage, cpuload)
		}

		//ram available
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

		//timedelta
		cmdTime := `date`
		localTime := time.Now()
		ctime := v.RunCommand(cmdTime)
		re4 := regexp.MustCompile(`\n`)
		stt := re4.ReplaceAllString(ctime, "")
		st := strings.Split(stt, " ")
		fmt.Printf("\nDate command output: %v %v", st, v.ServerName)
		i, _ := strconv.ParseInt(st[2], 10, 64)
		fmt.Printf("Value of i is %v",i)
		var serverTimeStr string
		if i < 10 {
			serverTimeStr = fmt.Sprintf("%v %v 0%v %v ICT", st[5], st[1], st[2], st[3])
		} else {
			serverTimeStr = fmt.Sprintf("%v %v %v %v ICT", st[5], st[1], st[2], st[3])
		}
		fmt.Printf("\ndate output sprintf %v\n", serverTimeStr)
		timeFormat := "2006 Jan 02 15:04:05 MST"
		serverTime, errTime := time.Parse(timeFormat, serverTimeStr)
		if errTime != nil {
			fmt.Println(errTime)
		}
		duration := serverTime.Sub(localTime)
		duraint := int64(duration / time.Millisecond)
		timedel := &v.TimeDelta
		*timedel = duraint
		fmt.Printf("\nDifference in time:%v, difference in int %v\n", duration, v.TimeDelta)

		//OSversion
		cmdOSver := `cat /etc/redhat-release`
		strOS := v.RunCommand(cmdOSver)
		OSVer := &v.OSVersion
		*OSVer = strOS

	}

	for _, b := range spoc.ClientConnections.Connections {
		fmt.Printf("\nLoadavg %v CPUAdjustloadavg %v nproc %v ramavailable %v osversion %v servername %v\n", b.LoadAverage, b.CpuAdjustedLoadAverage, b.CpuCount, b.RamAvailable, b.OSVersion, b.ServerName)
	}

}

func HDDUtilization(v *spoc.ConnInfo) {

	var hd spoc.HddUtil

	cmdHdd := `df -h | awk -F '[[:space:][:space:]]+' ' { print $2","$3","$4","$5 } '`
	s := v.RunCommand(cmdHdd)
	fmt.Printf("df -h:\n %+v", s)

	for _, vv := range strings.Split(s, "\n") {

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

	cmdTime := `date`
	ctime := v.RunCommand(cmdTime)
	re5 := regexp.MustCompile(`\n`)
	stt := re5.ReplaceAllString(ctime, "")
	st := strings.Split(stt, " ")

	i, _ := strconv.ParseInt(st[2], 10, 64)
	var serverTimeStr string
	if i < 10 {
		serverTimeStr = fmt.Sprintf("%v %v 0%v %v ICT", st[5], st[1], st[2], st[3])
	} else {
		serverTimeStr = fmt.Sprintf("%v %v %v %v ICT", st[5], st[1], st[2], st[3])
	}
	timeFormat := "2006 Jan 02 15:04:05 MST"
	serverTime, errTime := time.Parse(timeFormat, serverTimeStr)

	if errTime != nil {
		fmt.Println(errTime)
	} else {
		v.HddLastChecked = serverTime
	}

}
