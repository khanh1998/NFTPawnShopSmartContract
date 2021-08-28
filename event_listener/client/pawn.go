package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PawnClient struct {
	host string
}

const (
	Path = "/pawns"
)

func NewPawnClient(host string) *PawnClient {
	return &PawnClient{
		host: host,
	}
}

func (p *PawnClient) Post(id string, creator string, tokenAdd string, tokenId string, status uint8) bool {
	fullPath := fmt.Sprintf("%v%v", p.host, Path)
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":            id,
		"creator":       creator,
		"token_address": tokenAdd,
		"token_id":      tokenId,
		"status":        status,
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

func (p *PawnClient) Patch(id string, status int, bidId string) bool {
	log.Println("status and bid id", status, bidId)
	fullPath := fmt.Sprintf("%v%v/%v", p.host, Path, id)
	payload, err := json.Marshal(map[string]interface{}{
		"status": status,
		"bid":    bidId,
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
