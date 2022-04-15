package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ergochat/irc-go/ircmsg"
	"github.com/greboid/golog"
	"github.com/greboid/irc-bot/v5/plugins"
	"github.com/greboid/irc-bot/v5/rpc"
	"github.com/kouhin/envflag"
	"go.uber.org/zap"
)

var (
	RPCHost        = flag.String("rpc-host", "localhost", "gRPC server to connect to")
	RPCPort        = flag.Int("rpc-port", 8001, "gRPC server port")
	RPCToken       = flag.String("rpc-token", "", "gRPC authentication token")
	Debug          = flag.Bool("debug", false, "Show debugging info")
	HighlightWords = flag.String("highlight-words", "", "Comma separated highlighted words")
	Network        = flag.String("network", "", "Network to show in title of push notification")
	IglooPushToken = flag.String("igloo-token", "", "Igloo IRC Push Token - Found in client settings")
	log            *zap.SugaredLogger
	helper         *plugins.PluginHelper
)

type HighlightHandler struct {
	Highlights []string
}

func main() {
	log.Infof("Starting notifier plugin")
	err := envflag.Parse()
	if err != nil {
		log.Fatalf("Unable to load config: %s", err.Error())
		return
	}
	log = logger.MustCreateLogger(*Debug)
	helper, err = plugins.NewHelper(fmt.Sprintf("%s:%d", *RPCHost, uint16(*RPCPort)), *RPCToken)
	if err != nil {
		log.Fatalf("Unable to create plugin helper: %s", err.Error())
		return
	}
	handler := HighlightHandler{
		Highlights: parseHighlights(*HighlightWords),
	}
	err = helper.RegisterChannelMessageHandler("*", handler.handleChannelMessage)
	if err != nil {
		log.Fatalf("Error registering channel handler: %s", err.Error())
		return
	}
	log.Infof("Exiting")
}

func (h *HighlightHandler) handleChannelMessage(message *rpc.ChannelMessage) {
	if checkHighlight(message, h.Highlights) {
		nuh, err := ircmsg.ParseNUH(message.Source)
		if err != nil {
			sendNotification(*Network, message.Channel, message.Message, nuh.Name)
		}
	}
}

func sendNotification(network, channel, message, sender string) {
	params := url.Values{}
	params.Set("network", network)
	params.Set("channel", channel)
	params.Set("message", message)
	params.Set("sender", sender)
	params.Set("type", "")
	params.Set("device1", *IglooPushToken)
	encoded := params.Encode()
	req, err := http.NewRequest(http.MethodPost, "https://api.iglooirc.com/znc/push", strings.NewReader(encoded))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "ZNC Push/irc-notifier")

	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
	if err != nil {
		log.Errorf("Unable to send notification: %s", err.Error())
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending notification: %s", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading notification response: %s", err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		log.Errorf("Notification server error: %d: %s", resp.StatusCode, body)
	}
}

func checkHighlight(message *rpc.ChannelMessage, highlights []string) bool {
	for index := range highlights {
		if strings.Contains(strings.ToLower(message.Message), highlights[index]) {
			return true
		}
	}
	return false
}

func parseHighlights(users string) []string {
	highlights := make([]string, 0)
	splitHighlights := strings.Split(users, ",")
	for _, highlight := range splitHighlights {
		trimmedHighlight := strings.TrimSpace(strings.ToLower(highlight))
		if trimmedHighlight != "" {
			highlights = append(highlights, trimmedHighlight)
		}
	}
	return highlights
}
