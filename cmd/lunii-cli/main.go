package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	//"image/draw"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	lunii "github.com/olup/lunii-cli/pkg/lunii"
	"github.com/urfave/cli/v2"
)

func main() {

	// imgBin, _ := os.Open("./test/test.jpeg")
	// defer imgBin.Close()

	// source, _, err := image.Decode(imgBin)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // converting image to grayscale
	// grayscale := image.NewGray(image.Rect(0, 0, 536, 354))
	// draw.NearestNeighbor.Scale(grayscale, grayscale.Bounds(), source, source.Bounds(), draw.Over, nil)

	// by := bmp.GetBitmap(grayscale)

	// os.WriteFile("./test/compressed.bmp", by, 0777)

	startCli()
}

func startCli() {
	app := &cli.App{
		Name:  "lunii cli",
		Usage: "lunii cli",
		Commands: []*cli.Command{
			{
				Name:    "device",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Subcommands: []*cli.Command{
					{
						Name:    "infos",
						Aliases: []string{"d"},
						Usage:   "Read global config",
						Action: func(c *cli.Context) error {
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
						},
					},
				}},
			{
				Name:    "pack",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"d"},
						Usage:   "List installed packs on the device",
						Action: func(c *cli.Context) error {
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
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"d"},
						Usage:   "Remove one pack from device",
						Action: func(c *cli.Context) error {
							table := uitable.New()
							table.AddRow("UUID", "TITLE")
							table.AddRow("00000-00000-00000000-0000", "The title")
							fmt.Println(table)
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
