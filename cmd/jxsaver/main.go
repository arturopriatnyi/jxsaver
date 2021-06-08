// Package main is JXSaver entry point.
package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "JXSaver",
		Usage: "CLI tool for validating and saving JSON and XML as files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "json",
				Usage: "Parses, validates and saves JSON",
			},
			&cli.StringFlag{
				Name:  "xml",
				Usage: "Parses, validates and saves XML",
			},
		},
		Action: func(c *cli.Context) error {
			if len(c.FlagNames()) != 1 {
				return errors.New("invalid number of flags provided")
			}

			format := c.FlagNames()[0]
			data := c.String(format)
			log.Printf("Format: %s, Data: %s", format, data)

			// Logic goes here

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
