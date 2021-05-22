package main

import (
	"embed"
	"go-embed-sample/filebox"
	"go-embed-sample/server"
)

//go:embed statics
var embedFiles embed.FS

func main() {
	filebox.AddFiles(embedFiles)

	server.EchoStart(8080)
}
