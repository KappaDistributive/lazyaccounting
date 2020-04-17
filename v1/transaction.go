package main

import (
	"github.com/shopspring/decimal"
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
	value    decimal.Decimal
	currency Currency
}

func (amount Amount) String() string {
	return amount.value.String() + " " + amount.currency.String()
}

type Comment string

func (comment Comment) String() string {
	return string(comment)
}

type Cost Amount

func (cost Cost) String() string {
	return Amount(cost).String()
}

type Currency string

func (currency Currency) String() string {
	return string(currency)
}

type Flag int

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

const (
	Complete   Flag = 0
	Incomplete Flag = 1 // TODO are there more relevant flags?
)

type Link string

func (link Link) String() string {
	return string(link)
}

type Metadata struct {
	key   string
	value string
}

func (metadata Metadata) String() string {
	return metadata.key + ":" + metadata.value
}

type Narration string

func (narration Narration) String() string {
	return string(narration)
}

type Payee string

func (payee Payee) String() string {
	return string(payee)
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
	// TODO account for missing entries
	var represenation string

	represenation += posting.flag.String()
	represenation += " " + posting.account.String()
	represenation += " " + posting.amount.String()
	if string(posting.cost.currency) != "" {
		represenation += " " + posting.cost.String()
	}
	if string(posting.price.currency) != "" {
		represenation += " @ " + posting.price.String()
	}
	if len(posting.comment.String()) > 0 {

		represenation += " ; " + posting.comment.String()
	}

	return represenation
}

type Price Amount

func (price Price) String() string {
	return Amount(price).String()
}

type Tag string

func (tag Tag) String() string {
	return string(tag)
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
	// TODO account for missing entries
	var represenation string

	represenation = strings.Split(transactionHeader.date.Format(time.RFC3339), "T")[0]
	represenation += " " + transactionHeader.flag.String()
	represenation += " \"" + transactionHeader.payee.String() + "\""
	represenation += " \"" + transactionHeader.narration.String() + "\""
	represenation += " #" + transactionHeader.tag.String()
	represenation += " ^" + transactionHeader.link.String()
	represenation += " ; " + transactionHeader.comment.String()
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
