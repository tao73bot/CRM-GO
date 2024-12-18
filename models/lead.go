package models

import (
	"github.com/google/uuid"
)

type Lead struct {
	LeadID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Name   string    `gorm:"type:varchar(255);not null"`
	Email  string    `gorm:"type:varchar(255);unique;not null"`
	Phone  string    `gorm:"type:varchar(255);not null"`
	Status string    `gorm:"type:VARCHAR(255);check:status IN('new', 'contacted', 'qualified', 'lost', 'closed');default:'new'"`
	Source string    `gorm:"type:VARCHAR(255);check:source IN('website', 'phone', 'email', 'social media');default:'website'"`
	User   User      `gorm:"foreignKey:UserID;references:UserID"`
}
