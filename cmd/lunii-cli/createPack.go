package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	studiopackbuilder "github.com/olup/lunii-cli/pkg/pack-builder"
	"github.com/urfave/cli/v2"
)

func createPack(c *cli.Context) error {
	inputPath := c.String("directory")
	outputPath := c.String("output")

	if inputPath == "" {
		color.Red("Input directory is missing")
		return errors.New("Wrong argument")
	}

	if outputPath == "" || outputPath[len(outputPath)-4:] != ".zip" {
		color.Red("Output path missing or is not a Zip")
		return errors.New("Wrong argument")
	}

	fmt.Println("Generating a new studio pack from " + inputPath)
	_, err := studiopackbuilder.CreateStudioPack(inputPath, outputPath)
	if err != nil {
		color.Red("Could not create the studio pack")
		return err
	}
	color.Green("Studio pack created: " + outputPath)

	return err
}
