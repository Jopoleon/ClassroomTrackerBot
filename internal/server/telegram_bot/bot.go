package telegram_bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jopoleon/ClassroomBot/internal/clients/github_classroom"
	"github.com/sirupsen/logrus"
)

type ClassroomTrackerBot struct {
	ctx             context.Context
	bot             *bot.Bot
	classroomClient *github_classroom.ClassroomClient
	log             *logrus.Logger
}

func NewClassroomTrackerBot(botToken string, githubAccessToken string) (*ClassroomTrackerBot, error) {

	ctb := ClassroomTrackerBot{
		log: logrus.New(),
	}
	ctx := context.Background()
	opts := []bot.Option{
		bot.WithDefaultHandler(ctb.DefaultHandler),
	}
	b, err := bot.New(botToken, opts...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/list_assignments", bot.MatchTypeExact, ctb.ListAssignments)

	githubClient := github_classroom.NewClient(ctx, githubAccessToken)
	ctb.bot = b
	ctb.classroomClient = githubClient
	ctb.ctx = ctx
	return &ctb, nil
}

func (c *ClassroomTrackerBot) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	c.log.Info("Default handler use")
	if update.Message == nil {
		return
	}
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Unknown command",
	})
	if err != nil {
		c.log.Error("handler b.SendMessage error: ", err)
	}
}

func (c *ClassroomTrackerBot) ListAssignments(ctx context.Context, b *bot.Bot, update *models.Update) {
	c.log.Info("ListAssignments handler use")
	crs, err := c.classroomClient.ListClassrooms()
	if err != nil {
		c.log.Error("ListAssignments ListClassrooms err ", err)
		return
	}

	//pp.Println(crs)

	assignments, err := c.classroomClient.GetClassroomAssignments(crs[0])
	if err != nil {
		c.log.Error("ListAssignments GetClassroomAssignments err ", err)
		return
	}

	responseMessage := "List of all Assignments: \n"

	for _, a := range assignments {
		responseMessage = responseMessage + a.Title + "\n"
	}
	if update.Message == nil {
		return
	}
	//b.Chat
	//b.Men
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   responseMessage,
	})
	if err != nil {
		c.log.Error("ListAssignments b.SendMessage error: ", err)
	}
}

func (c *ClassroomTrackerBot) Start() {
	c.log.Info("Bot started")
	c.bot.Start(c.ctx)
}
