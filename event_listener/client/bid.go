package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BidClient struct {
	host string
	path string
}

func newBidClient(host string, path string) *BidClient {
	return &BidClient{
		host: host,
		path: path,
	}
}

func (b *BidClient) UpdateOne(id string, status int) bool {
	fullPath := fmt.Sprintf("%v%v/%v", b.host, b.path, id)
	payload, err := json.Marshal(map[string]interface{}{
		"status": status,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println(payload)
	client := &http.Client{}

	responseBody := bytes.NewBuffer(payload)
	req, err := http.NewRequest(http.MethodPatch, fullPath, responseBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	sb := string(body)
	fmt.Println(sb)
	return true
}
