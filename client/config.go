package client

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type Account struct {
	PhoneNum string `yaml:"phoneNum"`
	Password string `yaml:"password"`
}

type Configer struct {
	Srv struct {
		Addr    string `yaml:"addr"`
	}

	Log struct {
		Level    int  `yaml:"level"`
		Size     int  `yaml:"size"`
		Json     bool `yaml:"json"`
	}

	Accounts []Account `yaml:"accounts"`
}

var cfgIns *Configer
var once sync.Once

func Config() *Configer {

	once.Do(func() {
		cfgIns = &Configer{}
	})
	return cfgIns
}


func (c *Configer) Load(filename string) {

	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		goto FAILED
	}

	err = yaml.Unmarshal(buff, c)
	if err != nil {
		goto FAILED
	}
	return

FAILED:
	fmt.Printf("failed:%v",err)
	os.Exit(1)
}
