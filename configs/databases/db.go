package databases

import (
	"database/sql"
	"log"

	"github.com/yehpattana/api-yehpattana/configs"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	sqlServerMigrate "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitDB(config configs.DbConfigInterface) (*gorm.DB, error) {
	// connect to database
	db, err := gorm.Open(sqlserver.Open(config.Url()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Run migrations
	err = runMigrations(config.Url())
	if err != nil {
		log.Fatalf("Error running migrations: %v\n", err)
	}

	log.Println("Migrations applied successfully!")

	log.Println("Connected to database")
	checkVersion(db)
	return db, err
}

func checkVersion(db *gorm.DB) {
	if db == nil {
		log.Println("Database instance is nil")
		return
	}

	var version string

	err := db.Raw("SELECT @@version").Scan(&version).Error
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	log.Printf("%s\n", version)
}

func runMigrations(dbUrl string) error {
	db, err := sql.Open("sqlserver", dbUrl)
	if err != nil {
		log.Println(err)
	}
	driver, err := sqlServerMigrate.WithInstance(db, &sqlServerMigrate.Config{})
	if err != nil {
		log.Println(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlserver", driver)
	if err != nil {
		log.Println(err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
	}
	return nil
}
