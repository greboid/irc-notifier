package main

import (
	"testing"

	"github.com/greboid/irc-bot/v5/rpc"
)

func Test_checkHighlight(t *testing.T) {
	tests := []struct {
		name       string
		highlights []string
		message    *rpc.ChannelMessage
		want       bool
	}{
		{
			name:       "match",
			highlights: []string{"greg"},
			message: &rpc.ChannelMessage{
				Channel: "#greboid",
				Message: "greg is awesome",
				Source:  "test",
				Tags:    nil,
			},
			want: true,
		},
		{
			name:       "nothing to match",
			highlights: []string{},
			message: &rpc.ChannelMessage{
				Channel: "#greboid",
				Message: "greg is awesome",
				Source:  "test",
				Tags:    nil,
			},
			want: false,
		},
		{
			name:       "no match",
			highlights: []string{"rar"},
			message: &rpc.ChannelMessage{
				Channel: "#greboid",
				Message: "greg is awesome",
				Source:  "test",
				Tags:    nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkHighlight(tt.message, tt.highlights); got != tt.want {
				t.Errorf("checkHighlight() = %v, want %v", got, tt.want)
			}
		})
	}
}
