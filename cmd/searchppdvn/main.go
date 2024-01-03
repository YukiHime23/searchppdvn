package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	searchinppdvn "github.com/YukiHime23/search-in-ppdvn"
)

func main() {
	var nameQuery string

	nameQ := flag.String("name", "", "The name of the book you want to search for on PPDVN.")
	flag.Parse()
	if nameQ == nil {
		nameQuery = ""
	} else {
		nameQuery = strings.Replace(*nameQ, " ", "+", -1)

	}

	resultJson := searchinppdvn.Search(nameQuery)
	jsonData, err := json.Marshal(resultJson)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Println(jsonData)
}
