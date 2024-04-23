package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetURL(url string) string {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}