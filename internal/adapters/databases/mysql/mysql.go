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

var mysqlDB Mysql

func (m *Mysql) NewMysqlDatabase(dsn string) (*gorm.DB, error) {
	log.Println("Database created")
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating database: %v", err)
	}
	return gormDB, nil
}

func (m *Mysql) AutoMigrate() error {
	if err := m.db.AutoMigrate(&models.GormTag{}, &models.GormTaxonomy{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}
