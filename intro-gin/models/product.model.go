package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id    			uuid.UUID 		`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name  			string 			`gorm:"type:varchar(100)"`
	Price 			int    			`gorm:"check:price >= 0 AND price <= 1000000"`
	Stock 			*int	 		`gorm:"check:price >= 0 AND price <= 1000000"`
	Categories 		[]Category		`gorm:"many2many:product_categories"`
	ProductImages  	[]ProductImage	`gorm:"foreignKey:ProductId"` // One-to-Many
	CreatedAt 		time.Time		
	UpdatedAt 		time.Time		
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
}