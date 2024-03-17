package ghttp

import (
	"encoding/json"
	"fmt"
)

const TORURI = "socks5://127.0.0.1:9050"

// CheckConnectionTor checks the connection status to the Tor network by making a request to the Tor Project API.
func CheckConnectionTor() (string, error) {
	var marshal map[string]interface{}

	request := NewHttp()
	err := request.SetURL("https://check.torproject.org/api/ip")
	if err != nil {
		return "", err
	}

	err = request.SetMethod("GET")
	if err != nil {
		return "", err
	}

	err = request.OnTor()
	if err != nil {
		return "", err
	}

	response, err := request.Do()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(response.BRaw), &marshal)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(marshal["IP"]), nil
}
