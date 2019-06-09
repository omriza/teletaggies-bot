package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/omriza/teletaggies-bot/pkg/tagging"

	tb "gopkg.in/tucnak/telebot.v2"
)

func textHandler(b *tb.Bot) func(*tb.Message) {

	return func(msg *tb.Message) {
		opts := tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		}

		_ = b.Notify(msg.Sender, tb.Typing)

		linkTagger := tagging.LinkTagger{}
		tags, err := linkTagger.Tag(msg.Text)

		var txt string
		if err != nil {
			txt = "Couldn't extract links mate"
		} else {
			txt = "Tags: " + strings.Trim(strings.Replace(fmt.Sprint(tags), " ", ", ", -1), "[]")
		}

		log.Println(txt)
		s := fmt.Sprintf("```%s```\n%s", msg.Text, txt)
		_, _ = b.Send(msg.Sender, s, &opts)
	}
}
