package lunii

import (
	"os"
	"path/filepath"
)

func (device *Device) AddPack(studioPack StudioPack) error {
	// 1. Get path on devide
	rootPath := device.MountPoint
	contentPath := filepath.Join(rootPath, ".content", studioPack.Ref)

	// create directory for pack's files
	err := os.MkdirAll(contentPath, 0700)
	if err != nil {
		return err
	}

	// create image index - with exporter & lookup
	// create sound index - with exporter & lookup

	// write ni (index of stage nodes)
	// write li (index of action nodes)

	// write ri (index of image assets)
	// write si (index of sound assets)

	// copy images in rf
	// soupy sounds in sf

	// create boot file

	// update .pi root file with uuid
	err = device.AddPackToIndex(studioPack.Uuid)
	if err != nil {
		return err
	}

	return nil
}
