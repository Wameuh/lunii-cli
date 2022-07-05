package lunii

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	cp "github.com/otiai10/copy"
	yaml "gopkg.in/yaml.v3"
)

var wg sync.WaitGroup

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
	fmt.Println("Converting images ...")

	deviceImageDirectory := filepath.Join(tempPath, "rf", "000")
	os.MkdirAll(deviceImageDirectory, 0700)

	for i, image := range *imageIndex {
		wg.Add(1)
		go convertAndWriteImage(*reader, deviceImageDirectory, image, i)
	}

	wg.Wait()

	// copy audios in sf
	fmt.Println("Converting audios ...")

	deviceAudioDirectory := filepath.Join(tempPath, "sf", "000")
	os.MkdirAll(deviceAudioDirectory, 0700)

	for i, audio := range *audioIndex {
		wg.Add(1)
		go convertAndWriteAudio(*reader, deviceAudioDirectory, audio, i)
	}

	wg.Wait()

	// adding metadata
	fmt.Println("Writing metadata...")
	md := Metadata{
		Uuid:        studioPack.Uuid,
		Ref:         GetRefFromUUid(studioPack.Uuid),
		Title:       studioPack.Title,
		Description: studioPack.Description,
	}

	yaml, err := yaml.Marshal(&md)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(tempPath, "md"), yaml, 0777)
	if err != nil {
		return err
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
	_ = os.RemoveAll(tempPath)

	return nil
}

func convertAndWriteAudio(reader zip.ReadCloser, deviceAudioDirectory string, audio Asset, index int) error {
	defer wg.Done()

	file, err := reader.Open("assets/" + audio.SourceName)
	if err != nil {
		return err
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	mp3, err := AudioToMp3(fileBytes)
	if err != nil {
		return err
	}

	cypheredFile := cipherFirstBlockCommonKey(mp3)

	err = os.WriteFile(filepath.Join(deviceAudioDirectory, intTo8Chars(index)), cypheredFile, 0777)
	if err != nil {
		return err
	}

	return nil
}

func convertAndWriteImage(reader zip.ReadCloser, deviceImageDirectory string, image Asset, index int) error {
	defer wg.Done()

	file, err := reader.Open("assets/" + image.SourceName)
	if err != nil {
		return err
	}
	bmpFile, err := ImageToBmp4(file)
	if err != nil {
		return err
	}
	cypheredBmp := cipherFirstBlockCommonKey(bmpFile)
	err = os.WriteFile(filepath.Join(deviceImageDirectory, intTo8Chars(index)), cypheredBmp, 0777)
	if err != nil {
		return err
	}
	return nil
}
