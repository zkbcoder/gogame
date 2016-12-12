package util

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadXml(filename string, config interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, config)
	if err != nil {
		return err
	}
	return nil
}
