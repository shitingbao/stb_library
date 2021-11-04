package ddb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysqlClient(source string) (*gorm.DB, error) {
	d, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := d.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	return d, nil
}
