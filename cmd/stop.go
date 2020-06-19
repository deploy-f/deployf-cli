package cmd

import (
	"cli/client/application"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// Stop the application
func Stop(c *cli.Context) error {
	id, _ := strconv.Atoi(c.String("app-id"))

	params := application.NewStopParams()
	params.ID = int32(id)

	_, err := Api.Application.Stop(params, Auth)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("OK")

	return nil
}
