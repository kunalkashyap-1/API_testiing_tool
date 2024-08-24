package auth

import (
	"database/sql"
	"encoding/json"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/auth"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/models"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		var user models.User
		query := `SELECT id, password_hash FROM users WHERE username=$1`
		err := db.QueryRow(query, req.Username).Scan(&user.ID, &user.PasswordHash)
		if err == sql.ErrNoRows || !auth.CheckPasswordHash(req.Password, user.PasswordHash) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateJWT(req.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}
