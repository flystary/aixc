package net

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

func loadURL(path string)  error {
	io, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(io, &rules) 
	if err != nil {
		return  err
	}
	return nil
}