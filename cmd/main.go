package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/weeksontheweb/blockchain-test/cmd/blockchain"
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

			var block blockchain.Block
			var blockchain blockchain.BlockChain

			block, err := block.New([]byte("Torrenting also"))

			if err != nil {
				fmt.Println("Error = ", err)
			}

			fmt.Printf("%d\n", block.Version)
			fmt.Printf("%s\n", block.Payload)

			fmt.Printf("Length of blockchain (1) = %d\n", len(blockchain))

			blockchain, err = blockchain.Add(blockchain, block)

			if err != nil {
				fmt.Println("Error = ", err)
			}

			fmt.Printf("Length of blockchain (2) = %d\n", len(blockchain))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
