package models

// Chart models a chart in database
type Chart struct {
	UserAsset
	Title      string `json:"title"`
	AxisXTitle string `json:"axis_x_title"`
	AxisYTitle string `json:"axis_y_title"`
	Data       string `json:"data"`
}
