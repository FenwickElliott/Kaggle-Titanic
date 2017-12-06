package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Property struct {
	// name     string
	total    int
	survived int
	ratio    float64
}

type DataBase map[string]*Property

func main() {
	csvFile, _ := os.Open("./train.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	db := make(DataBase)
	db["global"] = &Property{}
	db["male"] = &Property{}
	db["female"] = &Property{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			fmt.Println(row)
		} else {
			db["global"].total++
			if row[1] == "1" {
				db["global"].survived++
			}

			if row[4] == "male" {
				db["male"].total++
				if row[1] == "1" {
					db["male"].survived++
				}
			} else {
				db["female"].total++
				if row[1] == "1" {
					db["female"].survived++
				}
			}
		}
	}

	fmt.Println(db["global"])
	fmt.Println(db["male"])
	fmt.Println(db["female"])
}
