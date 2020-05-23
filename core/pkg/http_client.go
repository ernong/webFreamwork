package pkg

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type HttpClientConfig struct {
	TimeOut             time.Duration
	KeepAlive           time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

func LoadHttpClientConfig() *HttpClientConfig {
	cfg := &HttpClientConfig{
		TimeOut:             30,
		KeepAlive:           30,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     30,
	}
	fmt.Printf("LoadHttpClientConfig:%#v\n", cfg)
	return cfg
}

func (cfg *HttpClientConfig) InitHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.TimeOut * time.Second,
				KeepAlive: cfg.KeepAlive * time.Second,
			}).DialContext,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			IdleConnTimeout:     cfg.IdleConnTimeout * time.Second,
		},
		Timeout: cfg.TimeOut * time.Second,
	}
	return client
}
