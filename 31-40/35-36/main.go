package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	dotenv "github.com/joho/godotenv"
	cpuUtil "github.com/shirou/gopsutil/v3/cpu"
	diskUtil "github.com/shirou/gopsutil/v3/disk"
	loadUtil "github.com/shirou/gopsutil/v3/load"
	memUtility "github.com/shirou/gopsutil/v3/mem"
	"log"
	"os"
	"time"
)

var MqttDetails = []string{
	"KG_MQTT_HOST", "KG_MQTT_PORT", "KG_MQTT_USER", "KG_MQTT_PASS",
}

var credentials = make(map[string]string)

var uniqMachId string
var err error

type swapM struct {
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
}

type virtM struct {
	VirtTotal     uint64
	VirtAvailable uint64
	VirtUsed      uint64
	VirtFree      uint64
}

type DiskD struct {
	DiskPath       string
	DiskTotal      uint64
	DiskUsed       uint64
	DiskFree       uint64
	DiskInodeTotal uint64
	DiskInodeUsed  uint64
	DiskInodeFree  uint64
}

type CPUData struct {
	CpuCore  int32
	CpuMHZ   float64
	CpuCache int32
	CpuCpu   int32
}

type LoadAvg struct {
	LoadForOneMin     float64
	LoadForFiveMin    float64
	LoadForFifteenMin float64
}

type Data struct {
	UniqID string
	swapM
	virtM
	DiskD []DiskD
	CPUData
	LoadAvg
}

func init() {
	fmt.Println("Initializing........")

	var MqttListNotPresent []string
	var DotfileListNotPresent []string
	var DotFileExists = true

	for _, list := range MqttDetails {
		_, ifPresent := os.LookupEnv(list)
		if !ifPresent {
			MqttListNotPresent = append(MqttListNotPresent, list)
		} else {
			credentials[list] = os.Getenv(list)
		}
	}

	if len(MqttListNotPresent) > 0 {
		fmt.Println("Environment variable(s) not declared. Will try reading from .env file.")
		content, err := dotenv.Read()

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(".env file doesn't exists, creating new one.")
			DotFileExists = false
		}

		if errors.Is(err, os.ErrPermission) {
			// handle the case where the file doesn't exist
			fmt.Println("Check permission of .env file.")
			os.Exit(0)
		}

		if err == nil {
			for _, list := range MqttListNotPresent {
				if _, ok := content[list]; !ok {
					DotfileListNotPresent = append(DotfileListNotPresent, list)
				} else {
					credentials[list] = content[list]
				}
			}
		}

		if (len(DotfileListNotPresent) > 0) || !DotFileExists {
			f, err := os.OpenFile(".env", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			rangeList := DotfileListNotPresent

			if !DotFileExists {
				rangeList = MqttListNotPresent
			}
			// Flush Content
			for _, list := range rangeList {
				var value string
				fmt.Println("Enter Value for ", list)
				_, err := fmt.Scanln(&value)
				if err != nil {
					return
				}
				toWrite := fmt.Sprintf("%s=%s\n", list, value)
				credentials[list] = value
				if _, err := f.Write([]byte(toWrite)); err != nil {
					log.Fatal(err)
				}
				fmt.Println(toWrite)
			}

		}

	}

	const appKey = "d6712094222c4eb9a71d1976e934b869e60cf4fbcb7279b653858ac3109815c0"
	uniqMachId, err = machineid.ProtectedID(appKey)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", credentials["KG_MQTT_HOST"], credentials["KG_MQTT_PORT"]))

	// Check if Auth is implemented or not.
	if credentials["KG_MQTT_USER"] != "" {
		opts.SetUsername(credentials["KG_MQTT_USER"])
		opts.SetPassword(credentials["KG_MQTT_PASS"])
	}
	CreateUniqID := fmt.Sprintf("%s-%s", "GO-CLIENT", uniqMachId)
	opts.SetClientID(CreateUniqID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	topic := fmt.Sprintf("%s/%s", "stats", uniqMachId)

	sub(client, fmt.Sprintf("%s/cmnd", topic))
	publish(client, topic, "Message")
	for {
		// Run Infinite Loop, because why not.
		swapMemDetails, err := memUtility.SwapMemory()
		errorHandling(err, "log")
		swapData := swapM{
			SwapTotal: swapMemDetails.Total,
			SwapUsed:  swapMemDetails.Free,
			SwapFree:  swapMemDetails.Free,
		}

		virtMemDetails, err := memUtility.VirtualMemory()
		errorHandling(err, "log")
		virtData := virtM{
			VirtTotal:     virtMemDetails.Total,
			VirtAvailable: virtMemDetails.Available,
			VirtUsed:      virtMemDetails.Used,
			VirtFree:      virtMemDetails.Free,
		}

		partitionDetails, err := diskUtil.Partitions(true)
		errorHandling(err, "log")
		var diskData []DiskD
		for _, disk := range partitionDetails {
			partDetails, err := diskUtil.Usage(disk.Mountpoint)
			errorHandling(err, "log")
			diskData = append(diskData, DiskD{
				DiskPath:       partDetails.Path,
				DiskTotal:      partDetails.Total,
				DiskUsed:       partDetails.Used,
				DiskFree:       partDetails.Free,
				DiskInodeTotal: partDetails.InodesTotal,
				DiskInodeUsed:  partDetails.InodesUsed,
				DiskInodeFree:  partDetails.InodesFree,
			})
		}

		cpuDetails, err := cpuUtil.Info()
		errorHandling(err, "log")
		var cpuData CPUData
		for _, cpu := range cpuDetails {
			cpuData = CPUData{
				CpuCore:  cpu.Cores,
				CpuMHZ:   cpu.Mhz,
				CpuCache: cpu.CacheSize,
				CpuCpu:   cpu.CPU,
			}

		}

		loadDetails, err := loadUtil.Avg()
		errorHandling(err, "log")
		loadData := LoadAvg{
			LoadForOneMin:     loadDetails.Load1,
			LoadForFiveMin:    loadDetails.Load5,
			LoadForFifteenMin: loadDetails.Load15,
		}
		finalData, err := json.Marshal(Data{uniqMachId, swapData, virtData, diskData, cpuData, loadData})
		errorHandling(err)
		publish(client, topic, string(finalData))
		time.Sleep(time.Second * 10)
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func publish(client mqtt.Client, topic string, text string) {
	token := client.Publish(topic, 0, false, text)
	token.Wait()
}

func errorHandling(err error, typeOfError ...string) {
	if err != nil {
		if len(typeOfError) != 0 {
			switch typeOfError[0] {
			case "log":
				log.Fatal(err)
			case "panic":
				panic(err)
			default:
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
