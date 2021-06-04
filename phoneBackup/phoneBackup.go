package phoneBackup

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func PhoneBackup() string {
	backup := false

	c, b := exec.Command("cfgutil", "list-backups"), new(strings.Builder) //cfgutil CLI provided by https://apps.apple.com/us/app/apple-configurator-2/id1037126344
	c.Stdout = b
	c.Run()
	s := b.String()
	splitted := strings.Split(s, ":")

	miniuteAndAmPm := strings.Split(splitted[4], " Name")[0]
	dayAndHour := splitted[3]

	dateString := dayAndHour + ":" + miniuteAndAmPm
	dateString = strings.ToUpper(dateString)

	print(dateString)

	if strings.Contains(dateString, "TODAY") { //check if it has atleast been around 24 hours since the last backup
		//no need to run backup, its already done today
	} else if strings.Contains(dateString, "YESTERDAY") {
		//check time
		var amOrpm = ""

		if strings.Contains(miniuteAndAmPm, " AM") {
			amOrpm = "AM"
		} else if strings.Contains(miniuteAndAmPm, " PM") {
			amOrpm = "PM"
		} else {
			print("INVALID TIME FORMATTING")
		}

		hour := strings.Split(strings.ToUpper(dayAndHour), "YESTERDAY AT")[1]
		//print(amOrpm)
		//print(hour)

		currentTime := time.Now()
		currentTimeString := currentTime.Format("3 PM")

		currentHour := strings.Split(currentTimeString, " ")[0]
		currentAmorPm := strings.Split(currentTimeString, " ")[1]

		//println("ok: " + currentHour + currentAmorPm)

		if currentAmorPm == amOrpm {
			//compare hour since AM and PM are same
			if currentHour >= hour {
				backup = true
			} else {
				backup = false
			}
		} else if currentAmorPm == "AM" && amOrpm == "PM" {
			backup = false
		} else {
			backup = true
		}

	} else {
		backup = true
	}

	if backup {
		fmt.Println("Trying To Backup...")

		cc, bb := exec.Command("cfgutil", "backup"), new(strings.Builder)
		cc.Stdout = bb
		cc.Run()
		backupRes := bb.String()

		currentTime := time.Now()

		if strings.Contains(backupRes, "ECID") {
			//ran successfully
			fmt.Println("Backed up successfully on " + currentTime.Format("01-02-2006 15:04:05"))
			return "Backed up successfully"
		} else {
			//Error, device most likely not found or available
			fmt.Println("Device probably not connected")
			return "Error occured, Device probably not connected"
		}
	}

	return "Not Backing Up"

}
