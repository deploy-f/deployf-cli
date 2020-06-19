package main

import (
	"cli/client"
	"cli/cmd"
	"log"
	"net/url"
	"os"
	"sort"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/urfave/cli/v2"
)

var transport *httptransport.Runtime

func main() {
	app := &cli.App{
		Description: "command line tool for access to deploy-f api",
		Version:     Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token, t",
				Usage:    "Personal Access Token for access to api",
				EnvVars:  []string{"DF_API_TOKEN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "host",
				Usage:   "Api host",
				Value:   "https://api.deploy-f.com",
				EnvVars: []string{"DF_HOST"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "set-image",
				Aliases: []string{"si"},
				Usage:   "Set an image to application",
				Action:  cmd.SetImage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "app-id, a",
						Usage:    "Application id",
						EnvVars:  []string{"DF_APP_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "image, i",
						Usage:    "Docker image",
						Required: true,
					},
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start or restart application",
				Action:  cmd.Start,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "app-id, a",
						Usage:    "Application id",
						EnvVars:  []string{"DF_APP_ID"},
						Required: true,
					},
				},
			},
			{
				Name:    "stop",
				Aliases: []string{"st"},
				Usage:   "Stop application",
				Action:  cmd.Stop,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "app-id, a",
						Usage:    "Application id",
						EnvVars:  []string{"DF_APP_ID"},
						Required: true,
					},
				},
			},
		},
		Before: initApp,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initApp(c *cli.Context) error {
	uri, err := url.Parse(c.String("host"))
	if err != nil {
		panic(err)
	}
	transport = httptransport.New(uri.Host, "", []string{uri.Scheme})
	transport.Producers["application/*+json"] = runtime.JSONProducer()
	cmd.Auth = httptransport.BearerToken(c.String("token"))
	cmd.Api = client.New(transport, strfmt.Default)
	return nil
}
