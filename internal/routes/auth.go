package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kunalkashyap-1/API_testiing_tool/internal/handlers/auth"
)

func RegisterAuthRoutes(api *mux.Router, db *sql.DB) {
	api.HandleFunc("/auth/register", auth.Register(db)).Methods("POST")
	api.HandleFunc("/auth/login", auth.Login(db)).Methods("POST")
	api.HandleFunc("/auth/refresh", auth.Refresh()).Methods("POST")
}
