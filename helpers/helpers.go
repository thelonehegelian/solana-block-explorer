package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToCSV(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
	}
}

func ToString(d interface{}) string {
	return fmt.Sprintf("%v", d)
}
