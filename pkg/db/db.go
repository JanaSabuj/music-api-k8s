package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB return a new db connection
func InitDB(dsn string) *gorm.DB {
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		log.Fatal(err)
	}

	d, err := db.DB()
	if err != nil {
		fmt.Println("mysql init failed")
		log.Fatal(err)
	}

	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
