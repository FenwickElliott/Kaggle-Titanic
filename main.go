package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Property struct {
	name     string
	total    int
	survived int
	ratio    float64
	col      int
	exp      string
}

func (p *Property) CalculateRatio() {
	p.ratio = float64(p.survived) / float64(p.total)
}

type DataBase struct {
	properties []*Property
}

func (db *DataBase) addProperty(name string, col int, exp string) {
	db.properties = append(db.properties, &Property{
		name: name,
		col:  col,
		exp:  exp,
	})
}

func main() {
	db := DataBase{}
	db.addProperty("male", 4, "male")
	db.addProperty("female", 4, "female")

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
			for _, p := range db.properties {
				if row[p.col] == p.exp {
					p.total++
					// if row[1] == "1" {
					// 	p.survived++
					// }
				}
			}
		}
	}

	fmt.Println(db.properties[0])
	fmt.Println(db.properties[1])
}
