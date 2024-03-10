package gnet

import (
	"fmt"
	"log"
	"testing"
)

func TestSetMethod(t *testing.T) {
	cases := []struct {
		method string
		valid  bool
	}{
		{"GET", true},
		{"POST", true},
		{"PUT", true},
		{"DELETE", true},
		{"INVALID_METHOD", false},
	}

	for _, c := range cases {
		t.Run(c.method, func(t *testing.T) {
			e := &EntityHttpRequest{}
			err := e.SetMethod(c.method)

			switch {
			case c.valid && err != nil:
				t.Errorf("Unexpected error for %s method: %v", c.method, err)
			case !c.valid && err == nil:
				t.Errorf("Expected error for %s method, but got none", c.method)
			}
		})
	}
}

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
