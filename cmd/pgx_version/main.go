package main

import (
	"context"
	"db-rnd/model"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const (
	DBConnStr = "postgres://puser:ppassword@localhost:6432/notifyDB?sslmode=disable"
)

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, DBConnStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close(ctx)
	fmt.Println("Conn to DB")
	rows, err := db.Query(ctx, `
	select u.id,
		u.name,
		COALESCE(comp.id, 0),
		COALESCE(comp.id, 0),
		COALESCE(comp.name, 'NO')
	from users as u
	left join companies as comp ON comp.id = u.company_id
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
	users := make([]model.User, 0, 10)

	for rows.Next() {
		var user = model.User{
			Company: &model.Company{},
		}
		if err := rows.Scan(&user.ID, &user.Name, &user.CompanyID, &user.Company.ID, &user.Company.Name); err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}
