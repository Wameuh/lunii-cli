package main

import (
	"github.com/fatih/color"
	"github.com/olup/lunii-cli/pkg/lunii"
	"github.com/urfave/cli/v2"
)

func ImportPack(c *cli.Context) error {
	path := c.Args().Get(0)

	device, err := lunii.GetDevice()
	if err != nil {
		return err
	}
	studioPack, err := lunii.ReadStudioPack(path)
	if err != nil {
		return err
	}

	err = device.AddStudioPack(studioPack)
	if err != nil {
		return err
	}

	color.Green("Updated pack " + studioPack.Uuid.String())
	return nil
}
