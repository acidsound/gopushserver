// +build !appengine
package main

import (
	"log"
	"net/http"
)
func main() {
	log.Fatal(http.ListenAndServe(string(c["serverURL"].(string)), nil))
}
