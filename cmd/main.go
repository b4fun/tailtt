package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/b4fun/tailtt/notification"
	"github.com/hpcloud/tail"
)

var errNoNotifier = errors.New("no notifier set")

type keywordPatterns []*regexp.Regexp

func (k keywordPatterns) String() string {
	return fmt.Sprintf("%v", []*regexp.Regexp(k))
}

func (k *keywordPatterns) Set(r string) error {
	pattern, err := regexp.Compile(r)
	if err != nil {
		return err
	}

	*k = append(*k, pattern)

	return nil
}

func (k keywordPatterns) IsEmpty() bool {
	return len(k) < 1
}

func (k keywordPatterns) MatchLine(line string) bool {
	for _, p := range k {
		if p.MatchString(line) {
			return true
		}
	}
	return false
}

func main() {
	file := flag.String("f", "", "file to tail (taiil -f)")

	keywords := &keywordPatterns{}
	flag.Var(keywords, "k", "keyword patterns (required)")

	notifyBearychatRTMToken := flag.String("notify-bearychat-rtm-token", "", "bearychat rtm token")
	notifyBearychatRTMChannel := flag.String("notify-bearychat-channel", "", "bearychat notification channel")

	flag.Parse()

	if *file == "" {
		flag.Usage()
		abort()
	}
	if keywords.IsEmpty() {
		flag.Usage()
		abort()
	}

	var notifies []notification.Notifier
	{
		token := *notifyBearychatRTMToken
		channel := *notifyBearychatRTMChannel
		if token != "" && channel != "" {
			notifies = append(notifies, notification.MustNewBearychatRTM(token, channel))
		}
	}
	if len(notifies) < 1 {
		abort(errNoNotifier)
	}
	notifier := notification.NewNotifier(notifies)

	t, err := tail.TailFile(*file, tail.Config{Follow: true})
	if err != nil {
		abort(err)
	}

	for line := range t.Lines {
		if keywords.MatchLine(line.Text) {
			ctx := context.Background()
			if err := notifier.Notify(ctx, line.Text); err != nil {
				abort(err)
			}
		}
	}
}

func abort(errs ...error) {
	if len(errs) > 0 {
		fmt.Println(errs[0])
	}
	os.Exit(1)
}
