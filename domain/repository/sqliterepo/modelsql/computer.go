package modelsql

import "mnimidamonbackend/domain/model"

type Computer struct {
	Entity

	OwnerID uint   `gorm:"uniqueIndex:idx_user_computers"`
	Name    string `gorm:"uniqueIndex:idx_user_computers; size:15"`

	Owner User `gorm:"foreignKey:OwnerID"`
}

func NewComputerFromBusinessModel(cm *model.Computer) *Computer {
	if cm == nil {
		return nil
	}

	return &Computer{
		Entity:  Entity{
			ID: cm.ID,
		},
		OwnerID: cm.OwnerID,
		Name:    cm.Name,
		Owner:   *NewUserFromBusinessModel(&cm.Owner),
	}
}

func (c *Computer) NewBusinessModel() *model.Computer {
	if c == nil {
		return nil
	}

	cm := new(model.Computer)
	c.CopyToBusinessModel(cm)
	return cm
}

func (c *Computer) CopyToBusinessModel(cm *model.Computer)  {
	if c == nil {
		cm = nil
		return
	}

	cm.Name = c.Name
	cm.OwnerID = c.OwnerID
	cm.ID = c.ID
	cm.Owner = *c.Owner.NewBusinessModel()
}