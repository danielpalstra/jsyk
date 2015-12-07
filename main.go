package main

import (
	"os"
	//"encoding/json"
	// "net/http"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "jsyk"
	app.Version = "0.1.0"
	app.Usage = "jsyk [command] [arguments]"
	app.Author = "Daniel Palstra"
	app.Email = "daniel@danielpalstra.com"
	app.Commands = Commands
	app.Run(os.Args)

}
