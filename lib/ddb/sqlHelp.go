package ddb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDb(source string) (dao *gorm.DB, err error) {
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
	return
}
