package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BidPawnClient struct {
	host string
	path string
}

func newBidPawnClient(host string, path string) *BidPawnClient {
	return &BidPawnClient{
		host: host,
		path: path,
	}
}

func (b *BidPawnClient) UpdateOne(bidId string) bool {
	fullPath := fmt.Sprintf("%v%v/%v", b.host, b.path, bidId)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, fullPath, nil)
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
