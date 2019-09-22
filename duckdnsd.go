package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"strconv"
)

const DUCKDNS_DEFAULT_UPDATE_URL = "https://www.duckdns.org/update?domains=%s&token=%s&verbose=true"
const DUCKDNS_DEFAULT_INTERVAL = 5 * 60  // Interval in seconds

func main () {
	log.Output(1, "Starting duckdns daemon...")

	url := os.Getenv("DUCKDNS_UPDATE_URL")
	if url == "" {
		url = DUCKDNS_DEFAULT_UPDATE_URL
	}
	log.Println("Update URL set to ", url)

	envInterval := os.Getenv("DUCKDNS_INTERVAL")
	interval := DUCKDNS_DEFAULT_INTERVAL
	if envInterval == "" {
		log.Println("No internval set in ENV")
	} else {
		interv, err := strconv.Atoi(envInterval)
		if err != nil {
			log.Println("Unable to parse interval from DUCKDNS_INTERVAL: ", err)
		} else if interv < 1 {
			log.Println("Cannot set an interval smaller than 1 second")
		} else {
			interval = interv
		}
	}
	log.Println("Interval set to", interval, "seconds")

	// get the key
	key := os.Getenv("DUCKDNS_KEY")
	if key == "" {
		log.Panic("No key specified")
	}

	domains := os.Getenv("DUCKDNS_DOMAINS")
	if domains == "" {
		log.Panic("No domains specified")
	}

	for {
		update(url, key, domains)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func update (url string, key string, domains string) {
	resp, err := http.Get(fmt.Sprintf(url, domains, key))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	bodyBuf, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBuf)

	if body[0:2] == "OK" {
		lines := strings.Split(body, "\n")
		if len(lines) >= 4 && strings.TrimSpace(lines[3]) == "NOCHANGE" {
			log.Output(1, fmt.Sprintf("IP address for %s unchanged, remains %s", domains, lines[1]))
		} else {
			log.Output(1, fmt.Sprintf("IP address for %s updated: %s", domains, lines[1]))
		}
	} else {
		log.Output(1, fmt.Sprintf("Update for %s failed", domains))
	}
}
