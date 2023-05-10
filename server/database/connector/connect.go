package connector

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/namazyak3/test-app/server/config/envloader"
)

func ConnectDB(env envloader.EnvType) {
	//dbconf := env.MYSQL_USER + ":" + env.MYSQL_PASSWORD + "@tcp(" + env.MYSQL_HOST + ":" + env.MYSQL_PORT + ")/" + env.MYSQL_DATABASE
	//↑これはまだできてない.

	dbconf := "root:rootpass@tcp(" + env.MYSQL_HOST + ":" + env.MYSQL_TCP_PORT + ")/" + env.MYSQL_DATABASE
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
