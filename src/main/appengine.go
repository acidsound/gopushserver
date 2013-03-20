package main

import (
	"net/http"
	"log"
)
func init() {
	log.Println(">> gopushserver : robust go push notification server\n")
	log.Println("initialize with push server configuration..\n\n")
	LoadConfiguration()
	http.HandleFunc("/sendMessage", SendMessageHandler)

	log.Printf(">> push notification server started on %v\n", c["serverURL"].(string))
}





