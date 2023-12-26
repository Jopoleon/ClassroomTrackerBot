package telegram

import (
	"context"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func StartBot(botToken string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		//bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/list_assignments", bot.MatchTypeExact, handler)
	//b.RegisterHandler(bot.HandlerTypeMessageText, "tttt", bot.MatchTypeContains, handler)
	logrus.Info("bot started")
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message == nil {
		return
	}

	exapleResponse := `
		List of classrooms: 
			1) qqqqq 
			2) tttt
	`

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   exapleResponse,
	})
	if err != nil {
		logrus.Error("handler b.SendMessage error: ", err)
	}
	//pp.Println(msg.From.FirstName + " " + msg.From.LastName + " ::: " + msg.From.Username)
	pp.Println(msg.Chat.FirstName + " " + msg.Chat.LastName + " ::: " + msg.Chat.Username)
	pp.Println(msg.Text)
	//pp.Println(msg.Chat.ID)

}
