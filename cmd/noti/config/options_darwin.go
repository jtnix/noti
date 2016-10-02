package config

import (
	"github.com/variadico/noti/nsuser"
	"github.com/variadico/noti/say"
	"github.com/variadico/noti/slack"
)

type Options struct {
	DefaultSet []string
	Banner     *nsuser.Notification
	Speech     *say.Notification
	Slack      *slack.Notification
}

func NewOptions() Options {
	return Options{
		DefaultSet: make([]string, 0),
		Banner:     new(nsuser.Notification),
		Speech:     new(say.Notification),
		Slack:      new(slack.Notification),
	}
}
