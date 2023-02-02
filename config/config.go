package config

import (
	"encoding/json"
	"healthcheck/errors"
	"io/ioutil"
	"os"
)

var conf *JSONStruct

func Get() *JSONStruct {
	InitDb()
	if conf == nil {
		conf = &JSONStruct{}
	}
	return conf
}

func (conf *JSONStruct) Make(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	defer file.Close()
	conf.parse(file)
}

func (conf *JSONStruct) parse(file *os.File) {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	err = json.Unmarshal(bytes, conf)
	if err != nil {
		errors.Print(err.Error())
	}
}
