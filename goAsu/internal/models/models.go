package models

type ObjectType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Object struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type Well struct {
	Well int `json:"well"`
	NGDU int `json:"ngdu"`
	CDNG int `json:"cdng"`
	Kust int `json:"kust"`
	Mest int `json:"mest"`
}

type WellDayHistory struct {
	Well          int     `json:"well"`
	DateFact      string  `json:"date_fact"`
	Debit         float64 `json:"debit"`
	EEConsume     float64 `json:"ee_consume"`
	Expenses      float64 `json:"expenses"`
	PumpOperating float64 `json:"pump_operating"`
}

type WellDayPlan struct {
	Well          int     `json:"well"`
	DatePlan      string  `json:"date_plan"`
	Debit         float64 `json:"debit"`
	EEConsume     float64 `json:"ee_consume"`
	Expenses      float64 `json:"expenses"`
	PumpOperating float64 `json:"pump_operating"`
}

const (
	BASE_IP  = "109.120.183.88"
	PORT     = 5432
	USERNAME = "hetsu"
	PASSWORD = "Admin1234567890!"
	BASENAME = "PostgreSQL-vitalick113"
)
