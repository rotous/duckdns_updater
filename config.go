package main

import (
	"os"
	"log"
	"strconv"
)

type config struct {
	url string
	interval int
	key string
	domains string
}

// returns a config object with values it got from
// the environment or default values
func getConfig() config {
	const DUCKDNS_DEFAULT_UPDATE_URL = "https://www.duckdns.org/update?domains=%s&token=%s&verbose=true"
	const DUCKDNS_DEFAULT_INTERVAL = 5 * 60  // Interval in seconds

	// Get the update url
	url, set := os.LookupEnv("DUCKDNS_UPDATE_URL")
	if !set {
		url = DUCKDNS_DEFAULT_UPDATE_URL
	}

	// Get the update interval
	interval := DUCKDNS_DEFAULT_INTERVAL
	env, set := os.LookupEnv("DUCKDNS_INTERVAL")
	if set {
		interv, err := strconv.Atoi(env)
		if err != nil {
			log.Println("Unable to parse interval from DUCKDNS_INTERVAL: ", err)
		} else if interv < 1 {
			log.Panic("Cannot set an interval smaller than 1 second")
		} else {
			// a parsable interval value was set as ENV var
			interval = interv
		}
	}

	// get the key
	key, set := os.LookupEnv("DUCKDNS_KEY")
	if !set {
		log.Panic("No key specified! Please use DUCKDNS_KEY env var to set it.")
	}

	// get the domains
	domains,set := os.LookupEnv("DUCKDNS_DOMAINS")
	if !set {
		log.Panic("No domains specified! Please use DUCKDNS_DOMAINS env var to set it.")
	}

	return config {
		url,
		interval,
		key,
		domains,
	}
}