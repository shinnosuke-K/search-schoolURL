package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"search-schoolURL/env"
)

func main() {

	files, err := ioutil.ReadDir(env.DirPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if fileName := file.Name(); fileName != ".DS_Store" {

			csvFile, err := os.Open(env.DirPath + fileName)
			if err != nil {
				panic(err)
			}

			reader := csv.NewReader(csvFile)
			var line []string

			for {
				line, err = reader.Read()
				if err != nil {
					break
				}
				fmt.Println(line[1])
			}
			csvFile.Close()
		}
	}

}
