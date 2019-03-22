package sql

import (
	"github.com/jinzhu/gorm"
)

var gdb *gorm.DB

func InitialDatabase(driver, url string) error {
	return setDB(driver, url)
}

func InitialDatabaseWithOption(driver, url string, option Options) error {
	if err := setDB(driver, url); err != nil {
		return err
	}
	SetOptions(option)
	return nil
}

type Options struct {
	LOGMODE       bool
	MAXIDLE       int
	MAXCONNECTION int
}

func SetOptions(option Options) {
	gdb.LogMode(option.LOGMODE)
	gdb.DB().SetMaxIdleConns(option.MAXIDLE)
	gdb.DB().SetMaxOpenConns(option.MAXCONNECTION)
}

func setDB(driver, sqlurl string) error {
	db, err := gorm.Open(driver, sqlurl)
	if err != nil {
		return err
	}

	gdb = db
	return nil
}
