
export interface PageRequest {
	search : string
	page_number : number
	page_size : number
	sort_by : string
	sort_direction : string
}

export interface CustomerData {
    id : number
    name : string
    email: string
    phone_number : string
    address: string
    account_number: string
    created_at: string
}

export interface BankData {
    id : number
    account_number : string
    balance : string
    account_type : string
    card_number : string
    cvc : string
    deposits : TermDepositData[]
}

export interface TermDepositTypeData {
    name : string
    min_amount : number
    max_amount : number
}

export interface TermDepositData {
    amount : string
    interest_rate : string
    start_date : string
    maturity_date : string
    status : string
    extension_instructions : string
    term_deposits_types : TermDepositTypeData
}

export interface PocketData {
    name : string
    balance : string
    currency : string
}

export interface CustomerDetailData {
    id : number
    photo : string
    name : string
    email : string
    phone : string
    address : string
    gender : string
    account_purpose : string
    source_of_income : string
    income_per_month : string
    jobs : string
    position : string
    industries : string
    company_name : string
    address_company : string
    banks: BankData
    total_balance : string
    total_deposits : string
    total_pockets : string
    created_at : string
    pockets: PocketData[]
}