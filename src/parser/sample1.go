package main

import (
    "github.com/geziyor/geziyor"
    "github.com/geziyor/geziyor/export"
    "github.com/geziyor/geziyor/client"
    "github.com/PuerkitoBio/goquery"
    "fmt"
)


func main() {
    geziyor.NewGeziyor(&geziyor.Options{
        StartURLs: []string{"https://bash.im"},
        ParseFunc: quotesParse,
        Exporters: []export.Exporter{&export.JSON{}},
    }).Start()
}

func quotesParse(g *geziyor.Geziyor, r *client.Response) {
    r.HTMLDoc.Find("article.quote").Each(func(i int, s *goquery.Selection) {
        g.Exports <- map[string]interface{}{

//            "text":   s.Find("span.text").Text(),
//            "author": s.Find("small.author").Text(),
	      "Quote": s.Find("div.quote__body").Text(),
        }
    })
    fmt.Println(string(r.Body))
    if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
	res, err := r.JoinURL(href)
	fmt.Println("res",res.String(), err)
        g.Get(res.String(), quotesParse)
    }
}

