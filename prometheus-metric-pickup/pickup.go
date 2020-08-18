package main

import (
	"encoding/json"
	"github.com/monaco-io/request"
	"log"
	"time"
)

func main() {
	url := ""
	timeStrings := getTimeStringInMonth()
	queryStringMap := getQueryStringMap()
	for _, timeString := range timeStrings {
		for k, v := range queryStringMap {
			log.Printf("%s - %s : %s", k, timeString, query(url, v, timeString))
		}
	}

	// query

	// queryByInstance

}

func query(url string, queryString string, timeString string) string {
	client := request.Client{
		URL:    url,
		Method: request.GET,
		Params: map[string]string{"query": queryString, "time": timeString},
	}

	resp, err := client.Do()
	if err != nil {
		log.Fatal(resp.Code, string(resp.Data), err)
	}

	response := VectorResponse{}
	err = json.Unmarshal(resp.Data, &response)

	if response.Status == "success" {
		for _, r := range response.Data.Result {
			for _, v := range r.Value {
				switch value := v.(type) {
				default:
				case string:
					return value
				}
			}
		}
	}

	return ""
}

type VectorResponse struct {
	Data struct {
		Result []struct {
			Metric struct{}      `json:"metric"`
			Value  []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

func getTimeStringInMonth() []string {
	var timeStrings []string
	now := time.Now()
	for day := 1; day <= now.Day(); day++ {
		timeStrings = append(timeStrings, time.Date(now.Year(), now.Month(), day, 0, 0, 0, 0, time.UTC).Format(time.RFC3339))
	}
	return timeStrings
}

func getQueryStringMap() map[string]string {
	return map[string]string{"key": "value"} // TODO: fill query string without instance through configuration
}

func getQueryStringByInstanceMap() map[string]string {
	return map[string]string{"key": "value"} // TODO: fill query string with instance through configuration
}

func getInstanceArray() map[string]string {
	return map[string]string{"key": "value"} // TODO: fill instance through configuration
}
