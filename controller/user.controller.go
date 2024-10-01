package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct to represent the user entity
type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
}

// CreateUser inserts a new user into the "users" table
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Invalid input:", err)
		return
	}

	// Insert the new user into the users table
	query := `INSERT INTO users (username, email, password, role_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(query, user.Username, user.Email, user.Password, user.RoleID).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Println("Insert query error:", err)
		return
	}

	// Respond with the created user's ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}
