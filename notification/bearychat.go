package notification

import (
	"context"
	"fmt"

	"github.com/bearyinnovative/bearychat-go"
)

// BearychatRTM implements bearychat notification.
type BearychatRTM struct {
	rtmClient *bearychat.RTMClient
	channel   string
}

func MustNewBearychatRTM(rtmToken, channel string) *BearychatRTM {
	rtmClient, err := bearychat.NewRTMClient(rtmToken)
	if err != nil {
		panic(err)
	}

	return &BearychatRTM{
		rtmClient: rtmClient,
		channel:   channel,
	}
}

func (b BearychatRTM) Notify(ctx context.Context, line string) error {
	return b.rtmClient.Incoming(bearychat.RTMIncoming{
		Text:       fmt.Sprintf("**line found**: %s", line),
		VChannelId: b.channel,
		Markdown:   true,
	})
}
