package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Cert holds definition for certs
type Cert struct {
	Domain string
}

func getExpired() []Cert {

	var body, err = Get(fmt.Sprintf("%s/expired", serviceURL))
	if err != nil {
		log.Printf("Unable to get expired certs, will try later")
		return []Cert{}
	}

	var deserialized []Cert

	err = json.Unmarshal(body, &deserialized)
	if err != nil {
		log.Panicf("json serialization failing %s", err.Error())
	}

	return deserialized
}

func renewCert(domain string) {

	var _, err = Get(fmt.Sprintf("%s/renew/%s", serviceURL, domain))
	if err != nil {
		log.Printf("Unable to renew cert, will try later. Error: %s", err)
	}
}

var serviceURL = "http://localhost:8088"

func main() {

	for {

		fmt.Print("Running cert watch...\n")

		var certs = getExpired()

		for _, r := range certs {

			fmt.Printf("Renewing %s\n", r.Domain)

			renewCert(r.Domain)

		}

		time.Sleep(10 * time.Second)
	}

}
