package utils

import (
	"errors"
	"net/url"
)

type ProxyParams struct {
	Server string
	Port   string
	Secret string
}

func GetURLParams(proxyUrl string) (ProxyParams, error) {
	parsedURL, _ := url.Parse(proxyUrl)

	queryParams := parsedURL.Query()

	server := queryParams.Get("server")
	port := queryParams.Get("port")
	secret := queryParams.Get("secret")

	if server == "" || port == "" || secret == "" {
		return ProxyParams{}, errors.New("server, port and secret are required")
	}

	return ProxyParams{server, port, secret}, nil
}
