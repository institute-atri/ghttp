package ghttp

import (
	"fmt"
	"net"
	"net/url"
)

func ThisIsURL(URL string) error {
	var uri, err = url.ParseRequestURI(URL)

	if err != nil {
		return err
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		return fmt.Errorf("invalid scheme")
	}

	return nil
}

func ThisIsHostValid(URL string) error {
	var uri, err = url.ParseRequestURI(URL)

	if err != nil {
		return err
	}

	_, err = net.LookupHost(uri.Host)

	if err != nil {
		return err
	}

	return nil
}

func GetHost(URL string) (string, error) {
	var uri, err = url.Parse(URL)

	if err != nil {
		return err.Error(), err
	}

	return uri.Host, nil
}
