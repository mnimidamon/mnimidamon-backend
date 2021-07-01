package modelsql

type User struct {
	UserID       uint   `gorm:"primaryKey"`
	Username     string `gorm:"size:15"`
	PasswordHash string
	Entity
}
