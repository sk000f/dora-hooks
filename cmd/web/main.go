package main

import (
	"github.com/sk000f/metrix/pkg/gitlab"
	"github.com/sk000f/metrix/pkg/metrix"
	"github.com/sk000f/metrix/pkg/sonar"
)

func main() {
	app := metrix.App{}
	app.Init()
	app.AddRoutes(gitlab.InitRoutes())
	app.AddRoutes(sonar.InitRoutes())
	app.Run()
}
