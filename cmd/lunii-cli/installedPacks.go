package main

import (
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/olup/lunii-cli/pkg/lunii"
	"github.com/urfave/cli/v2"
)

func DisplayInstalledPacks(c *cli.Context) error {
	device, err := lunii.GetDevice()
	if err != nil {
		return err
	}

	packs, err := device.ReadGlobalIndexFile()
	if err != nil {
		return err
	}

	table := uitable.New()
	table.AddRow("REF", "UUID", "TITLE")
	for _, pack := range packs {
		table.AddRow(pack.FolderName, pack.Uuid.String(), pack.Title)
	}
	fmt.Println(table)
	return nil
}
