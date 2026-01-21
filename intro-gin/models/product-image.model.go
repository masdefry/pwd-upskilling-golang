package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	Id        uuid.UUID			`gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Url		  string 			`json:"url"`
	ProductId uuid.UUID			`json:"productId"`// Foreign Key (From Model Product)	
	CreatedAt time.Time		 	`json:"createdAt"`
	UpdatedAt time.Time		 	`json:"updatedAt"`
	DeletedAt gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}