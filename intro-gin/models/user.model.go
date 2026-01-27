package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)


type User struct {
	Id    			uuid.UUID 		`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  			string 		`gorm:"type:varchar(100)"`
	Email  			string 			`gorm:"type:varchar(255)"`
	Password  			string 		`gorm:"type:varchar(225)"`
	
	Role UserRole 					`gorm:"type:varchar(20);default:'USER'"`

	CreatedAt 		time.Time		
	UpdatedAt 		time.Time		
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
}