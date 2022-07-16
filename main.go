package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// https://www.tokopedia.com/search?st=product&q=Zjiang%20ZJ-5890T&srp_component_id=02.01.00.00

func main() {

	c := colly.NewCollector()

	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("OnError")
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("OnResponse")
		fmt.Println("Request URL:", r.Request.URL, "responded with status code:", r.StatusCode)
	})

	c.Visit("https://www.tokopedia.com/search?st=product&q=Zjiang%20ZJ-5890T&srp_component_id=02.01.00.00")
}
