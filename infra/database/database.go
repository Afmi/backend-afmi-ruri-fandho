package database

import (
	"log"
	"user-test/config"
	models "user-test/models" // pastikan untuk mengganti dengan path model yang sesuai

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	DB  *gorm.DB
	err error
)

// DbConnection untuk menghubungkan ke database dan melakukan migrasi
func DbConnection() error {
	masterDSN, replicaDSN := config.DatabaseConfig()

	// Gunakan log level sesuai kebutuhan
	loglevel := logger.Info

	var err error
	DB, err = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to master DB: %v", err)
	}

	// Tambahkan replica jika DSN-nya tidak kosong
	if replicaDSN != "" {
		err = DB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
		if err != nil {
			log.Fatalf("❌ Failed to register replica: %v", err)
		}
	}

	// UUID extension
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("❌ Failed to enable uuid-ossp: %v", err)
	}

	// Migrasi model-model
	if err := DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Cart{},
		&models.Product{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate error: %v", err)
	}

	log.Println("📦 Master DB connected & models migrated")
	return nil
}
