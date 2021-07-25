// Belirtilen Websitenin kaynak kodunu çekmektedir.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://facebook.com")

	if err != nil {
		log.Fatal("Hata!", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	fmt.Print(string(data))
}
