package main

import (
	decimal "github.com/shopspring/decimal"
	"strings"
	"time"
)

const (
	OFFSET_METADATA int = 2
	OFFSET_POSTING  int = 2
)

type Account struct {
	components []string
}

func (account Account) String() string {
	return strings.Join(account.components, ":")
}

type Amount struct {
	value    decimal.NullDecimal
	currency Currency
}

func (amount Amount) String() string {
	return amount.value.Decimal.String() + " " + amount.currency.String()
}

type Comment string

func (comment Comment) IsActive() bool {
	return len(string(comment)) > 0
}

func (comment Comment) String() string {
	return "; " + string(comment)
}

type Cost Amount

func (cost Cost) IsActive() bool {
	return Amount(cost).value.Valid || len(string(Amount(cost).currency)) > 0
}

func (cost Cost) String() string {
	return Amount(cost).String()
}

type Currency string

func (currency Currency) String() string {
	return string(currency)
}

type Flag int

const (
	Inactive   Flag = -1
	Complete   Flag = 0
	Incomplete Flag = 1
)

func (flag Flag) IsActive() bool {
	return flag != Inactive
}

func (flag Flag) String() string {
	var represenation string
	switch flag {
	case -1:
		represenation = ""
	case 0:
		represenation = "!"
	case 1:
		represenation = "*"

	default:
		represenation = "?"
	}
	return represenation

}

type Link string

func (link Link) IsActive() bool {
	return len(string(link)) > 0
}

func (link Link) String() string {
	return "^" + string(link)
}

type Metadata struct {
	key   string
	value string
}

func (metadata Metadata) String() string {
	return metadata.key + ":" + metadata.value
}

type Narration string

func (narration Narration) IsActive() bool {
	return len(string(narration)) > 0
}

func (narration Narration) String() string {
	return "\"" + string(narration) + "\""
}

type Payee string

func (payee Payee) IsActive() bool {
	return len(string(payee)) > 0
}

func (payee Payee) String() string {
	return "\"" + string(payee) + "\""
}

type Posting struct {
	flag     Flag
	account  Account
	amount   Amount
	cost     Cost
	price    Price
	comment  Comment
	metadata []Metadata
}

func (posting Posting) String() string {
	var represenation string

	if posting.flag.IsActive() {
		represenation += posting.flag.String() + " "
	}
	represenation += posting.account.String()
	if posting.cost.IsActive() {
		represenation += " " + posting.cost.String()
	}
	if posting.price.IsActive() {
		represenation += " " + posting.price.String()
	}
	if len(posting.comment.String()) > 0 {
		represenation += " " + posting.comment.String()
	}

	return represenation
}

type Price Amount

func (price Price) IsActive() bool {
	return Amount(price).value.Valid || len(string(Amount(price).currency)) > 0
}

func (price Price) String() string {
	return "@ " + Amount(price).String()
}

type Tag string

func (tag Tag) IsActive() bool {
	return len(string(tag)) > 0
}

func (tag Tag) String() string {
	return "#" + string(tag)
}

type TransactionHeader struct {
	date      time.Time
	flag      Flag
	payee     Payee
	narration Narration
	tag       Tag
	link      Link
	comment   Comment
	metadata  []Metadata
}

func (transactionHeader TransactionHeader) String() string {
	var represenation string

	represenation = strings.Split(transactionHeader.date.Format(time.RFC3339), "T")[0]
	represenation += " " + transactionHeader.flag.String()
	if transactionHeader.payee.IsActive() {
		represenation += " " + transactionHeader.payee.String()
	}
	if transactionHeader.narration.IsActive() {
		represenation += " " + transactionHeader.narration.String()
	}
	if transactionHeader.tag.IsActive() {
		represenation += " " + transactionHeader.tag.String()
	}
	if transactionHeader.link.IsActive() {
		represenation += " " + transactionHeader.link.String()
	}
	if transactionHeader.comment.IsActive() {
		represenation += " " + transactionHeader.comment.String()
	}
	for _, entry := range transactionHeader.metadata {
		represenation += "\n" + strings.Repeat(" ", OFFSET_METADATA) + entry.String()
	}

	return represenation
}

type Transaction struct {
	header   TransactionHeader
	postings []Posting
}

func (transaction Transaction) String() string {
	represenation := transaction.header.String()

	for _, posting := range transaction.postings {
		represenation += "\n" + strings.Repeat(" ", OFFSET_POSTING) + posting.String()
	}

	return represenation
}
