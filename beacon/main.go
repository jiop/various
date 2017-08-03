package main

import (
	"fmt"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func gatherInfo() {
	virtMem, _ := mem.VirtualMemory()
	fmt.Print("VirtualMemory", virtMem, "\n\n")

	swapMem, _ := mem.SwapMemory()
	fmt.Print("SwapMemory", swapMem, "\n\n")

	interfaces, _ := net.Interfaces()
	fmt.Print("Network Interfaces", interfaces, "\n\n")

	cpuInfo, _ := cpu.Info()
	fmt.Print("CPU", cpuInfo, "\n\n")

	hostInfo, _ := host.Info()
	fmt.Print("Host", hostInfo, "\n\n")

	users, _ := host.Users()
	fmt.Print("Users", users, "\n\n")
}

func main() {
	gatherInfo()

	_, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%#v", resp)
}
