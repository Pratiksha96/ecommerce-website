package main

import (
	"ecommerce-website/app/server"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "e-commerce website"
	app.Usage = "Website to shop products"
	app.Commands = []cli.Command{
		{
			Name: "start-server",
			Action: func(c *cli.Context) error {
				server.StartServer()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
