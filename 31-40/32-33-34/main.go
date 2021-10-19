package main

import (
	"encoding/json"
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

func main() {
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
	fmt.Println(string(finalData))
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
