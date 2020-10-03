package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)


//filename for the bot config should always be cfg.json
//and always exist in the folder with bot.go
const cfg = "cfg.json"


//struct for bot command prefix and token
type Botcfg struct {
	Prefix string
	Token  string
}


//Init initializes bot with token & prefix for commands
func (botcfg *Botcfg) Init() {

	//read prefix and token from .json file
	data, err := ioutil.ReadFile("cfg.json")
	if err != nil {
		fmt.Println(err)
	}

	//unmarshal json into Botcfg struct
	err = json.Unmarshal(data, botcfg)
	if err != nil {
		fmt.Println(err)
	}

	return
}