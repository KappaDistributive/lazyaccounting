package main

import (
	"time"

	"github.com/shopspring/decimal"
)

// Iban represents an IBAN (International Bank Account Number).
type Iban string

// CsvMt940 represents a row in an Mt940-formatted csv file.
type CsvMt940 struct {
	OrderAccount   Iban
	BookingDate    time.Time
	ValueDate      time.Time
	Narration      string
	Purpose        string
	Target         string
	TargetAccount  string
	TargetBankCode string
	Amount         decimal.Decimal
	Currency       Currency
	Info           string
}

func (row CsvMt940) Postings(
	internalAccount Account,
	externalAccount Account) []Posting {
	internalPosting := Posting{
		flag:    -1,
		account: internalAccount,
		amount: Amount{
			value:    row.Amount,
			currency: row.Currency,
		},
		cost: Cost(Amount{
			value:    decimal.Decimal{},
			currency: row.Currency,
		}),
		price: Price(Amount{
			value:    decimal.Decimal{},
			currency: row.Currency,
		}),
		comment:  Comment(""),
		metadata: []Metadata{},
	}
	externalPosting := Posting{
		flag:    -1,
		account: externalAccount,
		amount: Amount{
			value:    row.Amount.Neg(),
			currency: row.Currency,
		},
		cost: Cost(Amount{
			value:    decimal.Decimal{},
			currency: row.Currency,
		}),
		price: Price(Amount{
			value:    decimal.Decimal{},
			currency: row.Currency,
		}),
		comment:  Comment(""),
		metadata: []Metadata{},
	}

	return []Posting{internalPosting, externalPosting}
}

func (row CsvMt940) TransactionHeader() TransactionHeader {
	return TransactionHeader{
		date:      row.ValueDate,
		flag:      1,
		payee:     Payee(""),
		narration: Narration(row.Narration),
		tag:       Tag(""),
		link:      Link(""),
		comment:   Comment(""),
		metadata:  []Metadata{},
	}
}

func (row CsvMt940) Transaction(
	internalAccount Account,
	externalAccount Account) Transaction {
	return Transaction{
		header:   row.TransactionHeader(),
		postings: row.Postings(internalAccount, externalAccount),
	}
}
