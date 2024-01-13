package telegram_bot

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-telegram/bot"
	githubClassroom "github.com/jopoleon/ClassroomBot/internal/clients/github_classroom"
	"github.com/jopoleon/ClassroomBot/internal/config"
)

type ClassroomTrackerBot struct {
	ctx             context.Context
	bot             *bot.Bot
	classroomClient *githubClassroom.ClassroomClient
	log             *slog.Logger
}

func NewClassroomTrackerBot(cfg *config.Config) (*ClassroomTrackerBot, error) {
	ctb := ClassroomTrackerBot{
		ctx: context.Background(),
		// TODO: can be placed in a separate package
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	opts := ctb.Options()
	newBot, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	ctb.bot = newBot
	ctb.classroomClient = githubClassroom.NewClient(ctb.ctx, cfg.GithubToken)

	return &ctb, nil
}

// Start launches the bot
func (c *ClassroomTrackerBot) Start() {
	c.log.Info("Bot started")
	c.routers()

	c.bot.Start(c.ctx)
}
