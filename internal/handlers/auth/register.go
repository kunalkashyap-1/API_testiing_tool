package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/kunalkashyap-1/API_testiing_tool/pkg/auth"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/models"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		hashedPassword, err := auth.HashPassword(req.Password)

		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		user := models.User{
			Username:     req.Username,
			PasswordHash: hashedPassword,
			Email:        req.Email,
			CreatedAt:    time.Now(),
		}

		query := `INSERT INTO users (username, password_hash, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
		err = db.QueryRow(query, user.Username, user.PasswordHash, user.Email, user.CreatedAt).Scan(&user.ID)

		if err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
