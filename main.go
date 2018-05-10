package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

//Global Variables
var (
	codename = "Blueberry"
	version  = "3.0"
	vColor   = 0x3498db
	api      = "https://api.coinbase.com/v2/prices/spot?currency="
)

//main function
func main() {
	var token string
	flag.StringVar(&token, "token", "nil", "Bot token")
	flag.Parse()
	discord, err := discordgo.New("Bot " + string(token))
	if err != nil {
		log.Println(err)
		return
	}
	discord.AddHandler(Response)
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

//GetPrice returns price
func GetPrice(currency string) (string, string, error) {
	tStart := GetTime()
	resp, err := http.Get(api + currency)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Invalid response for " + currency + " | <" + resp.Status + ">")
		tEnd := GetTime()
		return "Invalid " + resp.Status, strconv.FormatInt(tEnd-tStart, 10), nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	data := map[string]map[string]string{}
	json.Unmarshal(body, &data)
	tEnd := GetTime()
	return data["data"]["amount"], strconv.FormatInt(tEnd-tStart, 10), nil
}

//Response for commands
func Response(s *discordgo.Session, m *discordgo.MessageCreate) {
	go UpdateStatus(s)
	message := strings.ToLower(strings.TrimSpace(m.Content))
	if strings.HasPrefix(message, ">btc") || strings.Contains(strings.Split(message, " ")[0], s.State.User.ID) {
		curr := "USD"
		if strings.Contains(message, " ") {
			if strings.Split(message, " ")[1] == "help" {
				_, t, _ := GetPrice("USD")
				s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Title: "BitcoinBot Help",
					Color: vColor,
					Fields: []*discordgo.MessageEmbedField{
						CreateField("Usage", ">btc <currency> or @BitcoinBot <currency>", false),
						CreateField("Examples", ">btc, >btc USD, @BitcoinBot, @BitcoinBot usd", false),
						CreateField("BitcoinBot's BTC Address", "3KyXwJhu1FpaPukJnzG9bPzn46xJ2ggTAs", false),
						CreateField("Version", codename+" ("+version+")", true),
						CreateField("Website", "https://bit.ly/btcbot", true),
						CreateField("API Latency", t+"ms", true),
					},
				})
			}
			curr = strings.Split(message, " ")[1]
		}
		rate, _, err := GetPrice(curr)
		if err != nil {
			log.Println(err)
			Report(s, m.ChannelID)
			return
		}
		if strings.Contains(rate, "Invalid") {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "Error",
				Color:       0xe74c3c,
				Description: "Invalid response: " + strings.TrimPrefix(rate, "Invalid "),
				Fields: []*discordgo.MessageEmbedField{
					CreateField("Report", "https://bit.ly/btcBotReport", false),
					CreateField("Email", "bitcoinbot@muzzammil.xyz", false),
				},
			})
			return
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Bitcoin Price",
			Color:       0xf4a435,
			Description: "Current Bitcoin per " + strings.ToUpper(curr) + " price:",
			Fields: []*discordgo.MessageEmbedField{
				CreateField(strings.ToUpper(curr), rate, true),
			},
		})
		/*	For server count management: START	*/
		f, err := os.OpenFile("/path/to/bitcoinbot/servers", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			Report(s, m.ChannelID)
			return
		}
		f.WriteString(m.ChannelID + "\n")
		f.Close()
		/*	For server count management: END	*/
	}
}

//UpdateStatus updates status on discord
func UpdateStatus(discord *discordgo.Session) {
	rate, t, err := GetPrice("USD")
	if err != nil {
		log.Println(err)
		return
	}
	discord.UpdateStatus(0, "$"+strings.Split(rate, ".")[0]+" | "+t+"ms"+" | >btc help")
}

//CreateField creates Message Embed Field and returns its address
func CreateField(name string, value string, inline bool) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}

//GetTime returns unix timestamp in milliseconds
func GetTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Report error
func Report(s *discordgo.Session, m string) {
	s.ChannelMessageSendEmbed(m, &discordgo.MessageEmbed{
		Title:       "Error",
		Color:       0xe74c3c,
		Description: "An error occurred. Please report this so I can fix this.",
		Fields: []*discordgo.MessageEmbedField{
			CreateField("Report", "https://bit.ly/btcBotReport", false),
			CreateField("Email", "bitcoinbot@muzzammil.xyz", false),
		},
	})
}
