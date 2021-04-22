package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	give_current_price_of()
}

func give_current_price_of() {

	c := colly.NewCollector(
		colly.AllowedDomains("google.com", "www.google.com"),
	)

	found := false

	c.OnHTML("div", func(e *colly.HTMLElement) {

		if found {
			return
		}

		if e.Attr("class") == "BNeawe iBp4i AP7Wnd" {
			fmt.Printf("$ = %s\n", e.Text)
			found = true
		}
	})

	c.Visit("https://www.google.com/search?q=dolar+preco")

}
