package gnet

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpRequest struct {
	URL                 string
	Method              string
	UserAgent           string
	ContentType         string
	Data                io.Reader
	Sleep               time.Duration
	Cookies             []http.Cookie
	Proxy               func(*http.Request) (*url.URL, error)
	TlsCertificateCheck bool

	TimeIn time.Time

	EntityFeatures struct {
		OnRandomUserAgent bool
		OnTor             bool
	}
}

type HttpResponse struct {
	URL           string
	Method        string
	UserAgent     string
	BRaw          string
	StatusCode    int
	ContentLength int64

	Duration time.Duration
	Request  HttpRequest
}

func NewHttp() *HttpRequest {
	var entity = &HttpRequest{
		UserAgent:   "GNET - Advanced Technology Research Institute",
		Data:        nil,
		ContentType: "text/html; charset=UTF-8",
		EntityFeatures: struct {
			OnRandomUserAgent bool
			OnTor             bool
		}{false, false},
	}

	return entity
}

func (e *HttpRequest) SetURL(URL string) error {
	if ThisIsURL(URL) != nil {
		return fmt.Errorf("this url is invalid for use")
	}

	e.URL = URL

	return nil
}

func (e *HttpRequest) SetMethod(method string) error {
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
		return fmt.Errorf("non-existent or unacceptable method")
	}

	return nil
}

func (e *HttpRequest) SetUserAgent(agent string) {
	e.UserAgent = agent
}

func (e *HttpRequest) SetData(data string) {
	e.Data = strings.NewReader(data)
}

func (e *HttpRequest) SetContentType(content string) {
	e.ContentType = content
}

func (e *HttpRequest) SetSleep(time time.Duration) {
	e.Sleep = time
}

func (e *HttpRequest) SetCookies(cookies []http.Cookie) {
	e.Cookies = append(e.Cookies, cookies...)
}

func (e *HttpRequest) SetProxy(proxy string) error {
	var proxyURL, err = url.Parse(proxy)

	e.Proxy = http.ProxyURL(proxyURL)

	return err
}

func (e *HttpRequest) OnRandomUserAgent() {
	e.EntityFeatures.OnRandomUserAgent = true
}

func (e *HttpRequest) OnTor() {
	e.EntityFeatures.OnTor = true

	e.SetProxy(TORURI)
}

func (e *HttpRequest) OnFirewallDetect() {}

func (e *HttpRequest) Do() (*HttpResponse, error) {
	var client = &http.Client{
		CheckRedirect: nil,
		Transport: &http.Transport{
			Proxy: e.Proxy,
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

	return &HttpResponse{
		URL:           e.URL,
		Method:        e.Method,
		StatusCode:    response.StatusCode,
		ContentLength: response.ContentLength,
		UserAgent:     response.Request.UserAgent(),
		Duration:      time.Since(e.TimeIn),
		BRaw:          string(braw),
		Request:       *e,
	}, nil
}
