package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("skyscanner.net"),
	)
	c.OnHTML("div[class=EcoTicketWrapper_ecoContainer_*]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)

	})
	c.Visit("https://www.skyscanner.net/transport/flights/kut/opo/250823/?adultsv2=1&cabinclass=economy&childrenv2=&ref=home&rtn=0&preferdirects=false&outboundaltsenabled=false&inboundaltsenabled=false")
}

type item struct {
	Name   string `json"name"`
	Price  string `json"price"`
	ImgUrl string `json"imgurl"`
}
