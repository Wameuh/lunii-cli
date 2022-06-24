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

type Story struct {
	uuid       uuid.UUID
	folderName string
}

func ReadGlobalIndexFile() ([]Story, error) {

	stories := []Story{}
	// read .pi file and get
	data, err := os.Open("/Volumes/lunii/.pi")
	if err != nil {
		return nil, errors.New("Could not reach the pack index file")
	}
	defer data.Close()

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

		uuidString := uuid.String()
		story := Story{
			uuid:       uuid,
			folderName: strings.ToUpper(strings.ReplaceAll(uuidString[len(uuidString)-8:], "_", "")),
		}
		stories = append(stories, story)
	}
	return stories, nil
}

func WriteGlobalIndexFile(stories []Story) error {
	var buf []byte
	for _, story := range stories {
		buf = append(buf, story.uuid[:]...)
	}
	fmt.Println(buf)
	err := os.WriteFile(filepath.Join("/Volumes/lunii", ".pi"), buf, 0777)
	return err
}
