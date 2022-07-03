package studiopackbuilder

import "io/ioutil"

func CopyFile(from string, to string) error {
	input, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(to, input, 0777)
	if err != nil {
		return err
	}
	return nil
}
