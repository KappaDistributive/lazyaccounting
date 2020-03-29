package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func main() {

	account := Account{
		components: []string{"Assets", "MyBank", "Checking"},
	}
	fmt.Println("Sample account:", account)

	temp, _ := decimal.NewFromString("1.5")
	value := decimal.NullDecimal{
		Decimal: temp,
		Valid:   true,
	}

	amount := Amount{
		value:    value,
		currency: "EUR",
	}
	fmt.Println("Sample amount:", amount)

	comment := Comment("test comment")
	fmt.Println("Sample comment:", comment)

	cost := Cost{
		value:    value,
		currency: "EUR",
	}
	fmt.Println("Sample cost:", cost)

	currency := Currency("EUR")
	fmt.Println("Sample currencty:", currency)

	incomplete := Flag(0)
	fmt.Println("Incomplete flag:", incomplete)

	complete := Flag(1)
	fmt.Println("Complete flag:", complete)

	link := Link("test link")
	fmt.Println("Sample link:", link)

	metadata := Metadata{
		key:   "test key",
		value: "test value",
	}
	fmt.Println("Sample metadata:", metadata)

	narration := Narration("test narration")
	fmt.Println("Sample narration:", narration)

	payee := Payee("test payee")
	fmt.Println("Sample payee:", payee)

	price := Price{
		amount: Amount{
			value:    value,
			currency: "EUR",
		},
		kind: 0,
	}
	fmt.Println("Sample price:", price)

	tag := Tag("test tag")
	fmt.Println("Sample tag:", tag)

	posting := Posting{
		flag:     complete,
		account:  account,
		amount:   amount,
		cost:     cost,
		price:    price,
		comment:  comment,
		metadata: []Metadata{metadata},
	}
	fmt.Println("Sample posting:")
	fmt.Println(posting)

	header := TransactionHeader{
		date:      time.Now(),
		flag:      complete,
		payee:     payee,
		narration: narration,
		tag:       tag,
		link:      link,
		comment:   comment,
		metadata:  []Metadata{metadata},
	}
	fmt.Println("Sample transactionheader:")
	fmt.Println(header)

	transaction := Transaction{
		header:   header,
		postings: []Posting{posting},
	}
	fmt.Println("Sample transaction:")
	fmt.Println(transaction)

	posting.flag = -1
	transaction = Transaction{
		header:   header,
		postings: []Posting{posting},
	}
	fmt.Println("Sample transaction with inactive flag:")
	fmt.Println(transaction)

}
