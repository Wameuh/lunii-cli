package lunii

import (
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
)

func getStudioDbPath() string {
	dirname, _ := os.UserHomeDir()
	return filepath.Join(dirname, ".studio", "db", "unofficial.json")
}

func GetStudioMetadataDb(dbPath string) (*Db, error) {
	var db *Db
	// if path is empty get default db path
	if dbPath == "" {
		dbPath = getStudioDbPath()
	}

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

		db.stories = append(db.stories, Story{
			Uuid:  uuid.MustParse(storyUuid),
			Title: title, Description: description,
			PackType: "custom",
		})
		return nil
	}, "response")

	return db, nil
}
