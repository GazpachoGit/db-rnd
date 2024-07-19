package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name  string
	Users []User
}

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   *Company
}

func (u User) String() string {
	var sb strings.Builder
	sb.WriteString("UserID: ")
	sb.WriteString(strconv.Itoa(int(u.ID)))
	sb.WriteString("; ")

	sb.WriteString("UserName: ")
	sb.WriteString(u.Name)
	sb.WriteString("; ")

	sb.WriteString("CompanyID: ")
	sb.WriteString(strconv.Itoa(int(u.CompanyID)))
	sb.WriteString("; ")
	if u.CompanyID != 0 {
		sb.WriteString("CompanyName: ")
		sb.WriteString(u.Company.Name)
		sb.WriteString("; ")
	}

	return sb.String()
}

func (c Company) String() string {
	var sb strings.Builder
	sb.WriteString("CompID: ")
	sb.WriteString(strconv.Itoa(int(c.ID)))
	sb.WriteString("; ")

	sb.WriteString("CompName: ")
	sb.WriteString(c.Name)
	sb.WriteString("; ")

	sb.WriteString("Users: [")
	for _, u := range c.Users {
		sb.WriteString(u.Name)
		sb.WriteString(",")
	}
	sb.WriteString("]")

	return sb.String()
}
