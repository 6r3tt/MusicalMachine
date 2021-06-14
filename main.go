package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "NjEyMzMxNjExMjA3MzY4NzYx.XVg0QA.NU7rfhDpYIdeS1C8_Wr5kmmv5RQ"

var BotID string = "612331611207368761"

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}
