package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name    string `json:"name"`
	Article string `json:"article"`
	Price   string `json:"price"`
	ImgUrl  string `json:"imgurl"`
}

func main() {

	c := colly.NewCollector(
	// colly.AllowedDomains("skyscanner.ie"),
	)

	var items []item

	// c.OnHTML("div[class^='FlightsResults_dayViewItems'] div[class*='Ticket']", func(h *colly.HTMLElement) {
	c.OnHTML("div[class^='FlightsResults_dayViewItems__NzJiY']", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
		fmt.Println("1")
		item := item{

			Name: h.ChildText("a.catalog_list_name"),
			// Article: h.ChildText("span.catalog_list_article"),
			Price: h.ChildText("span.catalog_list_actual_price"),
			// ImgUrl:  h.ChildAttr(".catalog_list_img_wrap_item img", "src"),
		}
		// fmt.Printf("Item is : %v\n", items)
		items = append(items, item)
	})

	c.OnHTML("ul.pagin_list a.next", func(h *colly.HTMLElement) {

		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		fmt.Printf(next_page)
		c.Visit(next_page)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://www.skyscanner.ie/transport/flights/bus/nap/250822/?adultsv2=1&cabinclass=economy&childrenv2=&inboundaltsenabled=false&outboundaltsenabled=false&preferdirects=false&rtn=0")
	fmt.Println(items)
	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}

