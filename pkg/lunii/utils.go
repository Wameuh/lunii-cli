package lunii

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/olup/lunii-cli/pkg/bmp4"
	"golang.org/x/image/draw"

	_ "github.com/qiniu/audio/mp3"
	_ "github.com/qiniu/audio/ogg"
	_ "github.com/qiniu/audio/wav"
	"github.com/tosone/minimp3"
	"github.com/viert/go-lame"
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

func ImageToBmp4(file io.Reader) ([]byte, error) {
	source, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// converting image to grayscale
	grayscale := image.NewGray(image.Rect(0, 0, 320, 240))
	draw.NearestNeighbor.Scale(grayscale, grayscale.Bounds(), source, source.Bounds(), draw.Over, nil)

	return bmp4.GetBitmap(grayscale), nil
}

func AudioToMp3(file io.ReadSeeker) ([]byte, error) {

	fileByte, _ := ioutil.ReadAll(file)
	dec, audio, _ := minimp3.DecodeFull(fileByte)

	if dec.Channels == 1 && dec.SampleRate == 44100 {
		fmt.Println("Proper mp3 format, no need to rencode")
		return fileByte, nil
	}

	output := new(bytes.Buffer)
	enc := lame.NewEncoder(output)
	defer enc.Close()
	fmt.Println("one")

	enc.SetMode(3)
	fmt.Println("two")

	enc.SetInSamplerate(44100)
	enc.Write(audio)

	return output.Bytes(), nil

}
