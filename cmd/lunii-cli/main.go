package main

import (
	_ "image/jpeg"
	_ "image/png"

	//"image/draw"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
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
				Usage:   "About your lunii device",
				Subcommands: []*cli.Command{
					{
						Name:    "infos",
						Aliases: []string{"d"},
						Usage:   "Get general informations",
						Action:  DisplayDeviceInfos,
					},
				}},
			{
				Name:    "pack",
				Aliases: []string{"d"},
				Usage:   "Story packs",
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"d"},
						Usage:   "List installed packs on the device",
						Action:  DisplayInstalledPacks,
					},
					{
						Name:    "remove",
						Aliases: []string{"d"},
						Usage:   "Remove one pack from device",
						Action:  RemovePack,
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
