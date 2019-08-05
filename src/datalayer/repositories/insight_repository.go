package repositories

import (
	"platform2.0-go-challenge/src/helpers/drivers"
	"platform2.0-go-challenge/src/models"
)

func GetInsights(id int, onlyFavourites bool) ([]models.Insight, error) {
	var insight models.Insight
	result := []models.Insight{}

	rows, err := drivers.DB.Query("SELECT * FROM Insights WHERE UserID = $1 "+getQuerySuffix(onlyFavourites), id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&insight.ID,
			&insight.UserID,
			&insight.Text,
			&insight.Favourite,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, insight)
	}

	return result, nil
}

func GetInsightsPaginated(userID, limit, offset int, onlyFavourites bool) ([]models.Insight, error) {
	var insight models.Insight
	result := []models.Insight{}

	rows, err := drivers.DB.Query("SELECT * FROM Insights WHERE UserID = $1 "+getQuerySuffix(onlyFavourites)+" ORDER BY id DESC LIMIT $2 OFFSET $3", userID, getQuerySuffix(onlyFavourites), limit, offset)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(
			&insight.ID,
			&insight.UserID,
			&insight.Text,
			&insight.Favourite,
		)
		if err != nil {
			return result, err
		}
		result = append(result, insight)
	}

	return result, nil
}

func AddInsight(insight models.Insight) (int, error) {
	var insightID int
	row := drivers.DB.QueryRow("INSERT INTO Insights (UserID, Text, Favourite) VALUES ($1, $2, $3) RETURNING ID", insight.UserID, insight.Text, insight.Favourite)
	err := row.Scan(&insightID)
	if err != nil {
		return 0, err
	}
	return insightID, nil
}

func EditInsight(insight models.Insight) (int64, error) {
	result, err := drivers.DB.Exec("UPDATE Insights SET Text=$1, Favourite=$2 WHERE id=$3 RETURNING id", insight.Text, insight.Favourite, insight.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}

func DeleteInsight(id int) (int64, error) {
	result, err := drivers.DB.Exec("DELETE FROM Insights WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, err
}
