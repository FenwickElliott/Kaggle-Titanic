package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type DataBase struct {
	comparisonRow int
	comparisonExp string
	properties    []*Property
	count         int
}

type Property struct {
	name     string
	total    int
	survived int
	ratio    float64
	col      int
	exp      string
}

func (db *DataBase) addProperty(name string, col int, exp string) {
	db.properties = append(db.properties, &Property{
		name: name,
		col:  col,
		exp:  exp,
	})
}

func NewDataBase(comparisonRow int, comparisonExp string) *DataBase {
	temp := DataBase{comparisonRow: comparisonRow, comparisonExp: comparisonExp}
	temp.addProperty("Global", comparisonRow, comparisonExp)
	temp.count = -1
	return &temp
}

func (p *Property) CalculateRatio() {
	p.ratio = float64(p.survived) / float64(p.total)
}

func (db DataBase) Publish() {
	db.properties[0].total = db.count
	for _, p := range db.properties {
		fmt.Println(p.name, p.total, p.survived, float64(p.survived)/float64(p.total))
	}
}

func main() {
	db := NewDataBase(1, "1")
	db.addProperty("Male", 4, "male")
	db.addProperty("Female", 4, "female")

	// Interestingly publish from here does so with a count of -1
	// defer db.Publish()

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
