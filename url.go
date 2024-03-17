package ghttp

import (
	"fmt"
	"net"
	"net/url"
)

// ThisIsURL validates if a given URL has a valid scheme of either "http" or "https"
// and returns an error if not.
// It takes a string parameter, URL, representing the URL to be validated.
// Returns an error if the scheme is not valid, otherwise returns nil.
func ThisIsURL(URL string) error {
	uri, err := url.ParseRequestURI(URL)
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

// ThisIsHostValid validates if a given URL has a valid host by performing a DNS lookup on the host.
// It takes a string parameter, URL, representing the URL to be validated.
// Returns an error if the host is not valid, otherwise returns nil.
func ThisIsHostValid(URL string) error {
	uri, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	_, err = net.LookupHost(uri.Host)
	if err != nil {
		return err
	}

	return nil
}

// GetHost extracts the host from a given URL and returns it as a string.
// It takes a string parameter, URL, representing the URL from which the host will be extracted.
// Returns the host as a string and nil error if the URL is successfully parsed.
// Returns an error message and the error itself if the URL parsing fails.
func GetHost(URL string) (string, error) {
	uri, err := url.Parse(URL)
	if err != nil {
		return err.Error(), err
	}

	return uri.Host, nil
}
