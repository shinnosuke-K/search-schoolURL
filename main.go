package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"search-schoolURL/env"
	"strings"
	"time"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

func createCSVfile(fileName string) {
	if _, err := os.Stat("csv/" + fileName); err != nil {
		if _, err := os.Create("csv/" + fileName); err != nil {
			log.Fatal(err)
		}
	}
}

func searchURL(query string) string {

	// for google api 'Custom Search api'
	client := &http.Client{Transport: &transport.APIKey{Key: env.ApiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List(query).Cx(env.Cx).Num(1).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Items {
		if strings.Contains(result.Link, "ed.jp/") || strings.Contains(result.Link, "ac.jp/") {
			return result.Link
		}
	}
	return ""
	// end
}

func writeCSV(fileName, deviValue, schoolName, course, link string) {
	file, err := os.OpenFile("csv/"+fileName, os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{
		deviValue,
		schoolName,
		course,
		link,
	}

	writer := csv.NewWriter(file)
	writer.Write(info)

	writer.Flush()
}

func extraction(reader *csv.Reader, fileName string) {
	beforeName := ""
	link := ""
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		deviValue := line[0]
		schoolName := line[1]
		course := line[2]

		if schoolName != beforeName {
			link = searchURL(line[1])
			beforeName = schoolName
		}

		writeCSV(fileName, deviValue, schoolName, course, link)
		time.Sleep(time.Millisecond * 5)
	}
	return
}

func main() {

	files, err := ioutil.ReadDir(env.DirPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if fileName := file.Name(); fileName != ".DS_Store" {

			createCSVfile(fileName)
			csvFile, err := os.Open(env.DirPath + fileName)
			if err != nil {
				panic(err)
			}

			reader := csv.NewReader(csvFile)

			extraction(reader, fileName)

			csvFile.Close()
		}
	}

}
