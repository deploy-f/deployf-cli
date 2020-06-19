package cmd

import (
	"cli/client/application"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// Start start the app
func Start(c *cli.Context) error {
	id, _ := strconv.Atoi(c.String("app-id"))

	params := application.NewStartParams()
	params.ID = int32(id)

	_, err := Api.Application.Start(params, Auth)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("OK")

	return nil
}
