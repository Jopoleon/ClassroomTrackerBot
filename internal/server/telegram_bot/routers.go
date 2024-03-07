package telegram_bot

import "github.com/go-telegram/bot"

func (c *ClassroomTrackerBot) routers() {
	c.bot.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/list_assignments",
		bot.MatchTypeExact,
		c.ListAssignments)
}
