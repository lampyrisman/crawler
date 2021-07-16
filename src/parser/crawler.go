package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os/exec"
	"strings"
)

const url string = "https://bash.im"

func render(url string) (out []byte) {
	var err error
	out, err = exec.Command("/opt/google/chrome/chrome", "--headless", "--disable-gpu", "--dump-dom", "https://bash.im").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func main() {

	fromRead := bytes.NewReader(render(url))
	doc, err := goquery.NewDocumentFromReader(fromRead)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".quote__frame").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find(".quote__body").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

	doc.Find("a").Each(func(i int, s1 *goquery.Selection) {
		links, ok := s1.Attr("href")
		if ok {
			isRel := strings.Index(links, `/`)
			if isRel == 0 {
				links = "https://bash.im" + links
			}
			fmt.Println(i, isRel, links)
		}
	})

}
