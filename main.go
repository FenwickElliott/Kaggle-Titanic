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

func main() {
	csvFile, _ := os.Open("./train.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	properties := []*Property{}

	properties = append(properties, &Property{
		name: "male",
		col:  4,
		exp:  "male",
	})

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			fmt.Println(row)
		} else {
			for _, p := range properties {
				if row[p.col] == p.exp {
					p.total++
					if row[1] == "1" {
						p.survived++
					}
				}
			}
		}
	}

	fmt.Println(properties[0])
}
