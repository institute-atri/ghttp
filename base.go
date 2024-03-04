package main

import "log"

func GET(URL string) *HttpResponse {
	var r = NewHttp()
	r.SetURL(URL)
	r.SetMethod("GET")

	var response, err = r.Do()

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func POST(URL string, data string) *HttpResponse {
	var err error

	var r = NewHttp()

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

func PUT(URL string) *HttpResponse {
	var r = NewHttp()
	r.SetURL(URL)
	r.SetMethod("PUT")

	var response, err = r.Do()

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func DELETE(URL string) *HttpResponse {
	var r = NewHttp()
	r.SetURL(URL)
	r.SetMethod("DELETE")

	var response, err = r.Do()

	if err != nil {
		log.Fatal(err)
	}

	return response
}
