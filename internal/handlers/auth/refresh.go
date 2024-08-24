package auth

import (
	"encoding/json"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/auth"
	"net/http"
)

type RefreshRequest struct {
	Token string `json:"token"`
}

func Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		claims, err := auth.ValidateToken(req.Token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		newToken, err := auth.GenerateJWT(claims.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": newToken,
		})
	}
}
