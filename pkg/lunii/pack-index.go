package lunii

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Story struct {
	uuid       uuid.UUID
	folderName string
}

func ReadGlobalIndexFile() []Story {

	stories := []Story{}
	// read .pi file and get
	data, err := os.Open("/Volumes/lunii/.pi")
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}

		uuid, err := uuid.FromBytes(slice)
		if err != nil {
			log.Fatal(err)
		}

		uuidString := uuid.String()
		story := Story{
			uuid:       uuid,
			folderName: strings.ToUpper(strings.ReplaceAll(uuidString[len(uuidString)-8:], "_", "")),
		}
		stories = append(stories, story)
	}
	return stories
}

func WriteGlobalIndexFile(stories []Story) {

}
