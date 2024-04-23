package main

import (
	"fmt"
	"io/ioutil"
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

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}

func main() {
	res := GetURL("http://localhost:8080")
	fmt.Println(res)
}
