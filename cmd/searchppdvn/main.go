package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Book struct {
	STT           string
	ISBN          string
	TenSach       string
	TacGia        string
	BienTapVien   string
	NhaXuatBan    string
	DoiTacLienKet string
	NoiIn         string
	NgayNopLC     string
}

func main() {
	baseURL := "https://ppdvn.gov.vn/web/guest/tra-cuu-luu-chieu?query=%v&id_nxb=-1&p=1"

	var nameQuery string

	nameQ := flag.String("name", "", "The name of the book you want to search for on PPDVN.")
	flag.Parse()
	if nameQ == nil {
		nameQuery = ""
	} else {
		nameQuery = strings.Replace(*nameQ, " ", "+", -1)

	}

	queryURL := fmt.Sprintf(baseURL, nameQuery)

	collector := colly.NewCollector(
		colly.AllowedDomains("ppdvn.gov.vn"),
	)

	var books []Book
	collector.OnHTML("div#list_data_return.table tbody tr", func(e *colly.HTMLElement) {
		book := Book{}
		// Trích xuất dữ liệu từ các thẻ td trong tr
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			switch i {
			case 0:
				book.STT = td.Text
			case 1:
				book.ISBN = td.Text
			case 2:
				book.TenSach = td.Text
			case 3:
				book.TacGia = td.Text
			case 4:
				book.BienTapVien = td.Text
			case 5:
				book.NhaXuatBan = td.Text
			case 6:
				book.DoiTacLienKet = td.Text
			case 7:
				book.NoiIn = td.Text
			case 8:
				book.NgayNopLC = td.Text
			}
		})
		books = append(books, book)
	})

	if err := collector.Visit(queryURL); err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}
