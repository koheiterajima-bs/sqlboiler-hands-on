package main

import (
	"fmt"
	"log"

	"github.com/koheiterajima-bs/sqlboiler-hands-on/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := boil.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// ユーザーの新規作成
	user := models.User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	err = user.Insert(db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	// ユーザーの取得
	users, err := models.Users(qm.Where("name=?", "John Doe")).All(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("User: %s, Email: %s\n", user.Name, user.Email)
	}
}
