package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koheiterajima-bs/sqlboiler-hands-on/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	// DB接続の初期化
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3308)/sqlboiler_demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// boilのDB接続の設定
	boil.SetDB(db)

	// ユーザーの作成
	user := models.User{
		Name:  null.StringFrom("John Doe"),
		Email: null.StringFrom("johndoe@example.com"),
	}

	// ユーザーを挿入
	ctx := context.Background()              // contextの初期化
	err = user.Insert(ctx, db, boil.Infer()) // boil.Infer()は、データベースのテーブルに挿入する際、フィールドとカラムの対応を自動的に推測するオプション
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
