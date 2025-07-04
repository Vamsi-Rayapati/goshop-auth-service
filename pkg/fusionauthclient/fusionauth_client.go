package fusionauthclient

import (
	"net/http"
	"net/url"
	"time"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/smartbot/auth/pkg/config"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func NewClient() *fusionauth.FusionAuthClient {
	var baseURL, _ = url.Parse(config.Config.FaUrl)
	var client = fusionauth.NewClient(httpClient, baseURL, config.Config.FaApiKey)

	return client

}
