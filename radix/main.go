package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jiop/various/radix/radix"
)

type resp struct {
	content   interface{}
	updatedAt time.Time
}

func transformStr(s string) string {
	return strings.ToLower(strings.Replace(s, " ", "_", -1))
}

func loadRadix(csvFile string, radixStore *radix.Tree) error {
	f, err := os.Open(csvFile)
	defer f.Close()

	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	_, err = r.Read()
	if err != nil {
		return err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		radixStore.Insert(transformStr(record[1]), record[2])
	}

	return nil
}

func main() {
	radixStore := radix.New()
	csvFile := "csv/laposte_hexasmal.csv"

	if err := loadRadix(csvFile, radixStore); err != nil {
		log.Fatal(err)
	}

	city := "PARIS 02"
	res, ok := radixStore.Get(transformStr(city))
	if !ok {
		log.Print("not found.")
	}
	log.Printf("%v", res)
}
