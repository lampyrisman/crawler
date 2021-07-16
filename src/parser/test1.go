package main

import (
    "fmt"
    "regexp"
    )


func main() {
    testo := "http://ololo.com/dpc/res.org"
    r_url := regexp.MustCompile("^http[s]?://[^/]+")
    r_html := regexp.MustCompile("[.][a-zA-Z0-9]+$")
    r_ext := regexp.MustCompile("[.]net|[.]org")


    testo1 := r_url.ReplaceAllString(testo, "")
    fmt.Println("1",testo1)
    testo2 := r_html.FindString(testo1)
    fmt.Println("2",testo2)
    if r_ext.MatchString(testo2) == true{
	fmt.Println("3","HTML")
    } else {
	fmt.Println("3","OTHER")
    }

}