package lunii

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/olup/lunii-cli/pkg/bmp4"
	"golang.org/x/image/draw"
)

func logg(datas ...any) {
	for _, data := range datas {
		fmt.Println(data)
	}
	fmt.Println()
}

func boolToShort(boolean bool) int16 {
	if boolean {
		return 1
	} else {
		return 0
	}
}

func ConvertImageFromPath(path string) ([]byte, error) {
	imgBin, _ := os.Open(path)
	defer imgBin.Close()

	source, _, err := image.Decode(imgBin)
	if err != nil {
		return nil, err
	}

	// converting image to grayscale
	grayscale := image.NewGray(image.Rect(0, 0, 536, 354))
	draw.NearestNeighbor.Scale(grayscale, grayscale.Bounds(), source, source.Bounds(), draw.Over, nil)

	return bmp4.GetBitmap(grayscale), nil
}
