# Ibot

Based on [slackbot](https://github.com/trinchan/slackbot), but completely reworked and functionality ripped out.



Dependencies
============
Schema  - `go get github.com/gorilla/schema`

Installation
============
You can grab the source code using `go get github.com/trinchan/slackbot` and install like usual. `go install github.com/trinchan/slackbot`

Setup
=====
###Setup an Incoming Webhook
If you don't already have an [Incoming Webhook](https://my.slack.com/services/new/incoming-webhook) setup for your bot, you'll want to start here.  Set one up (make sure it's enabled) and don't be afraid to read through the setup instructions.  You're after the "Webhook URL" that slack generates for you.  At the end of the URL, is the token that slack uses to authenticate where the URL is coming from:
```
https://hooks.slack.com/services/AAAAAAAAA/BBBBBBBBB/YourSecretToken123456789
```
That's really all you care about right now.  You can set the default Icon, Name, and default Channel, but slack will let you override that information in the http requests you send.  So don't worry yourself with setting everything up.  You just need the token.

###Create a Domain Configuration File
Assuming you've already pulled the source, and successfully compiled/installed, you should have a `slackbot` executable in your `$GOPATH/bin`.  You need to create a file named `config.json` and give your bot the proper credentials to send messages to your slack server.  Feel free to place the file in the a sub-folder if you want to be all organized like that.  If you want to attach more than one slack server to your bot, you can simply add another entry under "domain_tokens".

The config file (config.json) has the following format:

```json
{
    "port": PORT_FOR_BOT,
    "domain_tokens": {
        "YOUR_SLACK_DOMAIN":       "YOUR_SLACK_INCOMING_WEBHOOK_TOKEN",
        "YOUR_OTHER_SLACK_DOMAIN": "MATCHING_INCOMING_WEBHOOK_TOKEN"
    }
}
```
Note that the last "domain_token" does NOT have a comma at the end of the line (but the others do)

###Send messages to your bot
This framework can respond to "slash commands" and "outgoing webhooks"  If you want users to be able to silently type `/ping`, and have the ping-bot respond in their channel, then you'll want to set up "slash commands".  Each bot will need it's own command setup.  The other option is to configure an outgoing webhook with a symbol for the prefix. Exe: `!ping`.  This option only requires one configuration, but the commands will be entered into the channel as regular messages.

#####Configuring an Outgoing Webhook
I use an [Outgoing Webhook](https://my.slack.com/services/new/outgoing-webhook)

1. Add a new Outgoing Webhook Integration.
2. Here, you can tell slack to ONLY pay attention to a specific channel, or to simply listen to all public channels.  Outgoing Webhooks can not listen to private channels/direct messages.
3. The {trigger_word} should only be one character (preferrably a symbol, such as ! or ?) and typing `{trigger_word}ping` will trigger the Ping bot.
TODO: Clean up the trigger_word configuration.  Maybe something can be added to the config?
4. The URL should follow the following format: `your_address.com:port/slack_hook` (no trailing /)
No other configuration is necessary.

The bot will respond to commands of the form `{trigger_word}bot param param param` in the specified channels
#####Configuring Slash Commands
Alternativly, each bot you make can respond to a corresponding [Slash Command](https://my.slack.com/services/new/slash-commands).

1. Add a new slash command, use the [bot's name](https://github.com/trinchan/slackbot/tree/master/robots) as the name of the command.
2. The URL should follow the following format: `your_address.com:port/slack` (no trailing /)
3. You want to use POST.
4. This bot does not currently pay attention to the payload's token.
TODO: Pay attention to the payload's token.
5. Repeat for each bot you want to enable.

The bot will respond to commands of the form `/bot param param param`

## Running

`./ibot -token <domain:token>`
