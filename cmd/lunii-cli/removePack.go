package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/olup/lunii-cli/pkg/lunii"
	"github.com/urfave/cli/v2"
)

func yesNo(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	return result == "Y" || result == "y"
}

func RemovePack(c *cli.Context) error {
	packRef := c.Args().Get(0)
	if packRef == "" {
		color.Red("Ref of pack is required")
		return errors.New("Ref of pack is required")
	}
	if yesNo(fmt.Sprintf("Are you sure you want to remove %s from your device ?", packRef)) == false {
		return nil
	}
	device, err := lunii.GetDevice()
	if err != nil {
		return err
	}
	err = device.RemovePackFromIndexFromRef(packRef)
	if err != nil {
		return err
	}
	contentPath := filepath.Join(device.MountPoint, ".content", packRef)

	_, err = os.Stat(contentPath)
	if err == nil {
		fmt.Println("Removing content directory")
		os.RemoveAll(contentPath)
	}

	return nil
}
