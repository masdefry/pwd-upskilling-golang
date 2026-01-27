package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string		 `gorm:"type:varchar(100)"`
	Products  []Product      `gorm:"many2many:product_categories"`
	CreatedAt time.Time		
	UpdatedAt time.Time		
	DeletedAt gorm.DeletedAt `gorm:"index"`
}