package utils

import (
	"errors"
	"fmt"
	"net/url"
)

type ProxyParams struct {
	Server string
	Port   string
	Secret string
}

func GetProxyUrlParams(proxyUrl string) (ProxyParams, error) {
	parsedURL, err := url.Parse(proxyUrl)
	if err != nil {
		return ProxyParams{}, err
	}

	if parsedURL.Scheme != "https" || parsedURL.Host != "t.me" || parsedURL.Path != "/proxy" {
		return ProxyParams{}, fmt.Errorf("URL must have a scheme and host")
	}

	queryParams := parsedURL.Query()

	server := queryParams.Get("server")
	port := queryParams.Get("port")
	secret := queryParams.Get("secret")

	if server == "" || port == "" || secret == "" {
		return ProxyParams{}, errors.New("server, port and secret are required")
	}

	return ProxyParams{server, port, secret}, nil
}
