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
			e := &HttpRequest{}
			err := e.SetMethod(c.method)

			switch c.valid {
			case true:
				if err != nil {
					t.Errorf("Unexpected error for %s method: %v", c.method, err)
				}
				if e.Method != c.method {
					t.Errorf("SetMethod() didn't set the method correctly. Expected: %s, Got: %s", c.method, e.Method)
				}
			case false:
				if err == nil {
					t.Errorf("Expected error for %s method, but got none", c.method)
				}
			}
		})
	}
}

