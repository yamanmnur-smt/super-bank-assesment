package data

type DashboardData struct {
	TotalCards DashboardTotalCard `json:"total_card"`
	PieData    []ChartData        `json:"pie_data"`
	BarData    []ChartData        `json:"bar_data"`
}

type DashboardTotalCard struct {
	TotalCustomers string `json:"total_customers"`
	TotalDeposits  string `json:"total_deposits"`
	TotalBalance   string `json:"total_balance"`
}

type ChartData struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
