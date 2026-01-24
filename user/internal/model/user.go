package model

type User struct {
	ID       uint   `gorm:"primarykey"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
}
