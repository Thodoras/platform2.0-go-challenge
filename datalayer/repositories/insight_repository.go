package repositories

import "platform2.0-go-challenge/models/assets"

func GetInsights(id int) ([]assets.Insight, error) {
	var insight assets.Insight
	result := []assets.Insight{}

	rows, err := DB.Query("SELECT * FROM Insights WHERE UserID = $1", id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&insight.ID, &insight.UserID, &insight.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, insight)
	}

	return result, nil
}

func AddInsight(insight assets.Insight) (int, error) {
	var insightID int
	row := DB.QueryRow("INSERT INTO Insights (UserID, Text) VALUES ($1, $2) RETURNING ID", insight.UserID, insight.Text)
	err := row.Scan(&insightID)
	if err != nil {
		return 0, err
	}
	return insightID, nil
}

func EditInsight(insight assets.Insight) (int64, error) {
	result, err := DB.Exec("UPDATE Insights SET Text=$1 WHERE id=$2 RETURNING id", insight.Text, insight.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}
