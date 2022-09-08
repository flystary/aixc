package net

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"aixc/conf"

)

func loadURL(path string) (conf.Rules, error) {
		var rules conf.Rules
		io, err := ioutil.ReadFile(path)
		if err != nil {
			return rules, err
		}
		err = yaml.Unmarshal(io, &rules) 
		if err != nil {
			return rules, err
		}
		return rules, nil
}