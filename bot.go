package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//snitch "github.com/michael-j-taylor/rust-snitch/rust-snitch/rustsnitch"
	"github.com/bwmarrin/discordgo"
)

//struct for bot command prefix and token
type botcfg struct {
	Prefix string
	Token  string
}

func main() {

	//read prefix and token from .json file
	data, err := ioutil.ReadFile("cfg.json")
	if err != nil {
		fmt.Println(err)
	}

	var cfg botcfg

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg.Prefix + " " + cfg.Token)

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//register onMessageCreate as handler for MessageCreate events
	dg.AddHandler(onMessageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	//open websocket connection
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running")

	//block forever
	<-make(chan int)
}

//onMessageCreate is a called every time a new message is created in a
//channel the bot can access
func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//ignore messages sent by bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	//TODO: handle commands
	return
}
