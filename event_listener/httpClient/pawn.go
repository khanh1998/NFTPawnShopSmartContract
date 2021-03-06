package httpClient

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
	path string
}

func newPawnClient(host string, path string) *PawnClient {
	return &PawnClient{
		host: host,
		path: path,
	}
}

func (p *PawnClient) InsertOne(id string, creator string, tokenAdd string, tokenId string, status uint8) (bool, string) {
	fullPath := fmt.Sprintf("%v%v", p.host, p.path)
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
		return false, ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		return false, ""
	}
	sb := string(body)
	fmt.Println(sb)
	return true, sb
}

func (p *PawnClient) UpdateOne(id string, status int, bidId string) (bool, string) {
	log.Println("status and bid id", status, bidId)
	fullPath := fmt.Sprintf("%v%v/%v", p.host, p.path, id)
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
	return true, sb
}
