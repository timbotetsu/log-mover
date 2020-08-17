package main

import (
	"github.com/monaco-io/request"
	"log"
)

func main() {
}

func query(url string, querystring string, timestring string) {
	client := request.Client{
		URL:    url,
		Method: request.GET,
		Params: map[string]string{"query": querystring, "time": timestring},
	}
	resp, err := client.Do()

	log.Println(resp.Code, string(resp.Data), err)
	// TODO: convert response data to struct
}
