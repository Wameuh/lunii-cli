package lunii

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Uuid        uuid.UUID `yaml:"uuid" json:"uuid"`
	Ref         string    `yaml:"ref" json:"ref"`
	Title       string    `yaml:"title" json:"title"`
	Description string    `yaml:"description" json:"description"`
	PackType    string    `yaml:"packType" json:"packType"`
}

func (device *Device) GetPacks() ([]Metadata, error) {
	var packs []Metadata

	uuids, err := device.ReadGlobalIndexFile()
	if err != nil {
		return nil, err
	}

	for _, storyUuid := range uuids { // Read md file
		metadata, _ := GetMetadataFromDevice(storyUuid, device)
		if metadata == nil {
			metadata, _ = GetMetadataFromLuniiDb(storyUuid)
		}
		if metadata == nil {
			metadata, _ = GetMetadataFromStudioDb(storyUuid)
		}
		if metadata == nil {
			metadata = &Metadata{
				Uuid:        storyUuid,
				Ref:         GetRefFromUUid(storyUuid),
				Title:       "",
				Description: "",
				PackType:    "undefined",
			}
		}
		packs = append(packs, *metadata)
	}
	return packs, nil
}

func GetMetadataFromDevice(uuid uuid.UUID, device *Device) (*Metadata, error) {
	mdFilePath := filepath.Join(device.MountPoint, ".content", GetRefFromUUid(uuid), "md")
	metadataFile, err := os.ReadFile(mdFilePath)
	if err != nil {
		return nil, err
	}
	metadata := Metadata{}
	err = yaml.Unmarshal(metadataFile, &metadata)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	metadata.PackType = "custom"
	return &metadata, nil
}

func GetMetadataFromLuniiDb(uuid uuid.UUID) (*Metadata, error) {
	luniiDb, err := GetLuniiMetadataDb()
	if err != nil {
		return nil, err
	}
	story := luniiDb.GetStoryById(uuid)

	if story == nil {
		return nil, errors.New("Could not found this uuid in DB")
	}

	return &Metadata{
		Uuid:        uuid,
		Ref:         GetRefFromUUid(uuid),
		Title:       story.Title,
		Description: story.Description,
		PackType:    "official",
	}, nil
}

func GetMetadataFromStudioDb(uuid uuid.UUID) (*Metadata, error) {
	studioDb, err := GetLStudioMetadataDb()
	if err != nil {
		return nil, err
	}
	story := studioDb.GetStoryById(uuid)

	if story == nil {
		return nil, errors.New("Could not found this uuid in DB")
	}

	return &Metadata{
		Uuid:        uuid,
		Ref:         GetRefFromUUid(uuid),
		Title:       story.Title,
		Description: story.Description,
		PackType:    "custom",
	}, nil
}
