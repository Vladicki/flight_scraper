package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	getFlights()
}

func getFlights() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.skyscanner.ie/transport/flights/rmo/bus/250822/?adultsv2=1&cabinclass=economy&childrenv2=&ref=home&rtn=0&preferdirects=false&outboundaltsenabled=false&inboundaltsenabled=false", g.Opt.ParseFunc)
		},
		// StartURLs: []string{"https://www.skyscanner.ie/transport/flights/rmo/bus/250822/?adultsv2=1&cabinclass=economy&childrenv2=&ref=home&rtn=0&preferdirects=false&outboundaltsenabled=false&inboundaltsenabled=false"},
		ParseFunc:       skyParse,
		Exporters:       []export.Exporter{&export.CSV{}},
		BrowserEndpoint: "ws://localhost:9222",
	}).Start()
}

func skyParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div[class ^='FlightsTicket_container'").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"text":   s.Find("div[class ^='Ticket'").Text(),
			"author": s.Find("small.author").Text(),
		}
	})
	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
		g.Get(r.JoinURL(href), skyParse)
	}

	//	func quotesParse(g *geziyor.Geziyor, r *client.Response) {
	//		r.HTMLDoc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
	//			g.Exports <- map[string]interface{}{
	//				"text":   s.Find("span.text").Text(),
	//				"author": s.Find("small.author").Text(),
	//			}
	//		})
	//		if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
	//			g.Get(r.JoinURL(href), quotesParse)
	//		}
}
