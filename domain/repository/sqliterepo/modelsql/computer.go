package modelsql

type Computer struct {
	Entity

	OwnerID uint   `gorm:"uniqueIndex:idx_user_computers"`
	Name    string `gorm:"uniqueIndex:idx_user_computers; size:15"`

	Owner User `gorm:"foreignKey:OwnerID"`
}
