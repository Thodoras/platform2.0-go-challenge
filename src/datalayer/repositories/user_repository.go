package repositories

import (
	"database/sql"
	"strings"

	"platform2.0-go-challenge/src/helpers/drivers"
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

func GetUserByName(name string) (*models.User, error) {
	var user models.User

	row := drivers.DB.QueryRow("SELECT * FROM Users WHERE Name = $1", name)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorutils.NewInvalidRequest("Invalid credentials.")
		}
		return nil, err
	}

	return &user, nil
}

func AddUser(user models.User) (int, error) {
	var userID int
	row := drivers.DB.QueryRow("INSERT INTO Users (Name, Password) VALUES ($1, $2) RETURNING id", user.Name, user.Password)
	err := row.Scan(&userID)
	if err != nil {
		if strings.Contains(err.Error(), errorutils.UniqueConstrainViolationString) {
			return 0, errorutils.NewUniqueConstrainViolation("User with given name already exists")
		}

		return 0, err
	}

	return userID, nil
}
