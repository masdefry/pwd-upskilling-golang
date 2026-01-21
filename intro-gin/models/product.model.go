package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id    			uuid.UUID 		`gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  			string 			`gorm:"type:varchar(100)" json:"name"`
	Price 			int    			`gorm:"check:price >= 0 AND price <= 1000000" json:"price"`
	Stock 			*int	 		`gorm:"check:price >= 0 AND price <= 1000000" json:"stock"`
	Categories 		[]Category		`gorm:"many2many:product_categories"`
	ProductImages  	[]ProductImage	`gorm:"foreignKey:ProductId"` // One-to-Many
	CreatedAt 		time.Time		`json:"createdAt"`
	UpdatedAt 		time.Time		`json:"updatedAt"`
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}