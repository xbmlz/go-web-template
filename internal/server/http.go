package server

import (
	"fmt"
	"net/http"

	"github.com/xbmlz/go-web-template/internal/logger"
)

type Config struct {
	Mode     string `json:"mode"` // debug or release
	Host     string `json:"host"`
	Port     string `json:"port"`
	BasePath string `json:"base_path"`
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *Config) IsDev() bool {
	return c.Mode == "debug"
}

func Run(h http.Handler, c *Config) {
	// build routes here
	srv := &http.Server{
		Addr:    c.Addr(),
		Handler: h,
	}

	logger.Infof("Starting server on %s", c.Addr())
	logger.Infof("Swagger UI available at http://%s/swagger/index.html", c.Addr())
	logger.Info("Press CTRL+C to stop")
	if err := srv.ListenAndServe(); err != nil {
		logger.Panicf("ListenAndServe: %v", err)
	}
}
