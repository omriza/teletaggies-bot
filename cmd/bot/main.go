package main

import (
	"github.com/omriza/teletaggies-bot/pkg/bot"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Starting...")

	bot.Start()
}
