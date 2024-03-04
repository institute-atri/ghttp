package main

import (
	"log"
	"testing"
)

func TestPost(t *testing.T) {
	var response = POST("http://httpbin.org/post", "user=admin&pass=123")

	log.Println(response.BRaw)
}

func TestGet(t *testing.T) {
	var response = GET("http://httpbin.org/get")

	log.Println(response.BRaw)
}
