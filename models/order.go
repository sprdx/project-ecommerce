package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Payment Payment
}

type Payment struct {
	E_Wallet string  `json:"wallet"`
	Amount   float64 `json:"amount"`
}
