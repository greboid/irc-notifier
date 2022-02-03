## IRC notifier

Plugin for [IRC-Bot](https://github.com/greboid/irc-bot)

Monitors incoming messages for certain words, if found sends a push notification via IglooIRC

#### Configuration

`RPC_HOST` [Required]  
The hostname of the bot to connect to

`RPC_TOKEN` [Required]  
The authorisation token to send to the bot

`NETWORK` [Optional]   
Name of the network, used as a title in notifications.

`HIGHLIGHT_WORDS` [Required]  
Comma separated list of words to match messages on, if found a notification will be sent.

`IGLOO_TOKEN` [Required]  
Igloo IRC Push Token, used to authorise messages sent to the push notification service.  This can be found in the Settings of the app under `Notifications (Authorized)` near the bottom.

All of these settings are available as CLI flags, lowercase the settings and replace underscores with hyphens, for example RPC_HOST would be -rpc-host

#### Example compose file

```
---
services:
  bot-notifier:
    image: ghcr.io/greboid/irc-notifier
    environment:
      RPC_HOST: bot
      RPC_TOKEN: <as configured on the bot>
      HIGHLIGHT_WORDS: greboid
      NETWORK: CoolNetwork
      IGLOO_TOKEN: pushtokenfromtheclient
```
