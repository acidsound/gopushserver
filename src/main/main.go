/**
 * Created with IntelliJ IDEA.
 * User: acidsound
 * Date: 13. 3. 10.
 * Time: 오후 8:23
 * To change this template use File | Settings | File Templates.
 */
// +build !appengine
package main

import (
	"log"
	"net/http"
)
func main() {
	log.Fatal(http.ListenAndServe(string(c["serverURL"].(string)), nil))
}
