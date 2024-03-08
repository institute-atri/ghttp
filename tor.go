package gnet

import (
	"encoding/json"
	"fmt"
)

const TORURI = "socks5://127.0.0.1:9050"

func CheckConnectionTor() (string, error) {
	var marshal map[string]interface{}

	var request = NewHttp()
	request.SetURL("https://check.torproject.org/api/ip")
	request.SetMethod("GET")
	request.OnTor()

	var response, err = request.Do()

	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(response.BRaw), &marshal)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(marshal["IP"]), nil
}
