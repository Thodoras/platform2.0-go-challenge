package repositories

import "platform2.0-go-challenge/models/assets"

func GetAudiences(id int) ([]assets.Audience, error) {
	var audience assets.Audience
	result := []assets.Audience{}

	rows, err := DB.Query("SELECT * FROM Audiences WHERE UserID = $1", id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&audience.ID, &audience.UserID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroups, &audience.HoursSpent, &audience.NumOfPurchasesPerMonth)
		if err != nil {
			return nil, err
		}
		result = append(result, audience)
	}

	return result, nil
}

func AddAudience(audience assets.Audience) (int, error) {
	var audienceID int
	row := DB.QueryRow("INSERT INTO Audiences (UserID, Gender, BirthCountry, AgeGroups, HoursSpent, NumOfPurchasesPerMonth) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", audience.UserID, audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth)
	err := row.Scan(&audienceID)
	if err != nil {
		return 0, err
	}
	return audienceID, nil
}

func EditAudience(audience assets.Audience) (int64, error) {
	result, err := DB.Exec("UPDATE Audiences SET Gender=$1, BirthCountry=$2, AgeGroups=$3, HoursSpent=$4, NumOfPurchasesPerMonth = $5 WHERE id=$6 RETURNING id", audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth, audience.ID)
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
	result, err := DB.Exec("DELETE FROM Audiences WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, err
}
