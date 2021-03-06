package account

import (
	"errors"
	"io/ioutil"
	"strings"
	"encoding/json"
	// "fmt"
)

type AccountConfig struct {
	Name			string
	ClientId		string
	ClientSecret	string
	Platform		string
	AuthType		string
	CallBack		string
}

var DefaultConfig = &AccountConfig{
	Name:"poetnote1",
	Platform:"poetnote",
	ClientId:"1",
	ClientSecret:"Y1KU5rX5mCtEolcTOF44ftGGy8LRn0pH1jlO7qIU",
	AuthType:"password",
	CallBack:"http://kindle.poetnote.com/poetnote/callback/",
}



func NewDefaultConfig() *AccountConfig {
	return DefaultConfig
}

func NewConfig(configFile string) *AccountConfig {
	if configFile != "" {
		var config = &AccountConfig{}
		err := parseFile(config, configFile)
		// fmt.Printf("%+v",config)
		if err == nil {
			return config
		}
	}
	return NewDefaultConfig()
}

func parseFile(config interface{},file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	switch {
	case strings.HasSuffix(file,".json"):
		err = unmarshalJSON(data,config)
		if err != nil {
			return err
		}
	default :
		return errors.New("invalid config file")
	}
	return nil
}
//解析json数据
func unmarshalJSON(data []byte,config interface{}) error{
	if json.Valid(data) != true {
		return errors.New("invalid json string")
	}
	err := json.Unmarshal(data,config)
	// fmt.Printf("%+v",config)
	if err != nil {
		return err
	}
	return nil
}