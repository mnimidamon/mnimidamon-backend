package model

// Represents a Computer that is included in the Group. Denotes how much storage it contributes to the group in MB.
type GroupComputer struct {
	Entity

	GroupID    uint
	ComputerID uint

	Group    Group
	Computer Computer

	StorageSize uint
}
