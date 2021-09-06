package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type NotifyClient struct {
	host string
	path string
}

type Notification struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	PawnID   string `json:"pawnid"`
	BidID    string `json:"bidid"`
	Borrower string `json:"borrower"`
	Lender   string `json:"lender"`
	Payload  string `json:"payload"`
}

func newNotifyClient(host string, path string) *NotifyClient {
	return &NotifyClient{
		host: host,
		path: path,
	}
}

func (n *NotifyClient) SendNotification(noti Notification) bool {
	fullPath := fmt.Sprintf("%v%v", n.host, n.path)
	payload, _ := json.Marshal(noti)
	res, err := http.Post(fullPath, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Panic(err)
		return false
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
		return false
	}
	fmt.Println(string(resBody))
	return true
}
