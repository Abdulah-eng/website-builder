package models

import (
	"time"

	"gorm.io/gorm"
)

type StoreStatus string

const (
	StoreStatusActive   StoreStatus = "active"
	StoreStatusInactive StoreStatus = "inactive"
	StoreStatusDraft    StoreStatus = "draft"
)

type Store struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Description string         `json:"description"`
	Logo        string         `json:"logo"`
	Favicon     string         `json:"favicon"`
	Domain      string         `json:"domain" gorm:"uniqueIndex"`
	Subdomain   string         `json:"subdomain" gorm:"uniqueIndex"`
	Status      StoreStatus    `json:"status" gorm:"default:'draft'"`
	OwnerID     uint           `json:"owner_id" gorm:"not null"`
	TemplateID  *uint          `json:"template_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Owner      User           `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Template   *Template      `json:"template,omitempty" gorm:"foreignKey:TemplateID"`
	Products   []Product      `json:"products,omitempty" gorm:"foreignKey:StoreID"`
	Categories []Category     `json:"categories,omitempty" gorm:"foreignKey:StoreID"`
	Orders     []Order        `json:"orders,omitempty" gorm:"foreignKey:StoreID"`
	Settings   *StoreSettings `json:"settings,omitempty" gorm:"foreignKey:StoreID"`
	Theme      *StoreTheme    `json:"theme,omitempty" gorm:"foreignKey:StoreID"`
	Pages      []Page         `json:"pages,omitempty" gorm:"foreignKey:StoreID"`
}

type StoreCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	TemplateID  *uint  `json:"template_id"`
}

type StoreUpdateRequest struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Logo        *string      `json:"logo,omitempty"`
	Favicon     *string      `json:"favicon,omitempty"`
	Domain      *string      `json:"domain,omitempty"`
	Status      *StoreStatus `json:"status,omitempty"`
}

type StoreResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Description string      `json:"description"`
	Logo        string      `json:"logo"`
	Favicon     string      `json:"favicon"`
	Domain      string      `json:"domain"`
	Subdomain   string      `json:"subdomain"`
	Status      StoreStatus `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
