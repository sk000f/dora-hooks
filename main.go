package main

import (
	"github.com/sk000f/metrix/pkg/gitlab"
)

func main() {
	app := gitlab.App{}
	app.Init()
	app.Run()
}
