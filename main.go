package main

import (
	"routine/builder"
)

func main() {
	app := builder.DefaultAppBuilder()

	app.Run()
}
