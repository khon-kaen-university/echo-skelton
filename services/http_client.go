package services

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/valyala/fastjson"
)

// HTTPClient custom http clent
var HTTPClient *http.Client
var jsonParser fastjson.Parser

var netTransport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
}

// CreateHTTPClient custom http client
func CreateHTTPClient() {
	HTTPClient = &http.Client{
		Timeout:   8 * time.Second,
		Transport: netTransport,
	}
	return
}