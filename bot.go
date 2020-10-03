package main

import (
	"fmt"

	//snitch "github.com/michael-j-taylor/rust-snitch/rust-snitch/rustsnitch"
	"github.com/bwmarrin/discordgo"
)

func main() {

	var cfg Botcfg
	cfg.Init()

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

	fmt.Println("bot now running")

	//TODO: read data into memory from servers.json and fire off
	//appropriate goroutines

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


