package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type EntityHttpRequest struct {
	URL                 string
	Method              string
	UserAgent           string
	ContentType         string
	Cookies             []http.Cookie
	Data                io.Reader
	Sleep               time.Duration
	TlsCertificateCheck bool

	TimeIn time.Time
}

type HttpResponse struct {
	URL       string
	Method    string
	UserAgent string
	BRaw      string

	Duration time.Duration
}

func NewHttp() *EntityHttpRequest {
	var entity = &EntityHttpRequest{
		UserAgent:   "GNET - Advanced Technology Research Institute",
		Data:        nil,
		ContentType: "text/html; charset=UTF-8",
	}

	return entity
}

func (e *EntityHttpRequest) SetURL(URL string) error {
	if ThisIsURL(URL) != nil {
		return fmt.Errorf("This URL is invalid for use.")
	}

	e.URL = URL

	return nil
}

func (e *EntityHttpRequest) SetMethod(method string) error {
	switch method {
	case "GET":
		e.Method = "GET"
	case "POST":
		e.Method = "POST"
	case "DELETE":
		e.Method = "DELETE"
	case "PUT":
		e.Method = "PUT"
	default:
		return fmt.Errorf("Non-existent or unacceptable method.")
	}

	return nil
}

func (e *EntityHttpRequest) SetUserAgent(agent string) {
	e.UserAgent = agent
}

func (e *EntityHttpRequest) SetData(data string) {
	e.Data = strings.NewReader(data)
}

func (e *EntityHttpRequest) SetContentType(content string) {
	e.ContentType = content
}

func (e *EntityHttpRequest) SetSleep(time time.Duration) {
	e.Sleep = time
}

func (e *EntityHttpRequest) SetCookies(cookies []http.Cookie) {
	for _, c := range cookies {
		e.Cookies = append(e.Cookies, c)
	}
}

func (e *EntityHttpRequest) OnFirewallDetect() {}

func (e *EntityHttpRequest) Do() (*HttpResponse, error) {
	var client = &http.Client{
		CheckRedirect: nil,
		Transport: &http.Transport{
			Proxy: nil,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: e.TlsCertificateCheck,
			},
		},
	}

	e.TimeIn = time.Now()

	var r, err = http.NewRequest(e.Method, e.URL, e.Data)

	if err != nil {
		return nil, err
	}

	r.Header.Set("User-Agent", e.UserAgent)
	r.Header.Set("Content-Type", e.ContentType)

	response, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	braw, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	defer time.Sleep(time.Duration(e.Sleep) * time.Second)

	return &HttpResponse{URL: e.URL, Method: e.Method, UserAgent: response.Request.UserAgent(), Duration: time.Since(e.TimeIn), BRaw: string(braw)}, nil
}
