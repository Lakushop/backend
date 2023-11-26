package router

import (
	"backend/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTION")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTION")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTION")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTION")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTION")

	return router
}
