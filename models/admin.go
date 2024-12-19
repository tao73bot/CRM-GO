package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	AdminID  uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	// User	 []User    `gorm:"foreignKey:AdminID;references:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
