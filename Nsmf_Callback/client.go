package Nsmf_Callback

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/net/http2"
)

type APIClient struct {
	cfg    *Configuration
	common service

	// API Services
	SmPolicyControlNotifyUpdateApi *SmPolicyControlNotifyUpdateApiService
}

type service struct {
	client *APIClient
}

func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
		cfg.httpClient.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API services
	c.SmPolicyControlNotifyUpdateApi = (*SmPolicyControlNotifyUpdateApiService)(&c.common)

	return c
}
