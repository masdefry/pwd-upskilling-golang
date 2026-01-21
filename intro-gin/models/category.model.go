package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string		 `gorm:"type:varchar(100)" json:"name"`
	Products  []Product      `gorm:"many2many:product_categories"`
	CreatedAt time.Time		 `json:"createdAt"`
	UpdatedAt time.Time		 `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}