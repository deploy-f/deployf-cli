package main

import (
	"cli/client"
	"cli/client/application"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/urfave/cli/v2"
)

var (
	transport *httptransport.Runtime
	api       *client.Deplyf
	auth      runtime.ClientAuthInfoWriter
)

func main() {
	app := &cli.App{
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
				Value:   "api.deploy-f.com",
				EnvVars: []string{"DF_HOST"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "set-image",
				Aliases: []string{"si"},
				Usage:   "Set an image to application",
				Action:  setImage,
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
				Action:  start,
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
				Action:  stop,
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

func setImage(c *cli.Context) error {
	return fmt.Errorf("Not implemented")
}

func start(c *cli.Context) error {
	id, _ := strconv.Atoi(c.String("app-id"))

	params := application.NewStartParams()
	params.ID = int32(id)

	_, err := api.Application.Start(params, auth)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("OK")

	return nil
}

func stop(c *cli.Context) error {
	id, _ := strconv.Atoi(c.String("app-id"))

	params := application.NewStopParams()
	params.ID = int32(id)

	_, err := api.Application.Stop(params, auth)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("OK")

	return nil
}

func initApp(c *cli.Context) error {
	transport = httptransport.New(c.String("host"), "", nil)
	auth = httptransport.BearerToken(c.String("token"))
	api = client.New(transport, strfmt.Default)
	return nil
}
