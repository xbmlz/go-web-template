package main

//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./cmd/server/main.go -d ../../ -o ../../docs

import (
	"flag"

	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
	"github.com/xbmlz/go-web-template/internal/config"
	"github.com/xbmlz/go-web-template/internal/database"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/internal/router"
	"github.com/xbmlz/go-web-template/internal/server"
	"github.com/xbmlz/go-web-template/internal/validator"
)

// @title go-web-template API
// @version 1.0
// @description This is a sample server for go-web-template
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @basePath /api
func main() {
	var confPath string

	flag.StringVar(&confPath, "c", "conf/config.yaml", "path to the configuration file, e.g. -c config.yaml")
	flag.Parse()

	// initialize the configuration
	c := config.MustInit(confPath)

	// initialize the logger
	logger.Init(c.Server.IsDev())

	// initialize the database connection
	db := database.MustInit(&c.Database)

	// migrate the database
	db.AutoMigrate(model.AllModels()...)

	// setup query
	query.SetDefault(db)

	// nitialize the validator
	validator.Init()

	// initialize the server
	r := router.Init(c)

	// start the server
	server.Run(r, &c.Server)
}
