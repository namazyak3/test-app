package main

import (
	"log"
	"net/http"

	"github.com/namazyak3/test-app/server/config"
	"github.com/namazyak3/test-app/server/controllers"
	"github.com/namazyak3/test-app/server/database"
)

func main() {
	env := config.GetEnv()

	database.ConnectDB(env)

	http.HandleFunc("/res1", controllers.Res1)
	http.HandleFunc("/res2", controllers.Res2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
