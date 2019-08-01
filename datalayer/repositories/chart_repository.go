package repositories

import "platform2.0-go-challenge/models/assets"

func GetCharts(id int) ([]assets.Chart, error) {
	var chart assets.Chart
	result := []assets.Chart{}

	rows, err := DB.Query("SELECT * FROM Charts WHERE UserID = $1", id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&chart.ID, &chart.UserID, &chart.Title, &chart.AxisXTitle, &chart.AxisYTitle, &chart.Data)
		if err != nil {
			return nil, err
		}
		result = append(result, chart)
	}

	return result, nil
}

func AddChart(chart assets.Chart) (int, error) {
	var chartID int
	row := DB.QueryRow("INSERT INTO Charts (UserID, Title, AxisXTitle, AxisYTitle, Data) VALUES ($1, $2, $3, $4, $5) RETURNING ID", chart.UserID, chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data)
	err := row.Scan(&chartID)
	if err != nil {
		return 0, err
	}
	return chartID, nil
}

func EditChart(chart assets.Chart) (int64, error) {
	result, err := DB.Exec("UPDATE Charts SET Title=$1, AxisXTitle=$2, AxisYTitle=$3, Data=$4 WHERE id=$5 RETURNING id", chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data, chart.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdate, err
}
