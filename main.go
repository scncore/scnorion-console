package main

import (
	"log"
	"os"

	"github.com/scncore/scnorion-console/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "scnorion-console",
		Commands:  getCommands(),
		Usage:     "The scnorion console allows and organization to manage its endpoints from a Web User Interface",
		Authors:   []*cli.Author{{Name: "Miguel Angel Alvarez Cabrerizo", Email: "mcabrerizo@scnorion.eu"}},
		Copyright: "2024-2025 - Miguel Angel Alvarez Cabrerizo <https://github.com/scnorion>",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() []*cli.Command {
	return []*cli.Command{
		commands.StartConsole(),
		commands.StopConsole(),
	}
}
