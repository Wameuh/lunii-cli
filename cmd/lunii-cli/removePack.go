package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
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
	packUuid := c.Args().Get(0)
	if packUuid == "" {
		color.Red("Ref of pack is required")
		return errors.New("Ref of pack is required")
	}
	if yesNo(fmt.Sprintf("Are you sure you want to remove %s from your device ?", packUuid)) == false {
		return nil
	}
	fmt.Sprintln("", packUuid)
	return nil
}
