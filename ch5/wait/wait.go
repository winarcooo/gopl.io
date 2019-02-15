package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer attemps to contact the server of a URL
// it tries for on minute using exponential back-off
// it reports an error if all atempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-of
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
