package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB return a new db connection
func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	d, err := db.DB()
	if err != nil {
		log.Fatal("mysql init failed", err)
	}

	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
