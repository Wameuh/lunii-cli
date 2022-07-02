package studiopackbuilder

// import (
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"

// 	"github.com/google/uuid"
// 	"github.com/olup/lunii-cli/pkg/lunii"
// 	"gopkg.in/yaml.v3"
// )

// func buildStudioPack(directoryPath string, outputPath string) (*lunii.StudioPack, error) {
// 	var metadata lunii.Metadata
// 	metadataPath := filepath.Join(directoryPath, "metadata.yaml")
// 	metadataBytes, err := os.ReadFile(metadataPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = yaml.Unmarshal(metadataBytes, metadata)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tempOutputPath := filepath.Join(os.TempDir(), "build", metadata.Ref)
// 	tempOutputAssetPath := filepath.Join(tempOutputPath, "assets")

// 	err = os.MkdirAll(tempOutputPath, 0700)
// 	err = os.MkdirAll(tempOutputAssetPath, 0700)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var nodeList []lunii.StageNode

// 	// start node grabbing
// 	nodeType := "title"

// 	// first node has the pack uuid, to keep compatibility with stuidio
// 	nodeUuid := metadata.Uuid
// 	nodeList[0] = lunii.StageNode{
// 		Uuid: nodeUuid,
// 	}

// 	// audio
// 	audioPath := filepath.Join(directoryPath, "title.mp3")
// 	_, err = os.Stat(audioPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	audioFileName := uuid.NewString() + ".mp3"
// 	err = copy(audioPath, filepath.Join(tempOutputAssetPath, audioFileName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	nodeList[0].Audio = audioFileName

// 	// cover
// 	imagePath := filepath.Join(directoryPath, "cover.png")
// 	_, err = os.Stat(imagePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	imageFileName := uuid.NewString() + ".png"
// 	copy(imagePath, filepath.Join(tempOutputAssetPath, imageFileName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	nodeList[0].Image = imageFileName

// 	storyAudioPath := filepath.Join(directoryPath, "title.mp3")
// 	_, err = os.Stat(audioPath)
// 	if err == nil {
// 		// create a story node
// 		// create a list node
// 		// add list node to title node
// 		// finish
// 	} else {
// 		// get a list form sub dire
// 	}

// 	// todo complete pack
// 	return &lunii.StudioPack{
// 		Uuid:        metadata.Uuid,
// 		Title:       metadata.Title,
// 		Description: metadata.Description,

// 		StageNodes: nodeList,
// 	}, nil
// }

// func copy(from string, to string) error {
// 	input, err := ioutil.ReadFile(from)
// 	if err != nil {
// 		return err
// 	}

// 	err = ioutil.WriteFile(to, input, 0777)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // func getNode(folder string, nodeList []lunii.StageNode) error {
// // 	nodeType := "title"
// // 	files, err := os.ReadDir(folder)
// // 	if err != nil {
// // 		return 0
// // 	}
// // 	for _, f := range files {

// // 	}

// // 	return nil
// // }
