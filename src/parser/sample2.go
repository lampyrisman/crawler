package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`ab`)
    fmt.Println(re.FindAllStringIndex("123ab678ab90", -1))
    fmt.Println(re.FindAllStringIndex("foo", -1) == nil)
}
