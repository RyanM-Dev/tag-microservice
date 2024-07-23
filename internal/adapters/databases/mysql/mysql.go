package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func (m *Mysql) NewDatabase(dsn string) (*gorm.DB, error) {
	log.Println("Database created")
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating database: %v", err)
	}
	return gormDB, nil
}
