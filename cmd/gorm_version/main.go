package main

import (
	"db-rnd/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBConnStr = "postgres://puser:ppassword@localhost:6432/notifyDB?sslmode=disable"
)

func main() {
	db, err := gorm.Open(postgres.Open(DBConnStr), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Conn to DB")

	err = db.AutoMigrate(&model.Company{}, &model.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	CreateSomeRecords(db)

	GetUsersAndCompaniesWithJoin(db)
	GetUsersAndCompanyWithPreload(db)
	GetCompaniesWithUsers(db)

	fmt.Println("Complete!")
}

// Join - executed as LEFT JOIN
func GetUsersAndCompaniesWithJoin(db *gorm.DB) {
	fmt.Println("GetUsersAndCompaniesWithJoin: ")
	var users = []model.User{}

	result := db.Joins("Company").Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// Preload - is executed in separate select
func GetUsersAndCompanyWithPreload(db *gorm.DB) {
	fmt.Println("GetUsersAndCompanyWithPreload: ")
	var users = []model.User{}

	result := db.Preload("Company").Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

func GetCompaniesWithUsers(db *gorm.DB) {
	fmt.Println("GetCompaniesWithUsers: ")
	var companies []model.Company
	err := db.Model(&model.Company{}).Preload("Users").Find(&companies).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range companies {
		fmt.Printf("%+v\n", c)
	}
}

func CreateSomeRecords(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		var Comp = &model.Company{
			Name: "Comp111",
		}
		var users = []model.User{
			{
				Name: "u1",
				Company: &model.Company{
					Name: "Comp2",
				},
			},
			{
				Name:    "u2",
				Company: Comp,
			},
			{
				Name:    "u3",
				Company: Comp,
			},
			{
				Name: "user without comp",
			},
		}
		db.Create(&users)
	}
}
