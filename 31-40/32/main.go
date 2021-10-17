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

func main() {
	const appKey = "d6712094222c4eb9a71d1976e934b869e60cf4fbcb7279b653858ac3109815c0"
	id, err := machineid.ProtectedID(appKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	fmt.Println(memUtility.SwapMemory())
	fmt.Println(memUtility.VirtualMemory())
	fmt.Println(diskUtil.Usage("/"))
	fmt.Println(diskUtil.Partitions(true))
	fmt.Println(cpuUtil.Info())
	fmt.Println(loadUtil.Avg())
}
