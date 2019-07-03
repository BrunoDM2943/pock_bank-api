package infra

import (
	"github.com/BrunoDM2943/pock_bank-api/domain"
	"github.com/jinzhu/gorm"

	//SQLlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//InitDB initializes sqlite3 db
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Account{})
	return db
}
