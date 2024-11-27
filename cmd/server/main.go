package main

import (
	"flag"

	"github.com/xbmlz/go-web-template/api/query"
	"github.com/xbmlz/go-web-template/internal/config"
	"github.com/xbmlz/go-web-template/internal/database"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/internal/router"
	"github.com/xbmlz/go-web-template/internal/server"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "config.yaml", "path to the configuration file, e.g. -c config.yaml")
}

func main() {

	// 1. initialize the configuration
	c := config.MustInit(configFile)

	// 2. initialize the logger
	logger.Init(c)

	// 3. initialize the database connection
	db := database.MustInit(&c.Database)

	// 4. setup query
	query.Use(db)

	// 5. initialize the server
	r := router.Init(c)

	// 6. start the server
	server.Run(r, &c.Server)
}
