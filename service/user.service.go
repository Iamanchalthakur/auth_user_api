package service

import (
	"database/sql"

	"github.com/tetsing/models" // Update this to the correct import path
)

func GetAllUsersFromDB(db *sql.DB) ([]models.User, error) {
	query := `SELECT u.id as user_id, title as role_name, u.username, u.email, u.role_id 
			  FROM users as u 
			  LEFT JOIN role as r ON u.role_id = r.id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.ROLE_NAME, &user.USERNAME, &user.EMAIL, &user.ROLE_ID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
