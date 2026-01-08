package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"api/internal/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	var dsn string
	
	// Cek apakah ada environment variable DATABASE_URL (biasanya dari Railway)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		// Parse URL: mysql://user:pass@host:port/dbname
		u, err := url.Parse(databaseURL)
		if err != nil {
			log.Printf("Warning: Failed to parse DATABASE_URL: %v. Falling back to individual vars.", err)
		} else {
			password, _ := u.User.Password()
			host := u.Host
			if !strings.Contains(host, ":") {
				host += ":3306"
			}
			
			dbname := strings.TrimPrefix(u.Path, "/")
			
			// Tambahkan parameter untuk performa dan timeout
			// interpolateParams=true: Mengurangi round-trip ke DB
			// timeout=10s: Batas waktu koneksi awal
			// readTimeout=30s: Batas waktu baca query
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true&timeout=10s&readTimeout=30s&writeTimeout=30s",
				u.User.Username(),
				password,
				host,
				dbname,
			)
			
			log.Printf("Connecting to DB Host: %s, DB Name: %s", host, dbname)
		}
	}

	// Jika DSN masih kosong
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true&timeout=10s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		log.Println("Connecting to DB using individual env vars")
	}

	// Configure GORM Logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Log jika query > 200ms
			LogLevel:                  logger.Info,            
			IgnoreRecordNotFoundError: true,                   
			ParameterizedQueries:      false,                  
			Colorful:                  false,                  
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		PrepareStmt: true, // Cache prepared statements untuk performa
	})
	if err != nil {
		log.Println("DB connection failed:", err)
		return nil
	}

	// Connection Pooling Configuration
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get underlying sql.DB:", err)
		return nil
	}

	// Tuning pool untuk environment container/serverless
	sqlDB.SetMaxIdleConns(5)        // Jangan terlalu banyak idle connection
	sqlDB.SetMaxOpenConns(50)       // Batas maksimal koneksi terbuka
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Refresh koneksi tiap 5 menit untuk menghindari stale connection

	// Auto Migrate (Hanya di local/dev, atau jika ENV != production untuk mempercepat start)
	// Di production, sebaiknya migrasi dijalankan manual atau via pipeline
	if os.Getenv("ENV") != "production" {
		log.Println("Starting AutoMigrate...")
		start := time.Now()
		
		db.AutoMigrate(&entity.User{})
		db.AutoMigrate(&entity.Absensi{})
		db.AutoMigrate(&entity.Cuti{})
		
		log.Printf("AutoMigrate finished in %v", time.Since(start))
	} else {
		log.Println("Skipping AutoMigrate in production environment")
	}

	return db
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file (ok in production)")
	}
}
