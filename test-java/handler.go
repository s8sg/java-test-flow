package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// makeQueryStringFromParam create query string from provided query
func makeQueryStringFromParam(params map[string][]string) string {
	if params == nil {
		return ""
	}
	result := ""
	for key, array := range params {
		for _, value := range array {
			keyVal := fmt.Sprintf("%s-%s", key, value)
			if result == "" {
				result = "?" + keyVal
			} else {
				result = result + "&" + keyVal
			}
		}
	}
	return result
}

// buildHttpRequest build upstream request for function
func buildHttpRequest(url string, method string, data []byte, params map[string][]string,
	headers map[string]string) (*http.Request, error) {

	queryString := makeQueryStringFromParam(params)
	if queryString != "" {
		url = url + queryString
	}

	httpreq, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		httpreq.Header.Add(key, value)
	}

	return httpreq, nil
}

// Handle a serverless request
func Handle(req []byte) string {
	gateway := os.Getenv("gateway")
	if gateway == "" {
		gateway = "gateway:8080"
	}
	function := os.Getenv("date-function")
	if function == "" {
		function = "date-java"
	}

	url := "http://" + gateway + "/function/" + function

	// Start

	/*
		resp, err := http.Post(url, "application/json", bytes.NewReader(req))
		if err != nil {
			log.Fatal(err)
		}
	*/

	respData := req
	for i := 0; i < 5; i++ {

		params := make(map[string][]string)
		headers := make(map[string]string)

		httpreq, err := buildHttpRequest(url, "POST", respData, params, headers)
		if err != nil {
			log.Fatalf("cannot connect to Function on URL: %s", url)
		}

		client := &http.Client{}
		resp, err := client.Do(httpreq)

		// End

		defer resp.Body.Close()
		var result []byte
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			result, _ = ioutil.ReadAll(resp.Body)
			err = fmt.Errorf("invalid return status %d while connecting %s", resp.StatusCode, url)
		} else {
			result, err = ioutil.ReadAll(resp.Body)
		}
		if err != nil {
			log.Fatal(err)
		}

		respData = result
	}
	return string(respData)
}
