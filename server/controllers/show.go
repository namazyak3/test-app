package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/namazyak3/test-app/server/database"
	"github.com/namazyak3/test-app/server/model"
)

func Show(w http.ResponseWriter, _ *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM test_table")
	if err != nil {
		log.Println(err.Error())
	}

	var tb model.Test_table

	err = rows.Scan(&tb.ID, &tb.NAME, &tb.DETAIL)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(rows)
}
