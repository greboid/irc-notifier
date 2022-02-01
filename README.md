## Igloo IRC notifications

Plugin for [IRC-Bot](https://github.com/greboid/irc-bot)

Monitors channel messages for certain words, if found sends a push notification via IglooIRC

#### Configuration

At a bare minimum you also need to give it a list of highlight words and an Igloo IRC Token and an RPC token.  
You'll like also want to specify the bot host.

#### Example running

```
---
version: "3.5"
service:
  bot-github:
    image: greboid/irc-github
    environment:
      RPC_HOST: bot
      RPC_TOKEN: <as configured on the bot>
      HIGHLIGHT_WORDS: greboid
      IGLOO_TOKEN: pushtokenfromtheclient
```

```
github -rpc-host bot -rpc-token <as configured on the bot> -highlight-words greboid -igloo-token pushtokenfromtheclient
```
