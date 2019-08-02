package repositories

import (
	"strings"

	"platform2.0-go-challenge/helpers/errorutils"
	"platform2.0-go-challenge/models"
)

func AddUser(user models.User) (int, error) {
	var userID int
	row := DB.QueryRow("INSERT INTO Users (Name, Password) VALUES ($1, $2) RETURNING id", user.Name, user.Password)
	err := row.Scan(&userID)
	if err != nil {
		if strings.Contains(err.Error(), errorutils.UniqueConstrainViolationString) {
			return 0, errorutils.NewUniqueConstrainViolation("User with same name already exists")
		}

		return 0, err
	}

	return userID, nil
}
