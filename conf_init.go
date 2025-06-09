package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

type SSHForm struct {
	Host     string
	Username string
	Password string
	Delay    int
}

type DATAForm struct {
	Prefix     string
	Delay      int
	Difference int
}

type MQTTForm struct {
	Host     string
	ClientID string
	Username string
	Password string
	Delay    int
}

func ConfInit() (*SSHForm, *DATAForm, *MQTTForm) {
	const fileName = "config.json"
	file, err := os.Open(fileName)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatalln("Error: Unable to create file {", fileName, "}!")
		}
		data, _ := memFS.ReadFile("default.json")
		file.Write(data)
		file.Close()
		log.Println("Error: Config file was not initialized. Fill out the file {", fileName, "}!")
		os.Exit(-1)
	}
	defer file.Close()

	sshStruct := &SSHForm{}
	dataStruct := &DATAForm{}
	mqttStruct := &MQTTForm{}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error: Unable to read file {", fileName, "}!")
	}

	var jsonInfo map[string]interface{}
	err = json.Unmarshal(data, &jsonInfo)
	if err != nil {
		log.Fatalln("Error: Unable to deserialize file {", fileName, "}!")
	}

	sshInfo := jsonInfo["SSH"].(map[string]interface{})
	sshStruct.Host = sshInfo["host"].(string)
	sshStruct.Username = sshInfo["username"].(string)
	sshStruct.Password = sshInfo["password"].(string)
	int_var, err := strconv.Atoi(sshInfo["delay"].(string))
	if sshStruct.Host == "" ||
		sshStruct.Username == "" ||
		sshStruct.Password == "" {
		log.Fatalln("Error: Invalid config [Empty Lines] !")
	} else if err != nil || int_var < 0 {
		log.Fatalln("Error: Invalid config [Wrong Values] !")
	} else {
		sshStruct.Delay = int_var
	}

	dataInfo := jsonInfo["DATA"].(map[string]interface{})
	dataStruct.Prefix = dataInfo["prefix"].(string)
	if dataStruct.Prefix == "" {
		log.Fatalln("Error: Invalid config [Empty Lines] !")
	}
	int_var, err = strconv.Atoi(dataInfo["delay"].(string))
	if err != nil || int_var < 0 {
		log.Fatalln("Error: Invalid config [Wrong Values] !")
	} else {
		dataStruct.Delay = int_var
	}
	int_var, err = strconv.Atoi(dataInfo["difference"].(string))
	if err != nil || int_var < 0 {
		log.Fatalln("Error: Invalid config [Wrong Values] !")
	} else {
		dataStruct.Difference = int_var
	}

	mqttInfo := jsonInfo["MQTT"].(map[string]interface{})
	mqttStruct.Host = mqttInfo["host"].(string)
	mqttStruct.ClientID = mqttInfo["client-id"].(string)
	mqttStruct.Username = mqttInfo["username"].(string)
	mqttStruct.Password = mqttInfo["password"].(string)
	int_var, err = strconv.Atoi(mqttInfo["delay"].(string))
	if mqttStruct.Host == "" ||
		mqttStruct.ClientID == "" ||
		mqttStruct.Username == "" ||
		mqttStruct.Password == "" {
		log.Fatalln("Error: Invalid config [Empty Lines] !")
	} else if err != nil || int_var < 0 {
		log.Fatalln("Error: Invalid config [Wrong Values] !")
	} else {
		mqttStruct.Delay = int_var
	}

	return sshStruct, dataStruct, mqttStruct
}
