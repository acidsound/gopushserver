/**
 * Created with IntelliJ IDEA.
 * User: acidsound
 * Date: 13. 3. 10.
 * Time: 오후 8:58
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	apns "github.com/pranavraja/apns"
	gcm "github.com/googollee/go-gcm"
	"fmt"
	"os"
)

var c map[string]interface{}

func LoadConfiguration() {
	FileBody, Err := ioutil.ReadFile("./config.json")
	if Err != nil {
		panic(Err)
	}
	Err = json.Unmarshal(FileBody, &c)
	if Err != nil {
		panic(Err)
	}

	log.Printf("API_KEY=%T|%v\n", c["APIKey"], c["APIKey"])
	log.Printf("SERVER_PORT=%T|%v", c["serverURL"], c["serverURL"])
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		targetDevice := r.FormValue("target")
		targetOS := r.FormValue("os")
		if targetOS == "ios" {
			result = SendAPNSNotification(targetDevice, string(body))
		} else {
			result = SendGCMNotification(targetDevice, string(body))
		}
		fmt.Fprintf(w, "{ result: \"%v\" }", result)
	}
}
func SendGCMNotification(target string, message string) string {
	// API Key
	client := gcm.New(c["APIKey"].(string))

	// client ID
	log.Printf("----------------------------------------------------------")
	log.Printf("sendGCMNotification regi_ID [%v]\n", target)
	log.Printf("sendGCMNotification payload [%v]\n", message)
	load := gcm.NewMessage(target)
	// load.AddRecipient("abc")
	load.SetPayload("data", message)
	// load.CollapseKey = "registerCar"
	load.DelayWhileIdle = true
	load.TimeToLive = 10

	resp, err := client.Send(load)

	log.Printf("sendResult id: %+v\n", resp)
	if err != nil {
		log.Println("err:", err)
		return "fail"
	}
	return "success"
}

func SendAPNSNotification(target string, message string) string {
	host := os.Getenv("APNS_HOST")
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")
	queue := []apns.NotificationAndPayload{apns.MakeNotification(1, "aef4429b", "message")}
	// Blocks until notifications are all sent
	err := apns.ConnectAndSend(host, certFile, keyFile, queue)
	if err != nil {
		log.Println("err:", err)
		return "fail"
	}
	return "success"
}
