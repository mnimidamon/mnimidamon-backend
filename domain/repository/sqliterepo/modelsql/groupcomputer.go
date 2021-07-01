package modelsql

// Represents a Computer that is included in the Group. Denotes how much storage it contributes to the group in MB.
type GroupComputer struct {
	Entity

	GroupID    uint
	ComputerID uint

	Group    Group    `gorm:"foreignKey:GroupID"`
	Computer Computer `gorm:"foreignKey:ComputerID"`

	StorageSize uint
}
