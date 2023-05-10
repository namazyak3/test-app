package main

import (
	"log"
	"net/http"

	"github.com/namazyak3/test-app/server/config/envloader"
	"github.com/namazyak3/test-app/server/controllers"
	"github.com/namazyak3/test-app/server/database/connector"
)

func main() {
	env := envloader.GetEnv()

	connector.ConnectDB(env)

	http.HandleFunc("/res1", controllers.Handler1)
	http.HandleFunc("/res2", controllers.Handler2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
