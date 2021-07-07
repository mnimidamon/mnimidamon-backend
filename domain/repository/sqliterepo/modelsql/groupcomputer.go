package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

// Represents a Computer that is included in the Group. Denotes how much storage it contributes to the group in MB.
type GroupComputer struct {
	Entity

	GroupID    uint
	ComputerID uint

	Group    Group    `gorm:"foreignKey:GroupID"`
	Computer Computer `gorm:"foreignKey:ComputerID"`

	StorageSize uint
}


func NewGroupComputerFromBusinessModel(gcm *model.GroupComputer) *GroupComputer {
	if gcm == nil {
		return nil
	}

	gc := &GroupComputer{
		Entity:      Entity{
			ID:        gcm.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		GroupID:     gcm.GroupID,
		ComputerID:  gcm.ComputerID,
		Group:       *NewGroupFromBusinessModel(&gcm.Group),
		Computer:    *NewComputerFromBusinessModel(&gcm.Computer),
		StorageSize: gcm.StorageSize,
	}

	return gc
}

func (gc *GroupComputer) NewBusinessModel() *model.GroupComputer {
	if gc == nil {
		return nil
	}

	gcm := new(model.GroupComputer)
	gc.CopyToBusinessModel(gcm)
	return gcm
}

func (gc *GroupComputer) CopyToBusinessModel(gcm *model.GroupComputer) {
	if gc == nil {
		gcm = nil
		return
	}

	gcm.ID = gc.ID
	gcm.GroupID = gc.GroupID
	gcm.ComputerID = gc.ComputerID
	gcm.Group = *gc.Group.NewBusinessModel()
	gcm.Computer = *gc.Computer.NewBusinessModel()
	gcm.StorageSize = gc.StorageSize
}
