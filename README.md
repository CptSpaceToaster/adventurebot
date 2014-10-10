adventurebot
===========

A silly text based adventure that gloriously mangled the framework to build bots for Slack.... thank you [trinchan](https://github.com/trinchan/slackbot)

Dependencies
============
Schema  - `go get github.com/gorilla/schema`

Installation
============
Is this your first project in GO?  You might find [this introduction](https://github.com/CptSpaceToaster/adventurebot/blob/master/INSTALLATION_NOTES.md) useful

`go get github.com/cptspacetoaster/adventurebot`  

Setup
=====
Create a config file (config.json) in `$GOPATH/bin` with the following format:

```
{
    "domain": "{YOUR_SLACK_DOMAIN}",
    "port": {PORT_FOR_BOT},
    "token": "{YOUR_SLACK_INCOMING_WEBHOOK_TOKEN}"
}
```

Make sure you have [Incoming Webhooks](https://slack.com/services/new/incoming-webhook) enabled and you are using that integration token for your config.

Adventurebot will respond to an [Outgoing Webhook](https://slack.com/services/new/outgoing-webhook). (hard set for patterns that begin with question-mark's at the moment)

TODO: Lots

Adding Rooms, and other configurations
===========

TODO:
copy the blank<thing>.json and fill in the fields... more to come, and I'm not completely finished with the formatting at the moment.

Running the Executable
=======
`cd $GOPATH/bin`  
`./adventurebot`
