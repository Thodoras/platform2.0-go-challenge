package audienceRepository

import (
	"database/sql"

	"platform2.0-go-challenge/src/models"
)

type AudienceRepository struct {
	db *sql.DB
}

func NewAudienceRepository(db *sql.DB) *AudienceRepository {
	return &AudienceRepository{db: db}
}

func (a *AudienceRepository) GetAudiences(userID int) ([]models.Audience, error) {
	var audience models.Audience
	result := []models.Audience{}

	rows, err := a.db.Query("SELECT * FROM Audiences WHERE UserID = $1", userID)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(&audience.ID, &audience.UserID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroups, &audience.HoursSpent, &audience.NumOfPurchasesPerMonth)
		if err != nil {
			return result, err
		}
		result = append(result, audience)
	}

	return result, nil
}

func (a *AudienceRepository) GetAudiencesPaginated(userID, limit, offset int) ([]models.Audience, error) {
	var audience models.Audience
	result := []models.Audience{}

	rows, err := a.db.Query("SELECT * FROM Audiences WHERE UserID = $1 ORDER BY id DESC LIMIT $2 OFFSET $3", userID, limit, offset)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(&audience.ID, &audience.UserID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroups, &audience.HoursSpent, &audience.NumOfPurchasesPerMonth)
		if err != nil {
			return result, err
		}
		result = append(result, audience)
	}

	return result, nil
}

func (a *AudienceRepository) AddAudience(audience models.Audience) (int, error) {
	var audienceID int
	row := a.db.QueryRow("INSERT INTO Audiences (UserID, Gender, BirthCountry, AgeGroups, HoursSpent, NumOfPurchasesPerMonth) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", audience.UserID, audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth)
	err := row.Scan(&audienceID)
	if err != nil {
		return 0, err
	}
	return audienceID, nil
}

func (a *AudienceRepository) EditAudience(audience models.Audience) (int64, error) {
	result, err := a.db.Exec("UPDATE Audiences SET Gender=$1, BirthCountry=$2, AgeGroups=$3, HoursSpent=$4, NumOfPurchasesPerMonth = $5 WHERE id=$6 RETURNING id", audience.Gender, audience.BirthCountry, audience.AgeGroups, audience.HoursSpent, audience.NumOfPurchasesPerMonth, audience.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}

func (a *AudienceRepository) DeleteAudience(id int) (int64, error) {
	result, err := a.db.Exec("DELETE FROM Audiences WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, err
}
