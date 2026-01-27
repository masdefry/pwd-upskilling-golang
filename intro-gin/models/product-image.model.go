package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	Id        uuid.UUID			`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Url		  string 	
	ProductId uuid.UUID			// Foreign Key (From Model Product)	
	CreatedAt time.Time		 	
	UpdatedAt time.Time		 	
	DeletedAt gorm.DeletedAt 	`gorm:"index"`
}