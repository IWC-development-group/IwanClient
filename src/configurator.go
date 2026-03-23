package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var (
	BaseIPS = [...]string{
		"localhost",
	}
)

const (
	BasePort = 8080
)

type Configurator struct {
	IPS  []string `json:"IPS"`
	Port int      `json:"Port`
}

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) InitConfig() {
	_, err := os.Stat("./Config/IwanConfig.json")
	if os.IsNotExist(err) {
		InitFiles()
	}

	jsonData, err := os.ReadFile("./Config/IwanConfig.json")
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
	err := os.MkdirAll("./Config", 0755)
	if err != nil {
		panic("Can't create config directory" + err.Error())
	}

	mainConfigFile, err := os.Create("./Config/IwanConfig.json")
	if err != nil {
		panic("Can't create main configuration file" + err.Error())
	}
	defer mainConfigFile.Close()

	base := Configurator{
		IPS:  BaseIPS[:],
		Port: BasePort,
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
