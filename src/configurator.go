package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var (
	BaseURLS = []string{
		"http://localhost:8080",
	}
)

type Configurator struct {
	URLS []string `json:"URLS"`
}

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) InitConfig() {
	pathToEx, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}

	_, err = os.Stat(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if os.IsNotExist(err) {
		InitFiles()
	}

	jsonData, err := os.ReadFile(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if err != nil {
		panic("Error read config file")
	}

	err = json.Unmarshal([]byte(jsonData), &c)
	if err != nil {
		panic("Error parse config file")
	}

	fmt.Println("Successfully loaded config!")
}

func InitFiles() {
	pathToEx, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}

	err = os.MkdirAll(path.Join(filepath.Dir(pathToEx), "Config"), 0755)
	if err != nil {
		panic("Can't create config directory" + err.Error())
	}

	mainConfigFile, err := os.Create(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if err != nil {
		panic("Can't create main configuration file" + err.Error())
	}
	defer mainConfigFile.Close()

	base := Configurator{
		URLS: BaseURLS,
	}

	jsonData, err := json.MarshalIndent(base, "", "	")
	if err != nil {
		panic("Bad config base, can't parse!")
	}

	length, err := mainConfigFile.Write(jsonData)
	if err != nil {
		panic("Can't save config base to file")
	}

	fmt.Println("Saved config base with " + strconv.Itoa(length) + " bytes")
}
