package repository

import "gnas/entities"

type IHardwareInfo interface {
	GetHardwareInfo() (*entities.HardwareInfo, error)
}
