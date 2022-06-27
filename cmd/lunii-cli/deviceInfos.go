package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/olup/lunii-cli/pkg/lunii"
	"github.com/urfave/cli/v2"
)

func DisplayDeviceInfos(c *cli.Context) error {
	device, err := lunii.GetDevice()
	if err != nil {
		color.Red("No device found")
		color.Yellow("Check if the device is plugged and mounted")
		color.Yellow("Note: Devices version 1 are not supported by the CLI")
		return err
	}

	fmt.Println("Device infos")
	fmt.Println("------------")
	table := uitable.New()
	table.AddRow("S/N:", device.SerialNumber)
	table.AddRow("Mountpoint:", device.MountPoint)
	table.AddRow("Version:", fmt.Sprint(device.FirmwareVersionMajor)+"."+fmt.Sprint(device.FirmwareVersionMinor))
	fmt.Println(table)
	color.Green("This device is supported by this CLI")
	return nil
}
