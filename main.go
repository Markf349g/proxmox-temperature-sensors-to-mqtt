package main

import (
	"embed"
	"log"
	"math"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//go:embed default.json
var memFS embed.FS

func main() {
	//Just make with google get http response
	sshConf, dataPtr, mqttPtr := ConfInit()

	sshPtr := SSHInit(sshConf)
	defer sshPtr.Close()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttPtr.Host)
	opts.SetClientID(mqttPtr.ClientID)
	opts.SetUsername(mqttPtr.Username)
	opts.SetPassword(mqttPtr.Password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(time.Duration(mqttPtr.Delay) * time.Millisecond)
	opts.CleanSession = false

	client := mqtt.NewClient(opts)

	for {
		token := client.Connect()
		token.Wait()
		if token.Error() != nil {
			log.Println(token.Error())
			time.Sleep(time.Duration(mqttPtr.Delay))
		} else {
			defer client.Disconnect(uint(mqttPtr.Delay))
			break
		}
	}

	var crt_map map[string]string
	lst_map := make(map[string]string)
	var key, value string
	var exists bool
	var token mqtt.Token
	var crt_value, lst_value int

	send := func(crt_map, lst_map *map[string]string, key, value string) {
		log.Println(key, ":", value)
		token = client.Publish(key, 1, true, value)
		(*lst_map)[key] = (*crt_map)[key]
		token.Wait()
	}

	for {
		WaitTheInternet(time.Duration(sshPtr.Form.Delay) * time.Millisecond)
		crt_map = telemetryRequest(sshPtr, dataPtr)
		for key, value = range crt_map {
			_, exists = lst_map[key]
			if !exists {
				send(&crt_map, &lst_map, key, value)
			} else if crt_map[key] != lst_map[key] {
				if crt_map[key] == "Unknown" || lst_map[key] == "Unknown" {
					send(&crt_map, &lst_map, key, value)
				} else {
					crt_value, _ = strconv.Atoi(crt_map[key])
					lst_value, _ = strconv.Atoi(lst_map[key])
					if (math.Abs(float64(lst_value)) - math.Abs(float64(crt_value))) >= float64(dataPtr.Difference) {
						send(&crt_map, &lst_map, key, value)
					} else {
						continue
					}
				}
			} else {
				continue
			}
		}
		time.Sleep(time.Duration(dataPtr.Delay) * time.Second)
	}
}
