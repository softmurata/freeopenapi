package Npcf_Callback

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/net/http2"

)

type APIClient struct {
	cfg *Configuration
	common service


	// API Services
	UdrDataModificationNotificationApi *UdrDataModificationNotificationApiService
	N1NotificationApi *N1NotificationApiService

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
	c.UdrDataModificationNotificationApi = (*UdrDataModificationNotificationApiService)(&c.common)
	c.N1NotificationApi = (*N1NotificationApiService)(&c.common)
	return c
}