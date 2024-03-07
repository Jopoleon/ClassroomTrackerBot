package main

import (
	"flag"
	"log/slog"

	"github.com/jopoleon/ClassroomBot/internal/config"
	tgbot "github.com/jopoleon/ClassroomBot/internal/server/telegram_bot"
)

func main() {
	path := flag.String("cfg", "", "path with config file")
	flag.Parse()

	cfg := config.MustConfig(*path)
	tgBOT, err := tgbot.NewClassroomTrackerBot(cfg)
	if err != nil {
		slog.Error("ListClassrooms err ", err)
		return
	}
	tgBOT.Start()
}
