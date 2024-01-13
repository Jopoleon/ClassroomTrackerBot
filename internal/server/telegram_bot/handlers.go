package telegram_bot

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const unknownCommand = "Unknown command"

func (c *ClassroomTrackerBot) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	c.log.Info("Default handler use")
	if update.Message == nil {
		return
	}

	_, err := c.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   unknownCommand,
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

	assignments, err := c.classroomClient.GetClassroomAssignments(crs[0])
	if err != nil {
		c.log.Error("ListAssignments GetClassroomAssignments err ", err)
		return
	}

	var sb strings.Builder
	sb.WriteString("List of all Assignments: \n")
	for _, a := range assignments {
		sb.WriteString(a.Title + "\n")
	}

	_, err = c.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   sb.String(),
	})
	if err != nil {
		c.log.Error("ListAssignments b.SendMessage error: ", err)
	}
}

func GetUpdatesRepo(ctx context.Context, b *bot.Bot, update *models.Update) {
	panic("impelement me")
}
