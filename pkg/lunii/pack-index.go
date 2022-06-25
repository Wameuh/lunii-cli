package lunii

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type PackMetadata struct {
	Uuid       uuid.UUID
	FolderName string
	Title      string
}

func GetFolderNameFromUUid(uuid uuid.UUID) string {
	uuidString := uuid.String()
	return strings.ToUpper(strings.ReplaceAll(uuidString[len(uuidString)-8:], "_", ""))
}

func (*Device) ReadGlobalIndexFile() ([]PackMetadata, error) {

	packs := []PackMetadata{}
	// read .pi file and get
	data, err := os.Open("/Volumes/lunii/.pi")
	if err != nil {
		return nil, errors.New("Could not reach the pack index file")
	}
	defer data.Close()

	luniiDb, err := GetMetadataDb()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(data)
	slice := make([]byte, 16)

	for {
		_, err = reader.Read(slice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("There was an error reading the pack index file")
		}

		uuid, err := uuid.FromBytes(slice)
		if err != nil {
			return nil, errors.New("There was an error getting UUID from the pack index file")
		}

		dbMetadata := luniiDb.GetStoryById(uuid)
		storyTitle := "..."
		if dbMetadata != nil {
			storyTitle = dbMetadata.Title
		}
		pack := PackMetadata{
			Uuid:       uuid,
			FolderName: GetFolderNameFromUUid(uuid),
			Title:      storyTitle,
		}
		packs = append(packs, pack)
	}
	return packs, nil
}

func (*Device) WriteGlobalIndexFile(stories []PackMetadata) error {
	var buf []byte
	for _, story := range stories {
		buf = append(buf, story.Uuid[:]...)
	}
	fmt.Println(buf)
	err := os.WriteFile(filepath.Join("/Volumes/lunii", ".pi"), buf, 0777)
	return err
}

func (device *Device) AddPackToIndex(uuid uuid.UUID) error {
	stories, err := device.ReadGlobalIndexFile()
	if err != nil {
		return err
	}
	stories = append(stories, PackMetadata{Uuid: uuid, FolderName: GetFolderNameFromUUid(uuid)})
	err = device.WriteGlobalIndexFile(stories)
	return err
}

func (device *Device) RemovePackFromIndex(uuid uuid.UUID) error {
	var stories []PackMetadata

	deviceStories, err := device.ReadGlobalIndexFile()
	if err != nil {
		return err
	}

	// filter out the specified UUID
	for _, story := range deviceStories {
		if story.Uuid != uuid {
			stories = append(stories, story)
		}
	}

	err = device.WriteGlobalIndexFile(stories)

	return err
}
