package converter

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// The line which contains the columns
const columnsLine = 11

// CSVBody represents a CSV file content
type CSVBody struct {
	columns []string
	rows    [][]string
}

// ToYNAB formats according to https://docs.youneedabudget.com/article/921-importing-a-csv-file

// Option 2
// This format has one small variation—the amount is one field instead of two. Outflows are identified by a negative sign in the Amount field.

// Date,Payee,Memo,Amount
// 06/22/21,Payee 1,Memo,-100.00
// 06/22/21,Payee 2,Memo,500.00

func ToYNAB(bt CSVBody) CSVBody {
	var ynab CSVBody

	ynab.columns = []string{"Date", "Payee", "Memo", "Amount"}

	for _, row := range bt.rows {
		var amount string
		var payee string

		description := row[2]
		payee = description
		if len(strings.Split(description, ";")) > 1 {
			payee = strings.Split(description, ";")[1]
		}

		debit, credit := row[4], row[5]

		if debit != "" {
			amount = debit
		} else {
			amount = credit
		}

		ynab.rows = append(ynab.rows, []string{row[0], payee, description, amount})
	}

	return ynab
}

// Convert will convert the file located in csvPath from Banca Transilvania's format to YNAB format
func Convert(csvPath string) error {
	f, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	var result CSVBody
	result.columns = make([]string, 0)
	result.rows = make([][]string, 0)

	currentLine := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		currentLine++

		if currentLine < columnsLine {
			continue
		}

		if currentLine == columnsLine {
			result.columns = record
		} else {
			result.rows = append(result.rows, record)
		}
	}

	return writeToStdout(ToYNAB(result))
}

// writeToStdout will write to stdout the CSV contents of body
func writeToStdout(body CSVBody) error {
	records := [][]string{
		body.columns,
	}
	records = append(records, body.rows...)

	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
