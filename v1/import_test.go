package main

import (
	"testing"
)

func TestPostings(t *testing.T) {
	csv := CsvMt940Fixture()

	postings := make([][]Posting, len(csv))

	for i, row := range csv {
		internalAccount := Account{
			components: []string{"internal", "account", string(i)},
		}
		externalAccount := Account{
			components: []string{"external", "account", string(i)},
		}

		postings[i] = row.Postings(internalAccount, externalAccount)
	}

	for i := 0; i < len(postings); i++ {
		if !postings[i][0].amount.value.Equal(postings[i][1].amount.value.Neg()) {
			t.Error("Transaction doesn't balance.")
		}
		if postings[i][0].amount.currency != postings[i][1].amount.currency {
			t.Error("Postings have different currencies.")
		}
	}
}

func TestTransactionHeader(t *testing.T) {
	csv := CsvMt940Fixture()

	headers := []TransactionHeader{}

	for _, row := range csv {
		headers = append(headers, row.TransactionHeader())
	}

	for i := 0; i < len(headers); i++ {
		// check date
		if headers[i].date != csv[i].ValueDate {
			t.Errorf("got: %s want: %s", headers[i].date, csv[i].ValueDate)
		}

		// check narration
		if string(headers[i].narration) != csv[i].Narration {
			t.Errorf("got: %s want: %s", headers[i].date, csv[i].ValueDate)

		}
	}
}
