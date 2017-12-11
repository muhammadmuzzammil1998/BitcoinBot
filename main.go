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
					if strings.Split(message, " ")[1] == "help" {
						s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
							Title: "BitcoinBot Help",
							Color: 0xf4a435,
							Fields: []*discordgo.MessageEmbedField{
								&discordgo.MessageEmbedField{
									Name:   "Usage",
									Value:  ">btc <currency> or @BitcoinBot#9430 <currency>",
									Inline: false,
								},
								&discordgo.MessageEmbedField{
									Name:   "Examples",
									Value:  ">btc, >btc USD, @BitcoinBot#9430, @BitcoinBot#9430 usd",
									Inline: false,
								},
							}})
					}
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

				s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Title:       "Bitcoin Price",
					Color:       0xf4a435,
					Description: "Current Bitcoin per " + strings.ToUpper(curr) + " price:",
					Fields: []*discordgo.MessageEmbedField{
						&discordgo.MessageEmbedField{
							Name:   data["data"]["currency"],
							Value:  data["data"]["amount"],
							Inline: true,
						},
					}})
			}
		})
	err = discord.Open()
	discord.UpdateStatus(0, ">btc help")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Bitcoin Bot is up!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
