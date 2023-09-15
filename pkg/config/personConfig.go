package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=dpg-ck2bmcfhdsdc73a05o8g-a port=5432 user=record_h59m_user dbname=RECORD password=FZ9EuZHcJH1iX8qwkOPZSPP0JCQXG6lO sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
