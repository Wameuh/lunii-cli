package lunii

import (
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
)

var studioDb *Db = nil

func getStudioDbPath() string {
	dirname, _ := os.UserHomeDir()
	return filepath.Join(dirname, ".studio", "db", "unofficial.json")
}

func GetLStudioMetadataDb() (*Db, error) {
	// Return DB if in cache
	if studioDb != nil {
		return studioDb, nil
	}

	// create DB
	studioDb = &Db{}

	// read DB from STUdio local DB
	dbBytes, err := os.ReadFile(getStudioDbPath())

	// if no temp file, return an error
	if err != nil {
		return nil, err
	}

	// parse db's json TODO
	jsonparser.ObjectEach(dbBytes, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		storyUuid, err := jsonparser.GetString(value, "uuid")
		if err != nil {
			return err
		}

		title, _ := jsonparser.GetString(value, "title")
		description, _ := jsonparser.GetString(value, "description")

		studioDb.stories = append(studioDb.stories, Story{Uuid: uuid.MustParse(storyUuid), Title: title, Description: description})
		return nil
	}, "response")

	return studioDb, nil
}
