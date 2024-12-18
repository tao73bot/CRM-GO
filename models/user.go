package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	AdminID   uuid.UUID  `gorm:"type:uuid;not null"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);unique;not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Role      string     `gorm:"type:VARCHAR(255);check:role IN('admin', 'user');default:'user'"`
	Admin     Admin      `gorm:"foreignKey:AdminID;references:AdminID"`
}
