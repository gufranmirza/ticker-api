package main

import (
	"github.com/gufranmirza/ticker-api/cmd"
)

// @title Ticker API Documentation
// @version 2.0
// @description Ticker API Documentation

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /ticker-api/v1
// @query.collection.format multi

func main() {
	cmd.Execute()
}
