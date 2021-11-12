package main

import (
	"database/sql"
	"fmt"
	"go-daily-food/db"
	_ "go-daily-food/docs"
	"go-daily-food/routers"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

var err error

func main() {
    db.DB, err = sql.Open("mysql", routers.ConnectionString)
	if(err != nil) {
		panic(err.Error())
	}
	fmt.Println("successfully connected to database.")

	router := mux.NewRouter()
	routers.SetupRoutesForRecords(router)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	port := ":8000"
	server := &http.Server{
		Handler: router,
		Addr: port,
		WriteTimeout:  15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}


