package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tetsing/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Println("Invalid input:", err)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.DefaultCost)

		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			log.Print("password hashing failed:", err)
		}
		user.PASSWORD = string(hashedPassword)

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
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.Login
		fmt.Println("=========login==========", login)

		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Println("Invalid input:", err)
			return
		}

		var user models.User
		query := `SELECT * FROM users WHERE email = $1`
		err = db.QueryRow(query, login.EMAIL).Scan(&user.ID, &user.USERNAME, &user.EMAIL, &user.PASSWORD, &user.ROLE_ID)

		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			log.Println("User not found:", err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(login.PASSWORD))
		fmt.Println("======= db passowrd======", user.PASSWORD)
		fmt.Println("======= payload passowrd======", user.PASSWORD)
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			log.Println("Invalid password:", err)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			log.Println("Error generating token:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "Login successful",
			"user_id":  user.ID,
			"username": user.USERNAME,
			"email":    user.EMAIL,
			"token":    tokenString,
		})
	}
}

func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := `SELECT u.id as user_id,title as role_name, u.username, u.email, u.role_id from users as u left join role as r on 
u.role_id = r.id `

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Could not get users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User

		for rows.Next() {
			var user models.User
			err := rows.Scan(&user.ID, &user.ROLE_NAME, &user.USERNAME, &user.EMAIL, &user.ROLE_ID)
			if err != nil {
				fmt.Println("---error---", err)
				http.Error(w, "Error reading users", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
			fmt.Println("====user=======", user)
			fmt.Println("====users=======", users)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func DeleteAllUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=====dummy-=====")

		query := `DELETE from users`
		_, err := db.Exec(query)
		if err != nil {
			http.Error(w, "Failed to delete all userws", http.StatusInternalServerError)
			log.Println("=====error in query=====", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "All users deleted successfully",
		})
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		fmt.Println("=====dummy-=====")

		query := `UPDATE users SET username = $2, email = $3, password = $4, role_id = $5 WHERE id = $1`

		// Send query to database.
		_, err := db.Exec(query, user.ID, user.USERNAME, user.EMAIL, user.PASSWORD, user.ROLE_ID)
		if err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			fmt.Println("Update query error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User updated successfully",
		})

	}
}
