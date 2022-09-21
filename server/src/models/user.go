package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID   `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email        string      `json:"email" gorm:"unique;not null"`
	Username     string      `json:"username" gorm:"unique;not null"`
	Password     string      `json:"password" gorm:"not null"`
	PasswordList []Passwords `json:"password_list"`
}

type Passwords struct {
	ID             uuid.UUID     `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Password       string        `json:"password" gorm:"not null"`
	PasswordName   PasswordNames `json:"password_name" gorm:"foreignKey:PasswordNameId"`
	PasswordNameId uuid.UUID     `json:"password_name_id"`
	UserId         uuid.UUID     `json:"user_id"  gorm:"not null"`
	User           User          `json:"-" gorm:"foreignKey:UserId"`
}

type PasswordNames struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"unique;not null"`
}
