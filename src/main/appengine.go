package main

/*
 * test curl
 * curl -d "요청하신 견적이 도착하였습니다." http://localhost:8001/sendMessage?target=APA91bHMYVVLT9C4MJ52TTwcUSa5ZndRk6bN0gZ01BVDtZi5zX9-Fv-CES2VuX7P9CpqT4tEqEP7pGRhNByTp9b8mQb195KqJ9iPpKso90YcPW4bB0ErBf_u2G70stW_yLGQgypiz3r4
 */

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





