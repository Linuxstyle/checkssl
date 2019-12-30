package conf

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type ConfInfo struct {
    Mails `yaml:"mails"`
    Domains `yaml:"domains"`
}
type Mails struct {
    Mail []string `yaml:"list"`
}
type Domains struct {
    Domain []string `yaml:"list"`
}

func NewConfigInfo() (*ConfInfo, error){
    var config  ConfInfo
    context ,err :=  ioutil.ReadFile("./conf/config.yaml")
    if err != nil {
        log.Println("Read Yaml File Failed:  ",err)
        return nil,err
    }
    yaml.Unmarshal(context,&config)
    //log.Println(config)
    return &config,nil
}
