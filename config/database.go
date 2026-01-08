package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"api/internal/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var dsn string
	
	// Cek apakah ada environment variable DATABASE_URL (biasanya dari Railway)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		// Parse URL: mysql://user:pass@host:port/dbname
		// GORM DSN: user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		
		u, err := url.Parse(databaseURL)
		if err != nil {
			log.Printf("Warning: Failed to parse DATABASE_URL: %v. Falling back to individual vars.", err)
		} else {
			password, _ := u.User.Password()
			host := u.Host
			// Handle case where host might not have port, though Railway usually includes it
			if !strings.Contains(host, ":") {
				host += ":3306" // Default port
			}
			
			dbname := strings.TrimPrefix(u.Path, "/")
			
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				u.User.Username(),
				password,
				host,
				dbname,
			)
		}
	}

	// Jika DSN masih kosong (tidak ada DATABASE_URL atau gagal parse), gunakan variabel individual
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	// Only User entity exists for now
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Absensi{})
	db.AutoMigrate(&entity.Cuti{})
	
	// db.AutoMigrate(&entity.Visitor{},&entity.Staff{},&entity.Cinema{},&entity.Theater{},&entity.Movie{},&entity.Schedule{},&entity.Booking{},&entity.Payment{},&entity.Log_aktivitas{},&entity.Seat{},&entity.Method{},&entity.Poster{})
	// db.AutoMigrate(&entity.Diskon{})
	// db.AutoMigrate(&entity.DetailMovie{})
	// db.AutoMigrate(&entity.Cast{})
	// db.AutoMigrate(&entity.Region{})

	return db
}

func init() {
	err := godotenv.Load()
	if err != nil {
		// Log but don't fail, as env might be set in system
		log.Println("Warning: Error loading .env file")
	}
}
