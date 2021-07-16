package main

import (
<<<<<<< HEAD
    "github.com/geziyor/geziyor"
    "github.com/geziyor/geziyor/client"
    "github.com/geziyor/geziyor/export"
    "github.com/PuerkitoBio/goquery"
=======
>>>>>>> ed311ac8bff6f4885370f24703ceb68cf921c8a1
    "fmt"
    "os/exec"
    "log"
    "github.com/PuerkitoBio/goquery"
    "bytes"
//    "strings"
    "net/http"
    "net/url"
    "io"

)

const startURL string = "http://hiddenchan.i2p/"

func render (url string) (out []byte){
    var err error
    out, err = exec.Command("/opt/google/chrome/chrome" , "--proxy-server=localhost:4444" , "--headless" , "--disable-gpu" , "--dump-dom" , startURL).Output()
    if err != nil {
	log.Fatal(err)
    }
    return out
}

<<<<<<< HEAD
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
=======

func getJumps (doc *goquery.Document)(urls []string){
    	urlBase := doc.Find("div#jumplinks").First()
	urlBase.Find("a").Each(func(i int, hrefQ *goquery.Selection) {
	links, ok := hrefQ.Attr("href")
	if ok {
    		fmt.Println(i, links)
		urls = append(urls, links)
	}
    })
	return urls
}

func getRedir (jumpURL []string) (redirURL string){
    proxyUrl, err := url.Parse("http://127.0.0.1:4444")
    fmt.Println(err)
    i2p := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	for _,jURL := range jumpURL {
	resp, err := i2p.Get(jURL)
	if resp.StatusCode != 200 {
	    fmt.Println( jURL, resp.StatusCode)
	    continue
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode,string(body),err)
    }
    return redirURL
}


func main () {
    fmt.Println(string(render(startURL)))
    fromRead := bytes.NewReader(render(startURL))
  doc, err := goquery.NewDocumentFromReader(fromRead)
  if err != nil {
    log.Fatal(err)
  }
	title := doc.Find("title").Text()
	fmt.Println(title)
	if title == "Website Unknown" {
	    jumpsURL := getJumps(doc)
	    redirURL := getRedir(jumpsURL)
	    fmt.Println(redirURL)
	}

/*
	doc.Find("a").Each(func(i int, hrefQ *goquery.Selection) {
	links, ok := hrefQ.Attr("href")
	if ok {
	    isRel := strings.Index(links, `/`)
	    if isRel == 0 {
		links = url + links
    		fmt.Println(i, isRel, links)
	    }


	    isRel = strings.Index(links, `./`)
	    if isRel == 0 {
		links = url + "/" +links
    		fmt.Println(i, isRel, links)
	    }


	}
    })
*/

>>>>>>> ed311ac8bff6f4885370f24703ceb68cf921c8a1
}

