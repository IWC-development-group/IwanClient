package iwanCore

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) AddUrl(URL string) {
	c.InitConfig()
	c.URLS = append(c.URLS, URL)

	pathToEx, err := os.Executable()
	if err != nil {
		Log(err.Error())
		os.Exit(1)
	}

	_, err = os.Stat(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if os.IsNotExist(err) {
		fmt.Println("Unknown error!")
		os.Exit(1)
	}

	updatedData, err := json.MarshalIndent(c, "", "	")

	err = os.WriteFile(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"), updatedData, 0755)
	if err != nil {
		fmt.Println("Error write config file")
		os.Exit(1)
	}

	fmt.Println("Successfully added new URL!")
}

func (c *Configurator) InitConfig() {
	pathToEx, err := os.Executable()
	if err != nil {
		Log(err.Error())
		os.Exit(1)
	}

	_, err = os.Stat(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if os.IsNotExist(err) {
		InitFiles()
	}

	jsonData, err := os.ReadFile(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if err != nil {
		Log("Error read config file")
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(jsonData), &c)
	if err != nil {
		Log("Error parse config file")
		os.Exit(1)
	}

	Log("Successfully loaded config!")
}

func InitFiles() {
	pathToEx, err := os.Executable()
	if err != nil {
		Log(err.Error())
		os.Exit(1)
	}

	err = os.MkdirAll(path.Join(filepath.Dir(pathToEx), "Config"), 0755)
	if err != nil {
		Log("Can't create config directory" + err.Error())
		os.Exit(1)

	}

	mainConfigFile, err := os.Create(path.Join(filepath.Dir(pathToEx), "Config/IwanConfig.json"))
	if err != nil {
		Log("Can't create main configuration file" + err.Error())
		os.Exit(1)
	}
	defer mainConfigFile.Close()

	base := Configurator{
		URLS: BaseURLS,
	}

	jsonData, err := json.MarshalIndent(base, "", "	")
	if err != nil {
		Log("Bad config base, can't parse!")
		os.Exit(1)
	}

	length, err := mainConfigFile.Write(jsonData)
	if err != nil {
		Log("Can't save config base to file")
		os.Exit(1)
	}

	Log("Saved config base with " + strconv.Itoa(length) + " bytes")
}
