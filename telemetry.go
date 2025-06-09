package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func tempRequest(sshPtr *SSHStruct, request string) string {
	var temp string
	output, err := SSHRequest(sshPtr, request)
	if err != nil {
		temp = "Unknown"
	} else {
		result, _ := strconv.Atoi(output)
		result /= 1000
		temp = strconv.Itoa(result)
	}
	return temp
}

func nameRequest(sshPtr *SSHStruct, request string) string {
	var label string
	output, err := SSHRequest(sshPtr, request)
	if err != nil {
		label = "Unknown"
	} else {
		label = strings.ToLower(strings.ReplaceAll(output, " ", "-"))
	}
	return label
}

func telemetryRequest(sshPtr *SSHStruct, dataPtr *DATAForm) map[string]string {
	result_map := make(map[string]string)

	hwmon_prefix := "/sys/class/hwmon"
	var output string
	var err error

	for {
		output, err = SSHRequest(sshPtr, fmt.Sprint("ls ", hwmon_prefix))
		if err != nil {
			log.Println("Command failed: ", err)
		} else {
			break
		}
	}

	hwmon_dirs := strings.Split(output, "\n")
	for _, hwmon_dir := range hwmon_dirs {
		name := nameRequest(sshPtr, fmt.Sprint("cat ", hwmon_prefix, "/", hwmon_dir, "/", "name"))

		var temp_prefix, label, temp, index_prefix string
		for index := 1; ; index++ {
			index_prefix = fmt.Sprint("/temp", index)
			temp_prefix = fmt.Sprint("cat ", hwmon_prefix, "/", hwmon_dir, index_prefix)

			temp = tempRequest(sshPtr, temp_prefix+"_input")
			if temp == "Unknown" {
				break
			}

			label = fmt.Sprint(dataPtr.Prefix, "/", name, index_prefix)
			result_map[label+"_input"] = temp

			temp = tempRequest(sshPtr, temp_prefix+"_min")
			result_map[label+"_min"] = temp
			temp = tempRequest(sshPtr, temp_prefix+"_max")
			result_map[label+"_max"] = temp
			temp = tempRequest(sshPtr, temp_prefix+"_crit")
			result_map[label+"_crit"] = temp
			temp = tempRequest(sshPtr, temp_prefix+"_crit_alarm")
			result_map[label+"_crit_alarm"] = temp
		}
	}
	return result_map
}
