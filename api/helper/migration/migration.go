package migration

import (
	"fmt"
	"go-aether/pkg/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
	SeverPort  string
	GeoKey     string
}

func iniConfig() *Config {
	var rs = new(Config)

	rs = LoadConfig()

	return rs

}

func LoadConfig() *Config {

	var result = new(Config)

	// rs := os.Getenv("HOST")

	if value, found := os.LookupEnv("DB_USER"); found {
		result.DBUsername = value
	}
	if value, found := os.LookupEnv("DB_PASS"); found {
		result.DBPassword = value
	}
	if value, found := os.LookupEnv("DB_PORT"); found {
		result.DBPort = value
	}
	if value, found := os.LookupEnv("DB_HOST"); found {
		result.DBHost = value
	}
	if value, found := os.LookupEnv("DB_NAME"); found {
		result.DBName = value
	}
	if value, found := os.LookupEnv("SERVER_PORT"); found {
		result.SeverPort = value
	}
	// if value, found := os.LookupEnv("GEOAPIFY_KEY"); found {
	// 	result.GeoKey = value
	// }

	return result
}

func InitDB(config *Config) {
	connStr := fmt.Sprintf("host='%s' user='%s' password='%s' dbname='%s' port='%s' sslmode=disable",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}
	MigrationDB(db)

}

func MigrationDB(db *gorm.DB) {

	err := db.AutoMigrate(&entities.PersonDB{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}

	log.Println("Migrasi berhasil!")

}
