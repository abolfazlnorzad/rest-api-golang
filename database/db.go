package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id     int
	Name   string
	Family string
}

type DBHandler struct {
	db *gorm.DB
}

func CreateConnection() (*DBHandler, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &DBHandler{
		db: db,
	}, nil
}

func (handler *DBHandler) GetAllUsers() []User {
	var users []User
	res, err := handler.db.Find(&users).Rows()
	if err != nil {
		log.Fatalln(err)
	}

	for res.Next() {
		err := handler.db.ScanRows(res, &users)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return users
}

func (handler *DBHandler) FindUserByName(name string) User {

	user := User{}
	queryResult, _ := handler.db.Find(&user, "name = ?", name).Rows()
	for queryResult.Next() {
		err := handler.db.ScanRows(queryResult, &user)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return user
}

func (handler *DBHandler) UpdateUserById(id int) User {
	handler.db.Table("users").Where("id = ?", id).Update("family", "FM07")
	return handler.FindUserByName("ilia")
}

func (handler *DBHandler) DeleteUserById(id int) {
	var user User
	handler.db.Table("users").Where("id = ?", id).Delete(&user)
}
