package gnet

import (
	"fmt"
	"log"
	"testing"
)

func TestPost(t *testing.T) {
	var response = POST("http://httpbin.org/post", "user=admin&pass=123")

	log.Println(response.BRaw)
	log.Println("Request Time: " + fmt.Sprint(response.Duration.Seconds()))
}

func TestGet(t *testing.T) {
	var response = GET("http://httpbin.org/get")

	log.Println(response.BRaw)
	log.Println("Request Time: " + fmt.Sprint(response.Duration.Seconds()))
}

func TestGetProxy(t *testing.T) {
	var request = NewHttp()

	request.SetURL("http://httpbin.org/get")
	request.SetMethod("GET")
	request.OnTor()

	var response, err = request.Do()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.BRaw)

	log.Println(CheckConnectionTor())
	log.Println("Request Time: " + fmt.Sprint(response.Duration.Seconds()))
}
