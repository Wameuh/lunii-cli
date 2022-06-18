package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

type Story struct {
	uuid       uuid.UUID
	folderName string
}

func main() {
	pack, _ := ReadPack("reference/pack.zip")

	imageAssets := getImageAssetListFromPack(*pack)
	soundAssets := getSoundAssetListFromPack(*pack)
	listNodeIndex := getListNodeIndex(&pack.ListNodes)

	// riBinary := cipherFirstBlockCommonKey(GenerateBinaryFromAssetIndex(&imageAssets))
	// siBinary := cipherFirstBlockCommonKey(GenerateBinaryFromAssetIndex(&soundAssets))
	// btBinary := cipherBlockSpecificKey(riBinary[:64])
	// liBinary := cipherFirstBlockCommonKey(GenerateBinaryFromListNodeIndex(&listNodeIndex, &pack.StageNodes))
	niBinary := generateNiBinary(pack, &pack.StageNodes, &listNodeIndex, &imageAssets, &soundAssets)

	logg("Ni Binary", niBinary)

}

func startCli() {
	app := &cli.App{
		Name:  "lunii cli",
		Usage: "lunii cli",
		Commands: []*cli.Command{
			{
				Name:    "device:infos",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Action: func(c *cli.Context) error {
					fmt.Println(GetDeviceInfos())
					return nil
				},
			},
			{
				Name:    "packs:list",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Action: func(c *cli.Context) error {
					fmt.Println(readGlobalIndexFile())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
