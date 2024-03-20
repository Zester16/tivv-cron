package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func PostCrawlGetBloombergNewsLetter(url string) (string, error) {

	const logPrefix = "network/bloomberg-crawler-go"
	body := NYTNewsBodyStruct{NewsUrl: url}
	bodyMarshal, _ := json.Marshal(body)

	URL := os.Getenv("blm_cron")
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(bodyMarshal))

	if err != nil {
		fmt.Println(logPrefix+" request", err)
	}

	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {

		fmt.Println(logPrefix+" http error", err)
		return "", err
	}

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
