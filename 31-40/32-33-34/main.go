package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	cpuUtil "github.com/shirou/gopsutil/v3/cpu"
	diskUtil "github.com/shirou/gopsutil/v3/disk"
	loadUtil "github.com/shirou/gopsutil/v3/load"
	memUtility "github.com/shirou/gopsutil/v3/mem"
	"log"
)

var uniqMachId string
var err error

func init() {
	const appKey = "d6712094222c4eb9a71d1976e934b869e60cf4fbcb7279b653858ac3109815c0"
	uniqMachId, err = machineid.ProtectedID(appKey)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(uniqMachId)
	swapMemDetails, err := memUtility.SwapMemory()
	errorHandling(err, "log")
	fmt.Println(swapMemDetails)

	virtMemDetails, err := memUtility.VirtualMemory()
	errorHandling(err, "log")
	fmt.Println(virtMemDetails)

	partitionDetails, err := diskUtil.Partitions(true)
	errorHandling(err, "log")
	fmt.Println(partitionDetails)
	//To loop each partition to get details
	//fmt.Println(diskUtil.Usage("/"))

	cpuDetails, err := cpuUtil.Info()
	errorHandling(err, "log")
	fmt.Println(cpuDetails)

	loadDetails, err := loadUtil.Avg()
	errorHandling(err, "log")
	fmt.Println(loadDetails)
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
