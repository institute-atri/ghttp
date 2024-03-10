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

type EntityHttpRequest struct {
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
		EntityFeatures: struct {
			OnRandomUserAgent bool
			OnTor             bool
		}{false, false},
	}

	return entity
}

func (e *EntityHttpRequest) SetURL(URL string) error {
	if ThisIsURL(URL) != nil {
		return fmt.Errorf("this url is invalid for use")
	}

	e.URL = URL

	return nil
}

func (e *EntityHttpRequest) SetMethod(method string) error {
	httpRequests := map[string]bool{
		"GET":    true,
		"POST":   true,
		"PUT":    true,
		"DELETE": true,
	}
	
	if request := httpRequests[method]; request {
		e.Method = method
		return nil
	}
	return fmt.Errorf("non-existent or unacceptable method")
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
	e.Cookies = append(e.Cookies, cookies...)
}

func (e *EntityHttpRequest) SetProxy(proxy string) error {
	var proxyURL, err = url.Parse(proxy)

	e.Proxy = http.ProxyURL(proxyURL)

	return err
}

func (e *EntityHttpRequest) OnRandomUserAgent() {
	e.EntityFeatures.OnRandomUserAgent = true
}

func (e *EntityHttpRequest) OnTor() {
	e.EntityFeatures.OnTor = true

	e.SetProxy(TORURI)
}

func (e *EntityHttpRequest) OnFirewallDetect() {}

func (e *EntityHttpRequest) Do() (*HttpResponse, error) {
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

	return &HttpResponse{URL: e.URL, Method: e.Method, UserAgent: response.Request.UserAgent(), Duration: time.Since(e.TimeIn), BRaw: string(braw)}, nil
}
