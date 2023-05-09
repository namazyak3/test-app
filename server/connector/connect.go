package connector

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/namzyak3/test-app/server/libs/envloader"
)

func ConnectDB(env envloader.EnvType) {
	//dbconf := env.USER + ":" + env.PASSWORD + "@tcp(" + env.HOST + ":" + env.PORT + ")/" + env.DATABASE
	//↑これはまだできてない.

	dbconf := "root:rootpass@tcp(" + env.HOST + ":" + env.PORT + ")/" + env.DATABASE
	fmt.Printf("アクセス中: %s\n", "root")
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		log.Println(err.Error())
	}

	defer db.Close()

	try := 0
	for try <= 10 {
		time.Sleep(3 * time.Second)
		err := db.Ping()

		if err == nil {
			fmt.Println("データベース接続成功")
			break
		}

		if try == 10 {
			fmt.Println("データベース接続失敗")
			log.Println(err.Error())
		}

		try++
	}
}
