Hello, Full Stack Web Development Lecturers✌️!

🖥️ How to Setup Gin Project?

        ▪️Step-01    : Create new directory

        ▪️Step-02    : Init `go module`

                        go mod init <module-name>

        ▪️Step-03    : Install Gin framework

                        go get -u github.com/gin-gonic/gin

        ▪️Step-04    : Create entry point (main.go)

                        package main

                        import (
                                "fmt"

                                "github.com/gin-gonic/gin"
                        )

                        func main() {
                                r := gin.Default()
                                        port := 8080;

                                r.GET("/ping", func(c *gin.Context) {
                                        c.JSON(200, gin.H{
                                        "message": "pong",
                                        })
                                })

                                r.Run(fmt.Sprintf(":%d", port));
                        }

        ▪️Step-05    : Running the application

                        go run main.go ___OR___ air

🖥️ How to Setup Gorm?

        ▪️Step-01    : Install this packages
                        
                        [MAIN-GORM] go get -u gorm.io/gorm

                        [DRIVER-POSTGRESQL] go get -u gorm.io/driver/postgres

                        [UUID-GENERATOR] go get github.com/jackc/pgx/pgtype/ext/gofrs-uuid

        ▪️Step-02    : Create models on `directory models`

                        type Product struct {
                                Id    uuid.UUID 			`gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
                                Name  string 				`json:"name"`
                                Price int    				`gorm:"check:price >= 0 AND price <= 1000000" json:"price"`
                                Stock int	 				`gorm:"check:price >= 0 AND price <= 1000000" json:"stock"`
                                Categories []Category		`gorm:"many2many:product_categories"`
                                ProductImages  []ProductImage      `gorm:"foreignKey:ProductId"` // One-to-Many
                                CreatedAt time.Time			`json:"createdAt"`
                                UpdatedAt time.Time			`json:"updatedAt"`
                                DeletedAt gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
                        }

                        ...
                        ...
                        ...

                        📝
                        ▪️DeletedAt harus bertipe gorm.DeletedAt.
                        ▪️Saat db.Delete(&product) dipanggil, GORM akan mengisi DeletedAt dengan timestamp
                        saat itu, record tidak hilang dari table.
                        ▪️Query normal seperti db.Find(&products) otomatis mengecualikan record yang sudah soft delete.
                        ▪️Jika ingin query tetap memasukan data yang telah di soft-delete, bisa gunakan syntax .Unscoped(): db.Unscoped().Find(&products).

🖥️ How to Setup Environment Variable in Golang?

        ▪️Step-01    : Install this package

                go get github.com/joho/godotenv

        ▪️Step-02    : Create `.env` file

                APP_PORT=8000
                DB_HOST=localhost
                DB_PORT=5432
                DB_USER=postgres
                DB_PASSWORD=password
                DB_NAME=product_db
                DB_SSLMODE=disable

        ▪️Step-03    : Update `database config` (see on: `config/database.config.go`)
