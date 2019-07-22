package entities

import (
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type HardwareInfo struct {
	Host        Host            `json:"host"`
	CPU         CPU             `json:"cpu"`
	Memory      []*Memory       `json:"memory"`
	Disk        []*Disk         `json:"disk"`
	GPU         []*GraphicsCard `json:"gpu"`
	NetworkCard []*NetworkCard  `json:"network_card"`
}

type Host host.InfoStat
type Memory mem.VirtualMemoryStat
type CPU cpu.InfoStat
type Disk disk.UsageStat
type NetworkCard net.IOCountersStat
type GraphicsCard ghw.GraphicsCard
