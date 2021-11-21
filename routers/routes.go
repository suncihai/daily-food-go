package routers

import (
	"go-daily-food/controllers"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func SetupRoutesForRecords(router *mux.Router) {
	enableCORS(router)
    
	router.HandleFunc("/", home)
	router.HandleFunc("/foodstore", controllers.GetFoodStore).Methods("GET")
	router.HandleFunc("/food", controllers.GetFood).Methods("GET")
	router.HandleFunc("/food", controllers.CreateFood).Methods("POST")
	router.HandleFunc("/food", controllers.EditFood).Methods("PUT")
	router.HandleFunc("/seasoningstore", controllers.GetSeasoningStore).Methods("GET")
	router.HandleFunc("/seasoning", controllers.GetSeasoning).Methods("GET")
	router.HandleFunc("/seasoning", controllers.CreateSeasoning).Methods("POST")
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Contro-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Home Page!")
}


