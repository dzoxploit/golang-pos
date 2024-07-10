cara untuk running project

pertama jalankan perintah go mod download

setelah itu ubah config.go dan ubah username password dan nama dbnya sesuai dengan environtment yang dipakai

dbURI := "<user>:<password>@tcp(localhost:3306)/<dbname>?charset=utf8&parseTime=True&loc=Local"

setelah itu jalankan go run main.go

Namun Jika Menggunakan docker ubah kode config.go seperti dibawah ini

package config

import (
"fmt"
"go-pos/models"
"os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

)

func NewDB() (\*gorm.DB, error) {
// Get database connection details from environment variables
dbUser := os.Getenv("DB_USER")
dbPassword := os.Getenv("DB_PASSWORD")
dbHost := os.Getenv("DB_HOST")
dbPort := os.Getenv("DB_PORT")
dbName := os.Getenv("DB_NAME")

    // Construct the database URI
    dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    	dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := gorm.Open("mariadb", dbURI)
    if err != nil {
    	return nil, err
    }

    // Enable debug mode during development for logging SQL queries
    db.LogMode(true)

    // AutoMigrate will automatically create the database tables based on the models
    db.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{})

    return db, nil

}

dan sesuaikan db_port db_user, db_password dan db_name di file .env dan docker-compose.yaml

Setelah itu jalankan perintah

docker-compose build
docker-compose up

Setalah itu import database di phpmyadmin atau sql file namun jika menggunakan docker import database dibagian image docker-mariadb untuk akses database

setelah itu import postman collection untuk menjalankan api yang sudah terdokumentasi terima kasih :)
