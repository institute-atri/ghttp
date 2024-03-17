package ghttp

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpRequest represents an HTTP request entity with various properties and methods.
type HttpRequest struct {
	URL                 string
	Method              string
	UserAgent           string
	ContentType         string
	Data                io.Reader
	Sleep               time.Duration
	Cookies             []http.Cookie
	Proxy               func(*http.Request) (*url.URL, error)
	Redirect            func(*http.Request, []*http.Request) error
	TlsCertificateCheck bool

	TimeIn time.Time

	EntityFeatures struct {
		OnRandomUserAgent bool
		OnTor             bool
	}
}

// HttpResponse represents an HTTP response entity with various properties and methods.
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

// NewHttp initializes a new HttpRequest object with default values and returns it.
// The HttpRequest struct is used to send HTTP requests and handle the responses.
// The returned object can be further configured by calling its setter methods.
// Example usage:
//
//	r := NewHttp()
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func NewHttp() *HttpRequest {
	entity := &HttpRequest{
		UserAgent:   "GHTTP",
		Data:        nil,
		ContentType: "text/html; charset=UTF-8",
		EntityFeatures: struct {
			OnRandomUserAgent bool
			OnTor             bool
		}{false, false},
	}

	return entity
}

// SetURL sets the URL for the HttpRequest object.
// It validates the URL using the ThisIsURL function.
// If the URL is invalid, it returns an error.
// Otherwise, it sets the URL and returns nil.
//
// Example usage:
//
//	r := NewHttp()
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetURL(URL string) error {
	if ThisIsURL(URL) != nil {
		return fmt.Errorf("this url is invalid for use")
	}

	e.URL = URL

	return nil
}

// SetMethod sets the request method for the HttpRequest object.
// It accepts "GET", "POST", "DELETE", or "PUT" as valid methods.
// If the method is not one of the valid options, it returns an error.
// Otherwise, it sets the method and returns nil.
// Example usage:
//
//	r := NewHttp()
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//	// Do something with the response
//
// Do not duplicate the examples provided above.
func (e *HttpRequest) SetMethod(method string) error {
	switch method {
	case "GET", "POST", "DELETE", "PUT":
		e.Method = method
	default:
		return fmt.Errorf("non-existent or unacceptable method")
	}
	return nil
}

// SetUserAgent sets the User-Agent header for the HttpRequest object.
//
// Example usage:
//
//	r := NewHttp()
//	r.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36")
//
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetUserAgent(agent string) {
	e.UserAgent = agent
}

// SetData sets the data for the HttpRequest object.
// It initializes the Data field of the HttpRequest object with a new Reader containing the provided data.
//
// Example usage:
//
//	r := NewHttp()
//	err := r.SetURL(URL)
//	if err != nil {
//		// Handle the error
//	}
//
//	err = r.SetMethod("POST")
//	if err != nil {
//		// Handle the error
//	}
//
//	r.SetData(data)
//	r.SetContentType("application/x-www-form-urlencoded")
//
//	response, err := r.Do()
//	if err != nil {
//		// Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetData(data string) {
	e.Data = strings.NewReader(data)
}

// SetContentType sets the content type for the HttpRequest object.
//
// Example usage:
//
//	r := NewHttp()
//	r.SetContentType("application/x-www-form-urlencoded")
func (e *HttpRequest) SetContentType(content string) {
	e.ContentType = content
}

// SetSleep sets the sleep duration for the HttpRequest object.
// The sleep duration is used to delay the execution of the request.
// The value is specified in seconds.
//
// Example usage:
//
//	r := NewHttp()
//	r.SetSleep(5) // Sleep for 5 seconds before executing the request
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetSleep(time time.Duration) {
	e.Sleep = time
}

// SetCookies adds the provided cookies to the HttpRequest object's cookie list.
// It appends the cookies to the existing list.
//
// Example usage:
//
//	r := NewHttp()
//
//	// Create cookies
//	cookie1 := http.Cookie{Name: "cookie1", Value: "value1"}
//	cookie2 := http.Cookie{Name: "cookie2", Value: "value2"}
//
//	// Set cookies
//	r.SetCookies([]http.Cookie{cookie1, cookie2})
//
//	// Perform request
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetCookies(cookies []http.Cookie) {
	e.Cookies = append(e.Cookies, cookies...)
}

// SetProxy sets the proxy for the HttpRequest object.
// It takes a string parameter "proxy" representing the proxy URL.
// The proxy URL is parsed and assigned to the Proxy field of the HttpRequest object.
// If there is an error in parsing the proxy URL, it is returned.
//
// Example usage:
//
//	r := NewHttp()
//	err := r.SetProxy(proxy)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetProxy(proxy string) error {
	proxyURL, err := url.Parse(proxy)

	e.Proxy = http.ProxyURL(proxyURL)

	return err
}

// SetRedirectFunc sets the redirect function for the HttpRequest object.
// The redirect function takes in a pointer to the current request and an array of prior requests.
// It should return an error if the redirect is invalid.
//
// Example usage:
//
//	func myRedirectFunc(req *http.Request, via []*http.Request) error {
//	    // Custom redirect logic
//	    return nil
//	}
//
//	r := NewHttp()
//	r.SetRedirectFunc(myRedirectFunc)
//
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) SetRedirectFunc(redirect func(*http.Request, []*http.Request) error) {
	e.Redirect = redirect
}

// OnRandomUserAgent sets the flag OnRandomUserAgent to true in the HttpRequest object.
// This flag indicates that a random user agent should be used in the HTTP request.
// By default, the flag is set to false.
//
// Example usage:
// req := &HttpRequest{}
// req.OnRandomUserAgent()
//
// The flag can be checked when making the request to decide whether to use a random user agent or not.
//
//	if req.EntityFeatures.OnRandomUserAgent {
//	    // Use a random user agent
//	} else {
//
//	    // Use a specific user agent
//	}
//
// req.Do()
func (e *HttpRequest) OnRandomUserAgent() {
	e.EntityFeatures.OnRandomUserAgent = true
}

// OnTor enables Tor functionality and sets the proxy for the HttpRequest object.
// It sets the OnTor flag in the EntityFeatures struct to true.
// If setting the proxy encounters an error, it returns that error.
// Otherwise, it returns nil.
// Example usage:
//
//	request := NewHttp()
//	err := request.SetURL("https://check.torproject.org/api/ip")
//	if err != nil {
//	    // Handle the error
//	}
//	err = request.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//	err = request.OnTor()
//	if err != nil {
//	    // Handle the error
//	}
//	response, err := request.Do()
//	if err != nil {
//	    // Handle the error
//	}
//	var result map[string]interface{}
//	err = json.Unmarshal([]byte(response.BRaw), &result)
//	if err != nil {
//	    // Handle the error
//	}
//	ipAddress := fmt.Sprint(result["IP"])
//	fmt.Println(ipAddress)
func (e *HttpRequest) OnTor() error {
	e.EntityFeatures.OnTor = true

	err := e.SetProxy(TORURI)
	if err != nil {
		return err
	}

	return nil
}

// Do send the HTTP request and returns the HTTP response and any error that occurred.
// It creates a new HTTP client with the given request settings.
// It sets the check redirect function, proxy, and TLS configuration.
// It initializes a new HTTP request with the specified method, URL, and data.
// It sets the User-Agent and Content-Type headers.
// It executes the HTTP request and reads the response body.
// It closes the response body and sleeps for the specified duration.
// It returns the HTTP response data as a string and the request information in a structured format.
// If an error occurs during any of the steps, it returns nil for the response and the error.
//
// Example usage:
//
//	r := NewHttp()
//	err := r.SetURL(URL)
//	if err != nil {
//	    // Handle the error
//	}
//
//	err = r.SetMethod("GET")
//	if err != nil {
//	    // Handle the error
//	}
//
//	response, err := r.Do()
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Do something with the response
func (e *HttpRequest) Do() (*HttpResponse, error) {
	client := &http.Client{
		CheckRedirect: e.Redirect,
		Transport: &http.Transport{
			Proxy: e.Proxy,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: e.TlsCertificateCheck,
			},
		},
	}

	e.TimeIn = time.Now()

	r, err := http.NewRequest(e.Method, e.URL, e.Data)

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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err)
		}
	}(response.Body)

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
