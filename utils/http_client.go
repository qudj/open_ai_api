package utils

import (
	"net"
	"net/http"
	"time"
)

func DefaultHttpClient() *http.Client {
	tr := &http.Transport{
		ResponseHeaderTimeout: 10 * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   10 * time.Second,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConnsPerHost:   50,
		ExpectContinueTimeout: 1 * time.Second,
		MaxConnsPerHost:       50,
	}
	return &http.Client{
		Transport: tr,
	}
}
