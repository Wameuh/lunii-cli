package lunii

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

func (device *Device) AddStudioPack(studioPack *StudioPack) error {
	// 1. Get path on devide
	rootPath := device.MountPoint
	contentPath := filepath.Join(rootPath, ".content", studioPack.Ref)

	fmt.Println("Generating Binaries...")

	stageNodeIndex := &studioPack.StageNodes
	// generate list node index
	listNodeIndex := GetListNodeIndex(&studioPack.ListNodes)

	// create image index - with exporter & lookup
	imageIndex := GetImageAssetListFromPack(studioPack)
	// create sound index - with exporter & lookup
	audioIndex := GetAudioAssetListFromPack(studioPack)

	// prepare ni (index of stage nodes)
	niBinary := GenerateNiBinary(studioPack, stageNodeIndex, listNodeIndex, imageIndex, audioIndex)

	// prepare li (index of action nodes)
	liBinary := generateLiBinary(listNodeIndex, stageNodeIndex)
	liBinaryCiphered := cipherFirstBlockCommonKey(liBinary)

	// prepare ri (index of image assets)
	riBinary := GenerateBinaryFromAssetIndex(imageIndex)
	riBinaryCiphered := cipherFirstBlockCommonKey(riBinary)

	// prepare si (index of sound assets)
	siBinary := GenerateBinaryFromAssetIndex(audioIndex)
	siBinaryCiphered := cipherFirstBlockCommonKey(siBinary)

	// create boot file
	btBinary := generateBtBinary(riBinaryCiphered)

	// prepare pack
	tempPath := filepath.Join(os.TempDir(), "packs", studioPack.Ref)
	err := os.MkdirAll(tempPath, 0700)
	if err != nil {
		return err
	}

	fmt.Println("Preparing asset in " + tempPath)

	err = os.WriteFile(filepath.Join(tempPath, "ni"), niBinary, 0777)
	err = os.WriteFile(filepath.Join(tempPath, "li"), liBinaryCiphered, 0777)
	err = os.WriteFile(filepath.Join(tempPath, "ri"), riBinaryCiphered, 0777)
	err = os.WriteFile(filepath.Join(tempPath, "si"), siBinaryCiphered, 0777)
	err = os.WriteFile(filepath.Join(tempPath, "bt"), btBinary, 0777)
	err = os.MkdirAll(filepath.Join(tempPath, "sf"), 0700)
	err = os.MkdirAll(filepath.Join(tempPath, "rf"), 0700)
	if err != nil {
		return err
	}

	// prepare zip reader
	reader, err := zip.OpenReader(studioPack.OriginalPath)
	if err != nil {
		return errors.New("Zip package could not be opened")
	}
	defer reader.Close()

	// copy images in rf
	deviceImageDirectory := filepath.Join(tempPath, "rf", "000")
	os.MkdirAll(deviceImageDirectory, 0700)

	for _, image := range *imageIndex {
		file, err := reader.Open("assets/" + image.SourceName)
		if err != nil {
			return err
		}
		bmpFile, err := ImageToBmp4(file)
		if err != nil {
			return err
		}
		cypheredBmp := cipherFirstBlockCommonKey(bmpFile)
		err = os.WriteFile(filepath.Join(deviceImageDirectory, image.DestinationName), cypheredBmp, 0777)
		if err != nil {
			return err
		}
	}

	// copy audios in sf
	deviceAudioDirectory := filepath.Join(tempPath, "sf", "000")
	os.MkdirAll(deviceAudioDirectory, 0700)

	for _, audio := range *audioIndex {
		file, err := reader.Open("assets/" + audio.SourceName)
		if err != nil {
			return err
		}
		fmt.Println(audio.SourceName)
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		mp3, err := AudioToMp3(bytes.NewReader(fileBytes))
		if err != nil {
			return err
		}

		cypheredFile := cipherFirstBlockCommonKey(mp3)

		err = os.WriteFile(filepath.Join(deviceAudioDirectory, audio.DestinationName), cypheredFile, 0777)
		if err != nil {
			return err
		}
	}

	// copy temp to lunii
	fmt.Println("Copying directory to the device...")
	cp.Copy(tempPath, contentPath)

	fmt.Println("Adding pack to root index...")
	// // update .pi root file with uuid
	err = device.AddPackToIndex(studioPack.Uuid)
	if err != nil {
		return err
	}

	fmt.Println("Cleaning...")
	// err = os.RemoveAll(tempPath)
	// if err != nil {
	// 	return err
	// }

	return nil
}
