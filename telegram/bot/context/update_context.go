package context

import (
	"github.com/EscanBE/go-lib/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramUpdateContext struct {
	bot      bot.TelegramBot
	update   tgbotapi.Update
	username string
}

func NewTelegramUpdateContext(update tgbotapi.Update, bot bot.TelegramBot) *TelegramUpdateContext {
	return &TelegramUpdateContext{
		bot:    bot,
		update: update,
	}
}

func (ctx *TelegramUpdateContext) WithUsername(username string) *TelegramUpdateContext {
	ctx.username = username
	return ctx
}

func (ctx TelegramUpdateContext) GetBot() bot.TelegramBot {
	return ctx.bot
}

func (ctx TelegramUpdateContext) GetCommand() string {
	return ctx.update.Message.Command()
}

func (ctx TelegramUpdateContext) GetCommandArg() string {
	return ctx.update.Message.CommandArguments()
}

func (ctx TelegramUpdateContext) GetUserId() int64 {
	return ctx.update.Message.From.ID
}

func (ctx TelegramUpdateContext) GetUsername() string {
	return ctx.username
}

func (ctx TelegramUpdateContext) HasUsername() bool {
	return len(ctx.username) > 0
}

func (ctx TelegramUpdateContext) GetChat() *tgbotapi.Chat {
	return ctx.update.Message.Chat
}

func (ctx TelegramUpdateContext) GetChatId() int64 {
	return ctx.update.Message.Chat.ID
}

func (ctx TelegramUpdateContext) NewResponseMessage(msgContent string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(ctx.update.Message.Chat.ID, msgContent)
}
