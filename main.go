package main

import (
	"flag"

	"github.com/ramabmtr/go-barebone/app"
	"github.com/ramabmtr/go-barebone/app/config"
)

// @title BareBone folder structure
// @version v1.0.0
// @description You can use this as base app structure for building another service

// @contact.name Rama Bramantara
// @contact.url https://www.linkedin.com/in/ramabmtr/
// @contact.email ramabmtr@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @basePath /
// @query.collection.format multi

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./conf.yaml", "path to config file")
	flag.Parse()

	config.InitConf(configPath)
	app.Run()
}
