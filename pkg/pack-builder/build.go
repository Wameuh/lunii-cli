package studiopackbuilder

import (
	"os"
	"path/filepath"

	"github.com/olup/lunii-cli/pkg/lunii"
	"gopkg.in/yaml.v3"
)

func buildStudioPack(directoryPath string, outputPath string) (*lunii.StudioPack, error) {
	var metadata lunii.Metadata
	metadataPath := filepath.Join(directoryPath, "metadata.yaml")
	metadataBytes, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(metadataBytes, metadata)
	if err != nil {
		return nil, err
	}

	tempOutputPath := filepath.Join(os.TempDir(), "build", metadata.Ref)
	tempOutputAssetPath := filepath.Join(tempOutputPath, "assets")

	err = os.MkdirAll(tempOutputPath, 0700)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(tempOutputAssetPath, 0700)
	if err != nil {
		return nil, err
	}

	// start node grabbing
	stageNodes, listNodes, err := GetTitleNode(directoryPath, tempOutputAssetPath)
	if err != nil {
		return nil, err
	}

	return &lunii.StudioPack{
		Uuid:        metadata.Uuid,
		Title:       metadata.Title,
		Ref:         metadata.Ref,
		Description: metadata.Description,

		StageNodes: stageNodes,
		ListNodes:  listNodes,
		PackType:   "",
		Version:    2,
	}, nil
}
