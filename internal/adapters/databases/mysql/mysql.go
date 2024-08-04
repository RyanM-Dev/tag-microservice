package mysql

import (
	"fmt"
	"log"
	"tagMicroservice/internal/adapters/databases/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func (m *Mysql) NewMysqlDatabase(dsn string) error {
	dsnWithoutDB := "root:831374@tcp(127.0.0.1:3306)/"
	dbName := "tag-microservice"
	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to MySQL server: %v", err)
	}
	err = tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName)).Error
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to MySQL server: %v", err)
	}

	m.db = gormDB
	log.Println("Database created")
	return nil
}
func (m *Mysql) GetDB() *gorm.DB {
	return m.db
}

func (m *Mysql) AutoMigrate() error {
	if err := m.db.AutoMigrate(&models.GormTag{}, &models.GormTaxonomy{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}
