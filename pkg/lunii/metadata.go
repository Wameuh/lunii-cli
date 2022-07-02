package lunii

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Uuid        uuid.UUID `yaml:"uuid"`
	Ref         string    `yaml:"ref"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
}

func (device *Device) GetPacks() ([]Metadata, error) {
	var packs []Metadata

	uuids, err := device.ReadGlobalIndexFile()
	if err != nil {
		return nil, err
	}
	luniiDb, _ := GetLuniiStoreDb()

	for _, storyUuid := range uuids { // Read md file
		metadata, _ := GetMetadataFromDevice(storyUuid, device)
		if metadata == nil && luniiDb != nil {
			metadata, _ = GetMetadataFromDb(storyUuid)
		}
		if metadata == nil {
			metadata = &Metadata{
				Uuid:        storyUuid,
				Ref:         GetRefFromUUid(storyUuid),
				Title:       "",
				Description: "",
			}
		}
		packs = append(packs, *metadata)
	}
	return packs, nil
}

func GetMetadataFromDevice(uuid uuid.UUID, device *Device) (*Metadata, error) {
	path := filepath.Join(device.MountPoint, ".content", GetRefFromUUid(uuid))
	metadataFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var metadata Metadata
	err = yaml.Unmarshal(metadataFile, metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func GetMetadataFromDb(uuid uuid.UUID) (*Metadata, error) {
	luniiDb, err := GetMetadataDb()
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
	}, nil
}
