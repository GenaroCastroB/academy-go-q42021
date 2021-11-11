package main

import (
	"golangBootcamp/m/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
