package middleware

import (
	"encoding/json"
	"net/http"
	"travel/internal/pkg/config"
	"travel/internal/pkg/converter"
	"travel/pkg/aviasales"
	"travel/pkg/logger"
)

//Middleware ...
type Middleware struct {
	Config *config.Config
	Logger *logger.Logger
}

//NewMiddleware ...
func NewMiddleware(cfg *config.Config) *Middleware {
	l := logger.NewLogger(cfg)
	return &Middleware{Config: cfg, Logger: l}
}

func (m *Middleware) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	client := aviasales.NewClient(m.Config)

	result, err := client.GetPlaces(r.URL.Query())
	if len(r.URL.RawQuery) == 0 {
		m.Logger.AddDebugLog("Not set GET parameters")
		m.Logger.AddDebugLog(r.URL.RequestURI())
		http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err != nil {
		m.Logger.AddDebugLog(err.Error())
		http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(converter.Convert(result))
	if err != nil {
		m.Logger.AddDebugLog(err.Error())
	}
}
