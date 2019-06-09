package bot

import (
	"time"

	"github.com/omriza/teletaggies-bot/pkg/security"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Start() {
	token, err := security.ReadTokenFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.WithField("Username", b.Me.Username).Info("Connected")

	b.Handle(tb.OnText, textHandler(b))

	b.Start()
}
