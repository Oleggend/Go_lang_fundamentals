package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://oleg-mgmt.auto.sentinelone.local")
	page, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", page)
}
