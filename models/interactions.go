package models

import (
	"github.com/google/uuid"
)

type Interaction struct {
	InteractionID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LeadID        uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID    uuid.UUID `gorm:"type:uuid;not null"`
	Type          string    `gorm:"type:VARCHAR(255);check:Type IN ('call', 'email', 'meeting', 'lunch', 'other');default:'call'"`
	Notes         string    `gorm:"type:text"`
	Lead          Lead      `gorm:"foreignKey:LeadID;references:LeadID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Customer      Customer  `gorm:"foreignKey:CustomerID;references:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
