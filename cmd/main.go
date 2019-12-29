package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	fmt.Println("Entry Point")
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "create-block",
				Value: "",
				Usage: "create a block",
			},
		},
		Action: func(c *cli.Context) error {
			//			if c.NArg() > 0 {
			//				name = c.Args().Get(0)
			//			}
			if c.String("create-block") == "y" {
				fmt.Println("creating a block")
			} else {
				fmt.Println("not creating a block")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
