package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type NetworkInterface struct {
	Mode          string   `json:"mode"`
	Index         int      `json:"index"`
	PortForwards  []string `json:"portForwards"`
	HostInterface string   `json:"hostInterface"`
	Address       string   `json:"address"`
	Hardware      string   `json:"hardware"`
}

type Drive struct {
	Removable bool   `json:"removable"`
	ID        string `json:"id"`
	Interface string `json:"interface"`
	HostSize  int    `json:"hostSize"`
}

type Configuration struct {
	NetworkInterfaces  []NetworkInterface `json:"networkInterfaces"`
	SerialPorts        []string           `json:"serialPorts"`
	Drives             []Drive            `json:"drives"`
	Machine            string             `json:"machine"`
	DirectoryShareMode string             `json:"directoryShareMode"`
	UEFI               bool               `json:"uefi"`
	Name               string             `json:"name"`
	Notes              string             `json:"notes"`
	Architecture       string             `json:"architecture"`
	Memory             int                `json:"memory"`
	Hypervisor         bool               `json:"hypervisor"`
	CPUCores           int                `json:"cpuCores"`
}

type VMInstance struct {
	Status        string        `json:"status"`
	ID            string        `json:"id"`
	Configuration Configuration `json:"configuration"`
	Backend       string        `json:"backend"`
	Pcls          string        `json:"pcls"`
	Name          string        `json:"name"`
}

//go:embed jxa/get_vms.js
var getVms string

func main() {
	cmd := exec.Command("/usr/bin/osascript", "-l", "JavaScript", "-e", getVms)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var VMInstances []VMInstance
	if err := json.Unmarshal(output, &VMInstances); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}
	// Now, you can work with the parsed data in the vmInstances slice
	for _, vm := range VMInstances {
		fmt.Printf("VM Name: %s\n", vm.Name)
		fmt.Printf("Status: %s\n", vm.Status)
		// Access other fields as needed
	}
}
