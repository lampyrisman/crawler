package main

import (
    "github.com/geziyor/geziyor"
    "github.com/geziyor/geziyor/client"
    "github.com/geziyor/geziyor/export"
    "github.com/PuerkitoBio/goquery"
    "fmt"
)

/*
func main() {
geziyor.NewGeziyor(&geziyor.Options{
    StartRequestsFunc: func(g *geziyor.Geziyor) {
        g.GetRendered("https://www.spletnik.ru", g.Opt.ParseFunc)
    },
    ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
        fmt.Println(string(r.Body))
    },
    //BrowserEndpoint: "ws://localhost:3000",
}).Start()}

*/

func main () {
    geziyor.NewGeziyor(&geziyor.Options{
//        StartURLs: []string{"http://quotes.toscrape.com/"},
    StartRequestsFunc: func(g *geziyor.Geziyor) {
        g.GetRendered("http://10.10.2.2/test2/index.html", g.Opt.ParseFunc)
    },
        ParseFunc: quotesParse,
        Exporters: []export.Exporter{&export.CSV{}},
    }).Start()
}

func getFullPage(g *geziyor.Geziyor, r *client.Response) {
    fmt.Println(string(r.Body))
}

func quotesParse(g *geziyor.Geziyor, r *client.Response) {
    r.HTMLDoc.Find("tbody.t431__tbody").Each(func(i int, s *goquery.Selection) {
        g.Exports <- map[string]interface{}{

	"tdata":   s.Find("td").Text(),

//	"Head": s.Find("h2").Text(),
//        "Quote": s.Find("div.newsarea").Text(),
        }
    })

//    if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
//    res, err := r.JoinURL(href)
//    fmt.Println("res",res.String(), err)
//        g.Get(res.String(), quotesParse)
//    }
}

