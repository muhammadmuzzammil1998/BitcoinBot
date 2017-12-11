package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "YOUR.BOT.TOKEN")
	if err != nil {
		log.Println(err)
		return
	}
	discord.AddHandler(
		func(s *discordgo.Session, m *discordgo.MessageCreate) {
			message := strings.TrimSpace(m.Content)
			if strings.HasPrefix(message, ">btc") || strings.HasPrefix(message, "<@388984248062967819>") {
				curr := "USD"
				if strings.Contains(message, " ") {
					curr = strings.Split(message, " ")[1]
				} else {
					curr = "USD"
				}
				curlData, err := exec.Command("/usr/bin/curl", "https://api.coinbase.com/v2/prices/spot?currency="+curr).Output()
				if err != nil {
					log.Println(err)
					return
				}
				data := map[string]map[string]string{}
				json.Unmarshal(curlData, &data)

				d := discordgo.MessageEmbed{
					Title:       "Bitcoin Price",
					Color:       0xf4a435,
					Description: "Current Bitcoin per " + strings.ToUpper(curr) + " price:",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   data["data"]["currency"],
							Value:  data["data"]["amount"],
							Inline: true,
						},
					}}
				s.ChannelMessageSendEmbed(m.ChannelID, &d)
			}
		})
	err = discord.Open()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Bitcoin Bot is up!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
