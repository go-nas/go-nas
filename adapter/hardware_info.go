package adapter

import (
	"gnas/entities"

	"github.com/shirou/gopsutil/cpu"

	"github.com/shirou/gopsutil/host"
)

type HardwareInfo struct {
}

func (hi *HardwareInfo) GetHardwareInfo() (*entities.HardwareInfo, error) {
	h, err := host.Info()
	if err != nil {
		return nil, err
	}
	c, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	//CPU         CPU             `json:"cpu"`
	//Memory      []*Memory       `json:"memory"`
	//Disk        []*Disk         `json:"disk"`
	//GPU         []*GraphicsCard `json:"gpu"`
	//NetworkCard []*NetworkCard  `json:"network_card"`
	return &entities.HardwareInfo{
		Host: entities.Host(*h),
		CPU:  entities.CPU(*c),
	}, nil
}
