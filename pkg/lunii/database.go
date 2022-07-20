package lunii

import (
	"log"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
)

type Story struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Uuid        uuid.UUID `json:"uuid"`
}

type Db struct {
	stories []Story
}

var db *Db = nil

func GetLuniiStoreDb() ([]byte, error) {
	log.Println("Fetching db from lunii store")

	resp, err := req.Get("https://server-auth-prod.lunii.com/guest/create")
	if err != nil {
		return nil, err
	}
	token, err := jsonparser.GetString(resp.Bytes(), "response", "token", "server")
	if err != nil {
		return nil, err
	}
	resp, err = req.SetHeader("X-AUTH-TOKEN", token).Get("https://server-data-prod.lunii.com/v2/packs")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

func GetMetadataDb() (*Db, error) {
	// Return DB if in cache
	if db != nil {
		return db, nil
	}

	// create DB
	db = &Db{}

	// read DB from temp file
	tempPath := os.TempDir()
	dbBytes, err := os.ReadFile(filepath.Join(tempPath, "lunii-store-db"))

	// if no temp file, fetch if from lunii api
	if err != nil {
		dbBytes, err := GetLuniiStoreDb()
		if err != nil {
			log.Println("Could not retrieve DB from lunii store. Skipping.")
		}
		err = os.WriteFile(filepath.Join(tempPath, "lunii-store-db"), dbBytes, 0777)
		if err != nil {
			log.Println("Could not write to temporary db file, skipping...")
		}
	}

	// parse db's json
	jsonparser.ObjectEach(dbBytes, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		storyUuid, err := jsonparser.GetString(value, "uuid")
		if err != nil {
			return err
		}

		title, _ := jsonparser.GetString(value, "localized_infos.fr_FR.title")
		description, _ := jsonparser.GetString(value, "localized_infos.fr_FR.description")

		db.stories = append(db.stories, Story{Uuid: uuid.MustParse(storyUuid), Title: title, Description: description})
		return nil
	}, "response")

	return db, nil
}

func (db *Db) GetStoryById(uuid uuid.UUID) *Story {
	for _, story := range db.stories {
		if story.Uuid == uuid {
			return &story
		}
	}
	return nil
}
