package server

import (
	"fmt"
	"net/http"
)

type Config struct {
	Mode string `json:"mode"` // debug or release
	Host string `json:"host"`
	Port string `json:"port"`
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

	fmt.Printf("Starting server on %s\n", c.Addr())
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
