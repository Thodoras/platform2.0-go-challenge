package repositories

import (
	"database/sql"

	"platform2.0-go-challenge/models/assets"
)

var DB *sql.DB

func GetAudiences(id string) ([]assets.Audience, error) {
	var audience assets.Audience
	result := []assets.Audience{}

	rows, err := DB.Query("SELECT * FROM Audiences")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&audience.ID, &audience.UserID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroups, &audience.NumOfPurchasesPerMonth)
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

func AddChart(chart assets.Chart) (int, error) {
	var chartID int
	row := DB.QueryRow("INSERT INTO Charts (UserID, Title, AxisXTitle, AxisYTitle, Data) VALUES ($1, $2, $3, $4, $5) RETURNING ID", chart.UserID, chart.Title, chart.AxisXTitle, chart.AxisYTitle, chart.Data)
	err := row.Scan(&chartID)
	if err != nil {
		return 0, err
	}
	return chartID, nil
}
