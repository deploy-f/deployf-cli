package main

import (
	"cli/client"
	"cli/client/home"
	"fmt"
	"log"
	"os"
	"sort"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/urfave/cli/v2"
)

func main() {

	transport := httptransport.New("localhost:5002", "", nil)
	apiclient := client.New(transport, strfmt.Default)
	r, e := apiclient.Home.GetHealth(home.NewGetHealthParams())

	if e != nil {
		fmt.Println("error:", e)
		return
	}

	fmt.Println(r)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang, l",
				Value: "english",
				Usage: "Language for the greeting",
			},
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action:  complete,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func complete(c *cli.Context) error {
	fmt.Printf("complete\n")
	return nil
}
