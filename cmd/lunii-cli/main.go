package main

import (
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
						Aliases: []string{"i"},
						Usage:   "Get general informations",
						Action:  DisplayDeviceInfos,
					},
				}},
			{
				Name:    "pack",
				Aliases: []string{"p"},
				Usage:   "Story packs",
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "List installed packs on the device",
						Action:  DisplayInstalledPacks,
					},
					{
						Name:    "remove",
						Aliases: []string{"rm"},
						Usage:   "Remove one pack from device",
						Action:  RemovePack,
					},
					{
						Name:    "import",
						Aliases: []string{"i"},
						Usage:   "Import a studio pack",
						Action:  ImportPack,
					},
					{
						Name:    "create",
						Aliases: []string{"c"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "directory",
								Usage: "The structured directory of the story",
							},
							&cli.StringFlag{
								Name:  "output",
								Usage: "Output archive for the pack",
							},
						},
						Usage:  "Create a studio pack from a structured directory",
						Action: createPack,
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
