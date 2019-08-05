package repositories

import (
	"platform2.0-go-challenge/src/helpers/drivers"
	"platform2.0-go-challenge/src/models"
)

func GetCharts(userID int, onlyFavourites bool) ([]models.Chart, error) {
	var chart models.Chart
	result := []models.Chart{}

	rows, err := drivers.DB.Query("SELECT * FROM Charts WHERE UserID = $1 "+getQuerySuffix(onlyFavourites), userID)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(
			&chart.ID,
			&chart.UserID,
			&chart.Title,
			&chart.AxisXTitle,
			&chart.AxisYTitle,
			&chart.Data,
			&chart.Favourite,
		)
		if err != nil {
			return result, err
		}
		result = append(result, chart)
	}

	return result, nil
}

func GetChartsPaginated(userID, limit, offset int, onlyFavourites bool) ([]models.Chart, error) {
	var chart models.Chart
	result := []models.Chart{}

	rows, err := drivers.DB.Query("SELECT * FROM Charts WHERE UserID = $1 "+getQuerySuffix(onlyFavourites)+" ORDER BY id DESC LIMIT $2 OFFSET $3", userID, getQuerySuffix(onlyFavourites), limit, offset)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(
			&chart.ID,
			&chart.UserID,
			&chart.Title,
			&chart.AxisXTitle,
			&chart.AxisYTitle,
			&chart.Data,
			&chart.Favourite,
		)
		if err != nil {
			return result, err
		}
		result = append(result, chart)
	}

	return result, nil
}

func AddChart(chart models.Chart) (int, error) {
	var chartID int
	row := drivers.DB.QueryRow("INSERT INTO Charts (UserID, Title, AxisXTitle, AxisYTitle, Data, Favourite) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", chart.UserID, chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data, chart.Favourite)
	err := row.Scan(&chartID)
	if err != nil {
		return 0, err
	}
	return chartID, nil
}

func EditChart(chart models.Chart) (int64, error) {
	result, err := drivers.DB.Exec("UPDATE Charts SET Title=$1, AxisXTitle=$2, AxisYTitle=$3, Data=$4, Favourite=$5 WHERE id=$6 RETURNING id", chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data, chart.Favourite, chart.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}

func DeleteChart(id int) (int64, error) {
	result, err := drivers.DB.Exec("DELETE FROM Charts WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, err
}
