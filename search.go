package searchinppdvn

import (
	"fmt"
	"log"
	"strconv"
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

	maxP := 1
	collector.OnHTML(".pagination a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		pValue := extractPValue(href)
		if pValue > maxP {
			maxP = pValue
		}
	})

	baseURL := "https://ppdvn.gov.vn/web/guest/tra-cuu-luu-chieu?query=%v&id_nxb=-1&p=%d"
	for p := 1; p <= maxP; p++ {
		queryURL := fmt.Sprintf(baseURL, nameQuery, p)
		fmt.Println(queryURL)

		if err := collector.Visit(queryURL); err != nil {
			log.Fatalln("Visit Error: ", err)
		}
	}

	return books
}

func extractPValue(url string) int {
	pIndex := strings.Index(url, "&p=")
	if pIndex == -1 {
		return 0
	}

	pValueStr := url[pIndex+3:]
	pValue, err := strconv.Atoi(pValueStr)
	if err != nil {
		return 0
	}

	return pValue
}
