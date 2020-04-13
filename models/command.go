package models

import (
	//"github.com/jinzhu/gorm"
)

type Command struct {
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
func GetCommands() ([]Command, error) {
	var (
		commands []Command
	)

	db.Find(&commands)

	return commands, nil
}

