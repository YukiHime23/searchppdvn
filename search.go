package searchinppdvn

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Book struct {
	STT            string
	ISBN           string
	Title          string
	Author         string
	Editor         string
	Publisher      string
	Affiliate      string
	PrintingPlace  string
	SubmissionDate string
}

func Search(nameQuery string) []Book {
	baseURL := "https://ppdvn.gov.vn/web/guest/tra-cuu-luu-chieu?query=%v&id_nxb=-1&p=1"

	queryURL := fmt.Sprintf(baseURL, nameQuery)

	collector := colly.NewCollector(
		colly.AllowedDomains("ppdvn.gov.vn"),
	)
	fmt.Println(queryURL)

	var books []Book
	collector.OnHTML("div#list_data_return.table tbody tr", func(e *colly.HTMLElement) {
		book := Book{}
		// Trích xuất dữ liệu từ các thẻ td trong tr
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			switch i {
			case 0:
				book.STT = strings.TrimSpace(td.Text)
			case 1:
				book.ISBN = strings.TrimSpace(td.Text)
			case 2:
				book.Title = strings.TrimSpace(td.Text)
			case 3:
				book.Author = strings.TrimSpace(td.Text)
			case 4:
				book.Editor = strings.TrimSpace(td.Text)
			case 5:
				book.Publisher = strings.TrimSpace(td.Text)
			case 6:
				book.Affiliate = strings.TrimSpace(td.Text)
			case 7:
				book.PrintingPlace = strings.TrimSpace(td.Text)
			case 8:
				book.SubmissionDate = strings.TrimSpace(td.Text)
			}
		})
		books = append(books, book)
	})

	if err := collector.Visit(queryURL); err != nil {
		log.Fatalln("Visit Error: ", err)
	}

	return books
}
