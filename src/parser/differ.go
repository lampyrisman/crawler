package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v8"
	"hash/crc32"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)
const r_url_v string = "^http[s]?://[^/]+"
const r_html_v string = "[.][a-zA-Z0-9]{1,4}$"
const r_ext_v string = "[.]net$|[.]com$|[.]org$|[.]ru$|[.][s]?htm[l]?$"
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var rdb_bu = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       1,  // use default DB
})
var ctx = context.Background()

func checkErr(err error) {
	if err != nil && err.Error() != "nil" {
		fmt.Print("Err--")
		fmt.Print(err)
		fmt.Println("--")
		panic(err)
	}
}




func chkNew(url string, content []byte) (isNew bool) {
	var oldCRC, newCRC uint32
	val, err := rdb.Get(ctx, url).Result()
	    checkErr(err)

	newCRC = crc32.ChecksumIEEE(content)
	if val == "" {
		err := rdb.Set(ctx, url, newCRC, 0).Err()
		checkErr(err)
//		fmt.Println("CHECK: ", val, "NONE, adding...")
	} else {
		oldCRCint, err := strconv.Atoi(val)
		checkErr(err)
		oldCRC = uint32(oldCRCint)
	}
//	fmt.Println("CHECK: ", url, "Old: ", oldCRC, "New: ", newCRC)
	if newCRC == oldCRC {
		return false
	}
	return true
}

func getLinks(url string, baseUrl string) (urlList []string) {
	var isRel int
	resp, err := http.Get(url)
	if resp.StatusCode != 200 { return urlList}
	fmt.Println(strings.Split(resp.Header["Content-Type"][0], `;`)[0])
	checkErr(err)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fromRead := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(fromRead)
	checkErr(err)
	doc.Find("a").Each(func(i int, s1 *goquery.Selection) {
		link, ok := s1.Attr("href")
		if ok {
			isDrop := 1
			if getBaseUrl(link) != "" {
			urlList = append(urlList, link) 
			isDrop = 0
		    }
//			fmt.Print("REL: ", link, " ")
			isRel = strings.Index(link, `//`)
			if isRel == 0 {
			isDrop = 0
				link = "https:" + link
//				fmt.Println(link)
				urlList = append(urlList, link)
			}

			isRel = strings.Index(link, `/`)
			if isRel == 0 {
			isDrop = 0
				link = baseUrl + link
//				fmt.Println(link)
				urlList = append(urlList, link)
			}

			isRel = strings.Index(link, `./`)
			if isRel == 0 {
			isDrop = 0
				link = url + link
//				fmt.Println(link)
				urlList = append(urlList, link)
			}
		    if isDrop == 1 { fmt.Println("DROP:", link) }
		}
	})
	return urlList
}

func redisScanAll(pos string) (newpos string, data []string) {
	val, err := rdb.Do(ctx, "SCAN", pos).Result()
	checkErr(err)
	val1 := val.([]interface{})
	newpos = val1[0].(string)
	for _, dataone := range val1[1].([]interface{}) {
		data = append(data, dataone.(string))
	}
	return newpos, data
}

func setChildUrl (url string)(){
	
	val, _ := rdb.Get(ctx, url).Result()
	if val == "" {
	err := rdb.Set(ctx, url, 0, 0).Err()
	checkErr(err)
    } else {
}
}


func getBaseUrl(url string)(baseUrl string){
    r_url := regexp.MustCompile(r_url_v)
    baseUrl = r_url.FindString(url)
    fmt.Println("GetBase", url, baseUrl)
    return baseUrl
}

func checkLink(url string)(status bool){
    r_url := regexp.MustCompile(r_url_v)
    r_html := regexp.MustCompile(r_html_v)
    r_ext := regexp.MustCompile(r_ext_v)

    step1 := r_url.ReplaceAllString(url, "")
//    fmt.Println("1",step1)
    step2 := r_html.FindString(step1)
//    fmt.Println("2",step2)
    if r_ext.MatchString(step2) == true || step2 == ""{
//    fmt.Println("3","HTML")
    return true
}
    fmt.Println("3","OTHER", url)
    return false

}


func main() {
	fmt.Println(time.Now().Unix())
	var next string = "0"
	var data []string

	for {
		next, data = redisScanAll(next)
		if len(data) == 0 {
			continue
		}
		for _, url := range data {
			if checkLink(url) == false { continue }
			baseUrl := getBaseUrl(url)
			err := rdb_bu.Set(ctx, baseUrl, 0, 0).Err()
			checkErr(err)
			urlList := getLinks(url, baseUrl)
			if len(urlList) == 0 {
				continue
			}
			isNew := chkNew(url, []byte(strings.Join(urlList, "")))
//			fmt.Println(urlList[0], isNew)
			if isNew == true {
				for _, setUrl := range urlList {
				if checkLink(setUrl) == false { continue }
				    setChildUrl(setUrl)
				}
			}
		}
		if next == "0" {
			break
		}

	}
}
