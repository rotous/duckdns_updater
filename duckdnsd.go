package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
)

func main () {
	log.Output(1, "Starting duckdns daemon...")

	config := getConfig()
	log.Println("Domains set to:", config.domains)
	log.Println("Update url set to:", config.url)
	log.Println("Update interval set to:", config.interval)

	for {
		update(config)
		time.Sleep(time.Duration(config.interval) * time.Second)
	}
}

func update (cfg config) {
	resp, err := http.Get(fmt.Sprintf(cfg.url, cfg.domains, cfg.key))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	bodyBuf, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBuf)

	if body[0:2] == "OK" {
		lines := strings.Split(body, "\n")
		if len(lines) >= 4 && strings.TrimSpace(lines[3]) == "NOCHANGE" {
			log.Printf("IP address for %s unchanged, remains %s", cfg.domains, lines[1])
		} else {
			log.Printf("IP address for %s updated to %s", cfg.domains, lines[1])
		}
	} else {
		log.Printf("Update for %s failed", cfg.domains)
	}
}
