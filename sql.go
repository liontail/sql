package sql

import (
	"github.com/jinzhu/gorm"
)

var gdb *gorm.DB

type Options struct {
	LOGMODE       bool
	MAXIDLE       int
	MAXCONNECTION int
}

func InitialDatabase(driver, url string) error {
	db, err := Connect(driver, url)
	gdb = db
	return err
}

func InitialDatabaseWithOption(driver, url string, option Options) error {
	db, err := Connect(driver, url)
	if err != nil {
		return err
	}

	SetOptions(db, option)
	gdb = db
	return nil
}

func SetOptions(db *gorm.DB, option Options) {
	db.LogMode(option.LOGMODE)
	db.DB().SetMaxIdleConns(option.MAXIDLE)
	db.DB().SetMaxOpenConns(option.MAXCONNECTION)
}

func Connect(driver, sqlurl string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, sqlurl)
	return db, err
}

func GetDB() *gorm.DB {
	return gdb
}
