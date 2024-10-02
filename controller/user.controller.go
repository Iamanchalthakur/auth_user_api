package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tetsing/models"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Println("Invalid input:", err)
			return
		}

		query := `INSERT INTO users (username, email, password, role_id) VALUES ($1, $2, $3, $4) RETURNING id`

		err = db.QueryRow(query, user.USERNAME, user.EMAIL, user.PASSWORD, user.ROLE_ID).Scan(&user.ID)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			log.Println("Insert query error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User created successfully",
			"user_id": user.ID,
		})
	}
}
