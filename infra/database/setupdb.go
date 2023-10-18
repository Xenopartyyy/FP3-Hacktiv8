package setupdb

import (
	"FP3-Hacktiv8/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var User = model.User{}
var Category = model.Category{}
var Task = model.Task{}
var err error

func seedAdmin() {
	admin := &model.User{
		Full_name: "superadmin",
		Email:     "superadmin@hacktiv8.com",
		Password:  "superadmin",
		Role:      "admin",
	}

	admin.HashPassword()

	DB.Create(admin)
	if err != nil {
		log.Fatalf("Gagal menghash password: %v", err)
	}

	log.Println("Akun admin berhasil di seed")

}

func ConnDB() {
	godotenv.Load(".env")
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Category, &User, &Task)

	if DB.Migrator().HasTable(&User) {
		var userCount int64
		if err := DB.Model(&User).Count(&userCount).Error; err != nil {
			log.Fatalf("Gagal memeriksa tabel user: %v", err)
		}

		if userCount == 0 {
			seedAdmin()
			if err != nil {
				log.Fatalf("Gagal melakukad seed admin: %v", err)
			}
		}
	}

}
