package main

import (
	"time"

	"github.com/shopspring/decimal"
)

// Iban represents an IBAN (International Bank Account Number).
type Iban string

// Currency represents a currency.
type Currency string

// CsvMt940 represents a row in an Mt940-formatted csv file.
type CsvMt940 struct {
	OrderAccount   Iban
	BookingDate    time.Time
	ValueDate      time.Time
	Narrative      string
	Purpose        string
	Target         string
	TargetAccount  string
	TargetBankCode string
	Amount         decimal.Decimal
	Currency       Currency
	Info           string
}
