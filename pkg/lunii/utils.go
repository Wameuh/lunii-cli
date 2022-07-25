package lunii

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
	"github.com/olup/lunii-cli/pkg/bmp4"
	"github.com/tosone/minimp3"
	"golang.org/x/image/draw"

	"github.com/qiniu/audio"
	_ "github.com/qiniu/audio/ogg"
	_ "github.com/qiniu/audio/wav"
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

func AudioToMp3(fileBytes []byte) ([]byte, error) {
	var audioBytes []byte

	start := time.Now()

	// maybe mp3 ?
	data, mp3Audio, err := minimp3.DecodeFull(fileBytes)

	fmt.Println("Ausio opened in ", time.Since(start))

	if err != nil {
		// if ogg or wav
		source, _, err := audio.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			return nil, err
		}
		audioBytes, err = ioutil.ReadAll(source)
		if err != nil {
			return nil, err
		}
	}

	// if data.Channels == 1 && data.SampleRate == 44100 {
	// 	fmt.Println(len(fileBytes))
	// 	fmt.Println("No conversion needed")
	// 	return fileBytes, nil
	// }

	start = time.Now()

	audioBytes = mp3Audio

	output := new(bytes.Buffer)
	enc := lame.NewEncoder(output)
	defer enc.Close()

	if data.Channels == 1 {
		enc.SetNumChannels(1)
	}

	enc.SetVBR(lame.VBRDefault)
	enc.SetVBRQuality(4)
	enc.SetQuality(4)
	enc.SetMode(lame.MpegMono)
	enc.SetWriteID3TagAutomatic(false)
	enc.Write(audioBytes)
	enc.Flush()
	fmt.Println("Audio converted in ", time.Since(start))

	return output.Bytes(), nil

}

func insert(array []uuid.UUID, element uuid.UUID, i int) []uuid.UUID {
	return append(array[:i], append([]uuid.UUID{element}, array[i:]...)...)
}
