package telegram_bot

import "github.com/go-telegram/bot"

func (c *ClassroomTrackerBot) Options() []bot.Option {
	return []bot.Option{
		bot.WithDefaultHandler(c.DefaultHandler),
	}
}
