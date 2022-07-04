// TODO : returns
package studiopackbuilder

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/olup/lunii-cli/pkg/lunii"
)

func getTitleNode(directoryPath string, tempOutputAssetPath string) ([]lunii.StageNode, []lunii.ListNode, error) {

	// Create title node
	nodeUuid := uuid.New()
	titleNode := lunii.StageNode{
		Uuid: nodeUuid,
		Name: filepath.Base(directoryPath) + "_title_node",
		Type: "node",
	}

	// audio
	audioPath := filepath.Join(directoryPath, "title.mp3")
	_, err := os.Stat(audioPath)
	if err != nil {
		return nil, nil, err
	}

	audioFileName := uuid.NewString() + ".mp3"
	err = copyFile(audioPath, filepath.Join(tempOutputAssetPath, audioFileName))
	if err != nil {
		return nil, nil, err
	}

	titleNode.Audio = audioFileName

	// cover
	imagePath := filepath.Join(directoryPath, "cover.jpeg")
	_, err = os.Stat(imagePath)
	if err != nil {
		return nil, nil, err
	}

	imageFileName := uuid.NewString() + ".png"
	copyFile(imagePath, filepath.Join(tempOutputAssetPath, imageFileName))
	if err != nil {
		return nil, nil, err
	}

	titleNode.Image = imageFileName

	// control settings for title node

	titleNode.ControlSettings = &lunii.ControlSettings{
		Wheel:    true,
		Ok:       true,
		Home:     true,
		Pause:    false,
		Autoplay: false,
	}

	titleNode.HomeTransition = nil

	// Is there a story node or more title nodes ?
	storyAudioPath := filepath.Join(directoryPath, "story.mp3")
	_, err = os.Stat(storyAudioPath)
	if err == nil {
		// We have a story node

		// copy audio
		audioFileName := uuid.NewString() + ".mp3"
		err = copyFile(storyAudioPath, filepath.Join(tempOutputAssetPath, audioFileName))
		if err != nil {
			return nil, nil, err
		}

		// create node
		storyNode := lunii.StageNode{
			Uuid:  uuid.New(),
			Name:  filepath.Base(directoryPath) + "_story_node",
			Audio: audioFileName,
			Type:  "node",
		}

		// create a list node
		listNode := lunii.ListNode{
			Id:      uuid.NewString(),
			Name:    filepath.Base(directoryPath) + "_story_list_node",
			Options: []uuid.UUID{storyNode.Uuid},
		}

		// add list node to title node
		titleNode.OkTransition = &lunii.Transition{
			ActionNode:  listNode.Id,
			OptionIndex: 0,
		}

		titleNode.HomeTransition = &lunii.Transition{
			ActionNode:  "", // TODO
			OptionIndex: 0,
		}

		// set story node control settings
		storyNode.ControlSettings = &lunii.ControlSettings{
			Wheel:    false,
			Ok:       false,
			Home:     true,
			Pause:    true,
			Autoplay: true,
		}

		// return nodes and lists
		return []lunii.StageNode{titleNode, storyNode}, []lunii.ListNode{listNode}, nil
	} else {
		// There is no story node - it is a title node
		stageNodes, listNodes, err := listNodesFomrDirectory(directoryPath, tempOutputAssetPath)
		if err != nil {
			return nil, nil, err
		}

		stageNodes = append([]lunii.StageNode{titleNode}, stageNodes...)
		return stageNodes, listNodes, nil
	}
}

func listNodesFomrDirectory(directoryPath string, tempOutputPath string) ([]lunii.StageNode, []lunii.ListNode, error) {
	var stageNodes []lunii.StageNode
	var listNodes []lunii.ListNode
	thisListNode := lunii.ListNode{
		Id:   uuid.NewString(),
		Name: filepath.Base(directoryPath) + "_title_list_node",
	}

	// read each files in directory
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, nil, err
	}

	// for each directory
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		thisStageNodes, thisListNodes, err := getTitleNode(filepath.Join(directoryPath, file.Name()), tempOutputPath)
		if err != nil {
			return nil, nil, err
		}
		stageNodes = append(stageNodes, thisStageNodes...)
		listNodes = append(listNodes, thisListNodes...)

		// link top stage node in this list
		thisListNode.Options = append(thisListNode.Options, stageNodes[0].Uuid)
	}

	finalListNodes := append([]lunii.ListNode{thisListNode}, listNodes...)

	return stageNodes, finalListNodes, nil
}
