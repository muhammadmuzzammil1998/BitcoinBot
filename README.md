# BitcoinBot · [![CircleCI](https://circleci.com/gh/muhammadmuzzammil1998/BitcoinBot.svg?style=svg)](https://circleci.com/gh/muhammadmuzzammil1998/BitcoinBot) 

A discord bot to check Bitcoin Price... yeah... just that...

[![Discord Bots](https://discordbots.org/api/widget/upvotes/388984248062967819.svg)](https://discordbots.org/bot/388984248062967819)
[![Discord Bots](https://discordbots.org/api/widget/servers/388984248062967819.svg)](https://discordbots.org/bot/388984248062967819) 
[![Discord Bots](https://discordbots.org/api/widget/status/388984248062967819.svg)](https://discordbots.org/bot/388984248062967819)

```diff
- I am not accepting any Pull Requests currently to avoid conflicts 
- because BitcoinBot is going under construction mode.  
```

## Build
    $ git clone https://github.com/muhammadmuzzammil1998/bitcoinbot.git
    $ cd bitcoinbot
    $ echo "YOUR.TOKEN.HERE" > token
    $ go build
## Run
To run in background

    $ ./bitcoinbot &

OR

    $ pkill bitcoinbot; cd /path/to/bitcoinbot; go build; nohup ./bitcoinbot &; disown; echo "\n\nBitcoinBot is up\n\n"; ps ax | grep ./bitcoinbot;

I prefer this block ^

## Usage in Discord
    >btc <currency>
    
OR

    @BitcoinBot#9430 <currency>
### Example

    @BitcoinBot#9430 USD
    
![Result](https://cdn.discordapp.com/attachments/364461767956365312/389885984789102595/unknown.png)

#### Help: `>btc help`

## [Invite](https://discordapp.com/api/oauth2/authorize?client_id=388984248062967819&permissions=2048&scope=bot)
Invite BitcoinBot to your Discord server.
##### Prerequisite: [Go](https://golang.org/doc/install) and [DiscordGo](https://github.com/bwmarrin/discordgo)
