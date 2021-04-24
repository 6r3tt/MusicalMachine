package main

import (
	"fmt"

	"github.com/dahop/discordgo"
)

const token string = "NjEyMzMxNjExMjA3MzY4NzYx.XVg0QA.3VDDI1U8zOx84y94EHMTxlApKNs"

var BotID string = ""

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
