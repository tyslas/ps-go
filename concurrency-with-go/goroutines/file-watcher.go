package main

import (
	"fmt"
	"os" //for working with the file system
	"time"
	"io/ioutil"
	"strings" //read contents of a file into a reader
	"encoding/csv" //parse csv files
	"strconv" //convert from non-string fields and back to correct datatype
	"runtime"
)

const watchedPath = "/Users/tito.yslas/workspace/go/source"

func main() {

	runtime.GOMAXPROCS(4)

	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("\nreceived invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data))
		}
		d.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

type Invoice struct {
	Number string
	Amount float64
	PurchaseOrderNumber int
	InvoiceDate time.Time
}