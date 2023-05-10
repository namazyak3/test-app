package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/namazyak3/test-app/server/configs"
)

var DB *sql.DB
var err error

func ConnectDB(env configs.EnvType) *sql.DB {
	dbconf := env.MYSQL_USER + ":" + env.MYSQL_PASSWORD + "@tcp(" + env.MYSQL_HOST + ":" + env.MYSQL_TCP_PORT + ")/" + env.MYSQL_DATABASE

	//dbconf := "root:rootpass@tcp(" + env.MYSQL_HOST + ":" + env.MYSQL_TCP_PORT + ")/" + env.MYSQL_DATABASE
	fmt.Printf("アクセス中: %s\n", env.MYSQL_USER)
	DB, err = sql.Open("mysql", dbconf)
	if err != nil {
		log.Println(err.Error())
	}

	defer DB.Close()

	try := 0
	for try <= 10 {
		time.Sleep(3 * time.Second)
		err := DB.Ping()

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

	return DB
}
