package aviasales

import (
	"fmt"
	"net/http"
	"time"
	"travel/internal/pkg/config"
	"travel/pkg/logger"
)

const (
	//VersionAPI version API
	VersionAPI = "v2"
	//ResponseTimeLimit response time limit seconds
	ResponseTimeLimit = 3
)

//LoggingRoundTripper This type implements the http.RoundTripper interface
type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	Logger  *logger.Logger
}

//RoundTrip ...
func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	lrt.Logger.AddDebugLog(fmt.Sprintf("Sending request to %v", req.URL))
	res, e = lrt.Proxied.RoundTrip(req)
	if e != nil {
		lrt.Logger.AddDebugLog(fmt.Sprintf("Error: %v", e))
	} else {
		lrt.Logger.AddDebugLog(fmt.Sprintf("Received %v response", res.Status))
	}
	return
}

//Client ...
type Client struct {
	Cfg    *config.Config
	Client *http.Client
	URL    string
	Logger *logger.Logger
}

//NewClient ...
func NewClient(cfg *config.Config) *Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    ResponseTimeLimit * time.Second,
		DisableCompression: true,
	}
	l := logger.NewLogger(cfg)
	client := &http.Client{Transport: tr}
	if cfg.Debug {
		transportWitLogger := LoggingRoundTripper{tr, l}
		client = &http.Client{Transport: transportWitLogger}
	}
	client.Timeout = ResponseTimeLimit * time.Second

	return &Client{Cfg: cfg,
		Client: client,
		URL: fmt.Sprintf("%s/%s",
			cfg.APIHost, VersionAPI),
		Logger: l}
}
