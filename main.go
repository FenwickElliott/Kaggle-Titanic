package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// DataBase holds all analysis data
type DataBase struct {
	comparisonRow int
	comparisonExp string
	properties    []*Property
	count         int
}

// Property describes a line of analysis and records the raw data as the sample is parsed
type Property struct {
	name     string
	total    int
	survived int
	ratio    float64
	col      int
	exp      string
}

// AddProperty adds Properties to DataBases
func (db *DataBase) AddProperty(name string, col int, exp string) {
	db.properties = append(db.properties, &Property{
		name: name,
		col:  col,
		exp:  exp,
	})
}

// NewDataBase constructs DataBase types
func NewDataBase(comparisonRow int, comparisonExp string) *DataBase {
	temp := DataBase{comparisonRow: comparisonRow, comparisonExp: comparisonExp}
	temp.AddProperty("Global", comparisonRow, comparisonExp)
	return &temp
}

// Publish makes all nessasery calculations and prints the results
func (db DataBase) Publish() {
	db.properties[0].total = db.count - 1
	for _, p := range db.properties {
		fmt.Println(p.name, p.total, p.survived, float64(p.survived)/float64(p.total))
	}
}

func main() {
	db := NewDataBase(1, "1")
	db.AddProperty("Male", 4, "male")
	db.AddProperty("Female", 4, "female")
	db.AddProperty("1st Class", 2, "1")
	db.AddProperty("2nd Class", 2, "2")
	db.AddProperty("3rd Class", 2, "3")
	db.AddProperty("Cherbourg", 11, "C")
	db.AddProperty("Queenstown", 11, "Q")
	db.AddProperty("Southhampton", 11, "S")

	csvFile, _ := os.Open("./train.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			fmt.Println(row)
		} else {
			db.count++
			for _, p := range db.properties {
				if row[p.col] == p.exp {
					p.total++
					if row[db.comparisonRow] == db.comparisonExp {
						p.survived++
					}
				}
			}
		}
	}
	db.Publish()
}
