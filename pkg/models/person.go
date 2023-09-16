package models

import (
	"log"

	"github.com/Zekeriyyah/stagetwo/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name    string `gorm:"" json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (p *User) CreateUser() *User {
	//db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (User, *gorm.DB, error) {
	var user User
	db := db.Where("ID = ?", Id).Find(&user)
	if db.Error != nil {
		log.Println("Id not found")
		return User{}, db, nil
	}
	return user, db, db.Error
}

func DeleteUser(Id int64) error {
	var user User
	err := db.Where("ID = ?", Id).Delete(&user).Error
	if err != nil {
		log.Println("Error while deleting: ", err)
		return err
	}
	return nil
}
