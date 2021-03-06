package main

import (
	"fmt"
	"log"
	"net/http"
	"search-schoolURL/env"
	"strings"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

func customSearchMain() {
	client := &http.Client{Transport: &transport.APIKey{Key: env.ApiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List(env.Query).Cx(env.Cx).Num(3).Do()
	if err != nil {
		log.Fatal(err)
	}

	for i, result := range resp.Items {
		if strings.Contains(result.Link, "ed.jp/") || strings.Contains(result.Link, "ac.jp/") {
			fmt.Printf("#%d: %s\n", i+1, result.Title)
			fmt.Printf("\t%s\n", result.Link)
			break
		}
	}
}

func main() {
	customSearchMain()
}
