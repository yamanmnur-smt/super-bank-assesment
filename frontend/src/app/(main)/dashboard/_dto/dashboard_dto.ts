export interface DashboardTotalCard {
    total_customers : string
    total_deposits : string
    total_balance : string
}

export interface ChartData {
    label : string
    value : string
}

export interface DashboardData {
    total_card : DashboardTotalCard
    pie_data : ChartData[]
    bar_data : ChartData[]
}
