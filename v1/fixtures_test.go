package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func CsvMt940Fixture() []CsvMt940 {
	fixture := []CsvMt940{}
	for i := 0; i < 5; i++ {
		fixture = append(fixture, CsvMt940{
			OrderAccount:   Iban(fmt.Sprintf("Iban:%d", i)),
			BookingDate:    time.Now(),
			ValueDate:      time.Now(),
			Narration:      "narration",
			Purpose:        "purpose",
			Target:         "target",
			TargetAccount:  "target account",
			TargetBankCode: "target bank code",
			Amount:         decimal.NewFromInt(101 * int64(-i)).Shift(-2),
			Currency:       "USD",
			Info:           "test info",
		})
	}
	return fixture
}
