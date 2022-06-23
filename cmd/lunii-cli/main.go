package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	//"image/draw"
	"log"
	"os"

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
				Name:    "device-infos",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Action: func(c *cli.Context) error {
					fmt.Println(lunii.GetDevice())
					return nil
				},
			},
			{
				Name:    "list-packs",
				Aliases: []string{"d"},
				Usage:   "Read global config",
				Action: func(c *cli.Context) error {
					fmt.Println(lunii.ReadGlobalIndexFile())
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
