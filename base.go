package ghttp

import "log"

// GET sends a GET request to the specified URL and returns the HTTP response.
// If an error occurs during the request process, it will be logged and a nil response will be returned.
func GET(URL string) *HttpResponse {
	r := NewHttp()
	err := r.SetURL(URL)
	if err != nil {
		return nil
	}

	err = r.SetMethod("GET")
	if err != nil {
		return nil
	}

	response, err := r.Do()
	if err != nil {
		log.Fatal(err)
	}

	return response
}

// POST sends a POST request to the specified URL with the provided data and returns the HTTP response.
// If an error occurs during the request process, it will be logged and a nil response will be returned.
func POST(URL string, data string) *HttpResponse {
	var err error

	r := NewHttp()

	err = r.SetURL(URL)
	if err != nil {
		log.Fatal(err)
	}

	err = r.SetMethod("POST")
	if err != nil {
		log.Fatal(err)
	}

	r.SetData(data)
	r.SetContentType("application/x-www-form-urlencoded")

	response, err := r.Do()
	if err != nil {
		log.Fatal(err)
	}

	return response
}

// PUT sends a PUT request to the specified URL and returns the HTTP response.
// If an error occurs during the request process, it will be logged and a nil response will be returned.
func PUT(URL string) *HttpResponse {
	r := NewHttp()

	err := r.SetURL(URL)
	if err != nil {
		return nil
	}

	err = r.SetMethod("PUT")
	if err != nil {
		return nil
	}

	response, err := r.Do()
	if err != nil {
		log.Fatal(err)
	}

	return response
}

// DELETE sends a DELETE request to the specified URL and returns the HTTP response.
// If an error occurs during the request process, it will be logged and a nil response will be returned.
func DELETE(URL string) *HttpResponse {
	r := NewHttp()

	err := r.SetURL(URL)
	if err != nil {
		return nil
	}

	err = r.SetMethod("DELETE")
	if err != nil {
		return nil
	}

	response, err := r.Do()

	if err != nil {
		log.Fatal(err)
	}

	return response
}
