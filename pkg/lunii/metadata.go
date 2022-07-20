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
	Uuid           uuid.UUID `yaml:"uuid" json:"uuid"`
	Ref            string    `yaml:"ref" json:"ref"`
	Title          string    `yaml:"title" json:"title"`
	Description    string    `yaml:"description" json:"description"`
	IsOfficialPack bool      `yaml:"-" json:"isOfficialPack"`
	IsUnknown      bool      `yaml:"-" json:"isUnknown"`
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
			metadata, _ = GetMetadataFromDb(storyUuid)
		}
		if metadata == nil {
			metadata = &Metadata{
				Uuid:           storyUuid,
				Ref:            GetRefFromUUid(storyUuid),
				Title:          "",
				Description:    "",
				IsOfficialPack: false,
				IsUnknown:      true,
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

	metadata.IsOfficialPack = false
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
		Uuid:           uuid,
		Ref:            GetRefFromUUid(uuid),
		Title:          story.Title,
		Description:    story.Description,
		IsOfficialPack: true,
		IsUnknown:      false,
	}, nil
}
