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

func NewBidClient(host string, path string) *BidClient {
	return &BidClient{
		host: host,
		path: path,
	}
}

func (b *BidClient) Post(
	bidId string, creator string, loanAmount string, interest string,
	startTime string, duration string, proRated bool, pawnId string,
) bool {
	fullPath := fmt.Sprintf("%v%v", b.host, b.path)
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":              bidId,
		"creator":         creator,
		"loan_amount":     loanAmount,
		"interest":        interest,
		"loan_start_time": startTime,
		"loan_duration":   duration,
		"pro_rated":       proRated,
		"pawn":            pawnId,
	})
	log.Println(postBody)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(fullPath, "application/json", responseBody)
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
