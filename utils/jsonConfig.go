package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configs struct{
	Master string `json:"master"`
	Port int `json:"port"`
	DB db `json:"db"`
	Tasks tasks `json:"tasks"`
}

type db struct {
	DbHost string `json:"dbHost"`
	DbUser string `json:"dbUser"`
	DbName string `json:"dbName"`
	DbPass string `json:"dbPass"`
}

type tasks struct {
	Trend string `json:"trend"`
} 

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (c *JsonStruct) Load(fileName string, v interface{}) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

var Config Configs

func InitRedisConfigs(conf string) {

	fmt.Println("path",conf)

	JsonParse := NewJsonStruct()
	JsonParse.Load(conf, &Config)

}