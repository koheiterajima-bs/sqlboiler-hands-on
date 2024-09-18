package main

import (
	"context"
	"fmt"
	"log"

	"github.com/koheiterajima-bs/sqlboiler-hands-on/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/volatiletech/null.v8"
)

func main() {
	// DB接続の初期化
	db := boil.GetDB()

	// ユーザーの作成
	user := models.User{
		Name:  null.StringFrom("John Doe"),
		Email: null.StringFrom("johndoe@example.com"),
	}

	// ユーザーを挿入
	ctx := context.Background()
	err := user.Insert(ctx, db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	// ユーザーの検索
	users, err := models.Users(qm.Where("name=?", "John Doe")).All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の出力
	for _, u := range users {
		fmt.Println(u.Name.String, u.Email.String)
	}
}
