package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

func CustomerSeeder() {
	customers := []models.Customer{
		{Name: "Yaman", Email: "yaman@mail.com", PhoneNumber: "1234567890", Address: "123 Street, City, Country", Username: "yaman", Gender: "male", AccountPurpose: "savings", SourceOfIncome: "salary", IncomePerMonth: "5000", Jobs: "developer", Position: "senior developer", Industries: "IT", CompanyName: "Tech Corp", AddressCompany: "456 Tech Street, City, Country"},
		{Name: "Alice", Email: "alice@mail.com", PhoneNumber: "1234567891", Address: "124 Street, City, Country", Username: "alice", Gender: "female", AccountPurpose: "investment", SourceOfIncome: "business", IncomePerMonth: "7000", Jobs: "entrepreneur", Position: "CEO", Industries: "Retail", CompanyName: "Alice Ventures", AddressCompany: "789 Business Ave, City, Country"},
		{Name: "Bob", Email: "bob@mail.com", PhoneNumber: "1234567892", Address: "125 Street, City, Country", Username: "bob", Gender: "male", AccountPurpose: "current", SourceOfIncome: "freelancing", IncomePerMonth: "4000", Jobs: "designer", Position: "graphic designer", Industries: "Creative", CompanyName: "Design Studio", AddressCompany: "321 Creative Lane, City, Country"},
		{Name: "Charlie", Email: "charlie@mail.com", PhoneNumber: "1234567893", Address: "126 Street, City, Country", Username: "charlie", Gender: "male", AccountPurpose: "savings", SourceOfIncome: "salary", IncomePerMonth: "6000", Jobs: "engineer", Position: "software engineer", Industries: "Technology", CompanyName: "Innovatech", AddressCompany: "654 Innovation Blvd, City, Country"},
		{Name: "David", Email: "david@mail.com", PhoneNumber: "1234567894", Address: "127 Street, City, Country", Username: "david", Gender: "male", AccountPurpose: "investment", SourceOfIncome: "real estate", IncomePerMonth: "8000", Jobs: "investor", Position: "real estate investor", Industries: "Real Estate", CompanyName: "David Properties", AddressCompany: "987 Realty Road, City, Country"},
		{Name: "Eve", Email: "eve@mail.com", PhoneNumber: "1234567895", Address: "128 Street, City, Country", Username: "eve", Gender: "female", AccountPurpose: "current", SourceOfIncome: "consulting", IncomePerMonth: "5500", Jobs: "consultant", Position: "business consultant", Industries: "Consulting", CompanyName: "Eve Consulting", AddressCompany: "159 Consultant St, City, Country"},
		{Name: "Frank", Email: "frank@mail.com", PhoneNumber: "1234567896", Address: "129 Street, City, Country", Username: "frank", Gender: "male", AccountPurpose: "savings", SourceOfIncome: "salary", IncomePerMonth: "4500", Jobs: "teacher", Position: "high school teacher", Industries: "Education", CompanyName: "City High School", AddressCompany: "753 Education Lane, City, Country"},
		{Name: "Grace", Email: "grace@mail.com", PhoneNumber: "1234567897", Address: "130 Street, City, Country", Username: "grace", Gender: "female", AccountPurpose: "investment", SourceOfIncome: "trading", IncomePerMonth: "9000", Jobs: "trader", Position: "stock trader", Industries: "Finance", CompanyName: "Grace Trading Co.", AddressCompany: "246 Finance Blvd, City, Country"},
		{Name: "Hank", Email: "hank@mail.com", PhoneNumber: "1234567898", Address: "131 Street, City, Country", Username: "hank", Gender: "male", AccountPurpose: "current", SourceOfIncome: "construction", IncomePerMonth: "5000", Jobs: "contractor", Position: "construction manager", Industries: "Construction", CompanyName: "Hank Builders", AddressCompany: "369 Build Ave, City, Country"},
		{Name: "Ivy", Email: "ivy@mail.com", PhoneNumber: "1234567899", Address: "132 Street, City, Country", Username: "ivy", Gender: "female", AccountPurpose: "savings", SourceOfIncome: "freelancing", IncomePerMonth: "4800", Jobs: "writer", Position: "content writer", Industries: "Media", CompanyName: "Ivy Media", AddressCompany: "951 Media Lane, City, Country"},
	}

	gormdb := db.GetDbHandler().DB
	for _, customer := range customers {
		gormdb.FirstOrCreate(&customer, models.Customer{Email: customer.Email})
	}
}
