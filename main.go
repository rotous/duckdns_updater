package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"io/ioutil"
)

const DUCKDNS_UPDATE_URL = "https://www.duckdns.org/update?domains=%s&token=%s"

func main () {
	fmt.Println("Updating DuckDNS...")

	// get the key
	key := os.Getenv("DUCKDNS_KEY")
	if key == "" {
		log.Panic("No key specified")
	}

	domains := os.Getenv("DUCKDNS_DOMAINS")
	if domains == "" {
		log.Panic("No domains specified")
	}

	resp, err := http.Get(fmt.Sprintf(DUCKDNS_UPDATE_URL, domains, key))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	bodyBuf, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBuf)

	if body == "KO" {
		log.Output(1, "Update failed")
	} else {
		fmt.Printf("Updated ip address for %s\n", domains)
		fmt.Println(body)
	}
}
