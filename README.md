Hello, Full Stack Web Development Lecturers✌️!

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