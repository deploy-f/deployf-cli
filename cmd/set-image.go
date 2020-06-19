package cmd

import (
	"cli/client/application"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// SetImage an image to the application
func SetImage(c *cli.Context) error {
	id, _ := strconv.Atoi(c.String("app-id"))
	image := c.String("image")

	params := application.NewGetCustomImageParams()
	params.ID = int32(id)

	existImage, err := Api.Application.GetCustomImage(params, Auth)

	if err != nil {
		return fmt.Errorf("get image error: %v", err)
	}

	newDto := *existImage.Payload
	newDto.Image = &image

	setParams := application.NewSetCustomImageParams()
	setParams.ID = int32(id)
	setParams.Data = &newDto

	_, err = Api.Application.SetCustomImage(setParams, Auth)

	if err != nil {
		return fmt.Errorf("set image error: %v", err)
	}

	fmt.Println("OK")
	return nil
}
