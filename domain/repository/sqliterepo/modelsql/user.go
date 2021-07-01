package modelsql

type User struct {
	Entity

	Username     string `gorm:"unique; index; size:15"`
	PasswordHash string

	Computers []Computer `gorm:"foreignKey:OwnerID"`
}
