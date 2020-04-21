package models

import (
	//"github.com/jinzhu/gorm"
)

type Task struct {
	Model

	Id     int
	ConfigId  string
	InstructionId string
	ScanDay     string
	ScanHour string
	ScanPeriod string
	AgentId string
	ScanNet string
	PortGroupId int
	OptionType int
	VulnelType int
        Description string
}


// GetTags gets a list of tags based on paging and constraints
func GetCommands() ([]Task, error) {
	var (
		commands []Task
	)

	db.Find(&commands)

	return commands, nil
}

