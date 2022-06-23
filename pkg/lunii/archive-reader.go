package lunii

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
)

type Pack struct {
	PackType     string
	OriginalPath string
	Format       string      `json:"format"` // enum ?
	Title        string      `json:"title"`
	Version      int         `json:"version"` // enum ?
	Description  string      `json:"description"`
	StageNodes   []StageNode `json:"stageNodes"`
	ListNodes    []ListNode  `json:"actionNodes"`
}

type StageNode struct {
	Uuid            uuid.UUID        `json:"uuid"`
	Type            string           `json:"type"` //enum ?
	Name            string           `json:"name"`
	Image           string           `json:"image"`
	Audio           string           `json:"audio"`
	OkTransition    *Transition      `json:"okTransition"`
	HomeTransition  *Transition      `json:"homeTransition"`
	ControlSettings *ControlSettings `json:"controlSettings"`
	SquareOne       bool             `json:"squareOne"`
}

type ListNode struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Options []uuid.UUID `json:"options"`
}

type ControlSettings struct {
	Wheel    bool `json:"wheel"`
	Ok       bool `json:"ok"`
	Home     bool `json:"home"`
	Pause    bool `json:"pause"`
	Autoplay bool `json:"autoplay"`
}

type Transition struct {
	ActionNode  string `json:"actionNode"`
	OptionIndex int    `json:"optionIndex"`
}

func ReadPack(path string) (*Pack, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, errors.New("package could not be found")
	}
	defer reader.Close()

	file, err := reader.Open("story.json")
	if err != nil {
		log.Fatal("story.json could not be found")
	}
	fileAsBytes, _ := ioutil.ReadAll(file)

	var pack Pack

	pack.OriginalPath = path
	json.Unmarshal(fileAsBytes, &pack)

	return &pack, nil
}

func (pack *Pack) writeToDevice() error {

	// get path from uuid
	// create directory on device

	// create image index - with exporter & lookup
	// create sound index - with exporter & lookup

	// write ni (index of stage nodes)
	// write li (index of action nodes)

	// write ri (index of image assets)
	// write si (index of sound assets)

	// copy images in rf
	// soupy sounds in sf

	// create boot file
	// update .pi root file with uuid

	return nil
}
