package repositories

import (
	"platform2.0-go-challenge/helpers/drivers"
	"platform2.0-go-challenge/models"
)

func GetCharts(id int) ([]models.Chart, error) {
	var chart models.Chart
	result := []models.Chart{}

	rows, err := drivers.DB.Query("SELECT * FROM Charts WHERE UserID = $1", id)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		err := rows.Scan(&chart.ID, &chart.UserID, &chart.Title, &chart.AxisXTitle, &chart.AxisYTitle, &chart.Data)
		if err != nil {
			return result, err
		}
		result = append(result, chart)
	}

	return result, nil
}

func AddChart(chart models.Chart) (int, error) {
	var chartID int
	row := drivers.DB.QueryRow("INSERT INTO Charts (UserID, Title, AxisXTitle, AxisYTitle, Data) VALUES ($1, $2, $3, $4, $5) RETURNING ID", chart.UserID, chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data)
	err := row.Scan(&chartID)
	if err != nil {
		return 0, err
	}
	return chartID, nil
}

func EditChart(chart models.Chart) (int64, error) {
	result, err := drivers.DB.Exec("UPDATE Charts SET Title=$1, AxisXTitle=$2, AxisYTitle=$3, Data=$4 WHERE id=$5 RETURNING id", chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data, chart.ID)
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
