package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var CONFIG_REL_DIR = "../configs"

type ConfigProvider interface{
	GetConfig(filename, configName string) interface{}
	GetConfigI(filename, configName string) int
	GetConfigI32(filename, configName string) int32
	GetConfigI64(filename, configName string) int64
	GetConfigString(filename, configName string) string
}

type configs struct{
	readFiles map[string]map[string]interface{}
}

func (c *configs) GetConfigI32(filename, configName string) int32 {
	return int32(c.GetConfig(filename,configName).(float64))
}

func (c *configs) GetConfigI64(filename, configName string) int64 {
	return int64(c.GetConfig(filename,configName).(float64))
}

func (c *configs) GetConfigI(filename, configName string) int{
	return int(c.GetConfig(filename,configName).(float64))
}

func (c *configs) GetConfigString(filename, configName string) string {
	return c.GetConfig(filename,configName).(string)
}

func MakeConfigs() ConfigProvider{
	res := configs{}
	res.readFiles = make(map[string]map[string]interface{})
	return &res
}



func (c *configs) readFile(fileName string){
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s.json",CONFIG_REL_DIR,fileName))
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	c.readFiles[fileName] = result
}

func (c* configs) GetConfig(fileName, configName string) interface{}{
	file, ok := c.readFiles[fileName]
	if !ok{
		c.readFile(fileName)
		file, _ = c.readFiles[fileName]
	}
	return file[configName]
}
