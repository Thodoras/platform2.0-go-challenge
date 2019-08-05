package repositories

import (
	"platform2.0-go-challenge/src/helpers/drivers"
	"platform2.0-go-challenge/src/models"
)

func GetAudiences(userID int, onlyFavourites bool) ([]models.Audience, error) {
	var audience models.Audience
	result := []models.Audience{}

	rows, err := drivers.DB.Query("SELECT * FROM Audiences WHERE UserID = $1"+getQuerySuffix(onlyFavourites), userID)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(
			&audience.ID,
			&audience.UserID,
			&audience.Gender,
			&audience.BirthCountry,
			&audience.AgeGroups,
			&audience.HoursSpent,
			&audience.NumOfPurchasesPerMonth,
			&audience.Favourite,
		)
		if err != nil {
			return result, err
		}
		result = append(result, audience)
	}

	return result, nil
}

func GetAudiencesPaginated(userID, limit, offset int, onlyFavourites bool) ([]models.Audience, error) {
	var audience models.Audience
	result := []models.Audience{}

	rows, err := drivers.DB.Query("SELECT * FROM Audiences WHERE UserID = $1 "+getQuerySuffix(onlyFavourites)+" ORDER BY id DESC LIMIT $2 OFFSET $3", userID, getQuerySuffix(onlyFavourites), limit, offset)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(
			&audience.ID,
			&audience.UserID,
			&audience.Gender,
			&audience.BirthCountry,
			&audience.AgeGroups,
			&audience.HoursSpent,
			&audience.NumOfPurchasesPerMonth,
			&audience.Favourite,
		)
		if err != nil {
			return result, err
		}
		result = append(result, audience)
	}

	return result, nil
}

func AddAudience(audience models.Audience) (int, error) {
	var audienceID int
	row := drivers.DB.QueryRow("INSERT INTO Audiences (UserID, Gender, BirthCountry, AgeGroups, HoursSpent, NumOfPurchasesPerMonth) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING ID", audience.UserID, audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth, audience.Favourite)
	err := row.Scan(&audienceID)
	if err != nil {
		return 0, err
	}
	return audienceID, nil
}

func EditAudience(audience models.Audience) (int64, error) {
	result, err := drivers.DB.Exec("UPDATE Audiences SET Gender=$1, BirthCountry=$2, AgeGroups=$3, HoursSpent=$4, NumOfPurchasesPerMonth = $5, Favourite = $6 WHERE id=$7 RETURNING id", audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth, audience.Favourite, audience.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}

func DeleteAudience(id int) (int64, error) {
	result, err := drivers.DB.Exec("DELETE FROM Audiences WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, err
}
