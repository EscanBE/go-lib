package context

import (
	"github.com/EscanBE/go-lib/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramUpdateContext hold update context when received an update, this struct provides some utilities
type TelegramUpdateContext struct {
	bot      bot.TelegramBot
	update   tgbotapi.Update
	username string
}

// NewTelegramUpdateContext wraps the new update thus can perform some utilities
func NewTelegramUpdateContext(update tgbotapi.Update, bot bot.TelegramBot) *TelegramUpdateContext {
	return &TelegramUpdateContext{
		bot:    bot,
		update: update,
	}
}

// WithUsername set the input username to the context. Because username does not include within the update instance thus manual provision is needed
func (ctx *TelegramUpdateContext) WithUsername(username string) *TelegramUpdateContext {
	ctx.username = username
	return ctx
}

// GetBot returns the underlying bot.TelegramBot instance
func (ctx TelegramUpdateContext) GetBot() bot.TelegramBot {
	return ctx.bot
}

// ExposeUpdate exposes the underlying tgbotapi.Update instance
func (ctx TelegramUpdateContext) ExposeUpdate() tgbotapi.Update {
	return ctx.update
}

// GetCommand returns command as string if this update is a command message
func (ctx TelegramUpdateContext) GetCommand() string {
	if !ctx.update.Message.IsCommand() {
		return ""
	}
	return ctx.update.Message.Command()
}

// GetCommandArg returns command argument as string if this update is a command message
func (ctx TelegramUpdateContext) GetCommandArg() string {
	if !ctx.update.Message.IsCommand() {
		return ""
	}
	return ctx.update.Message.CommandArguments()
}

// GetUserId returns the sender user id
func (ctx TelegramUpdateContext) GetUserId() int64 {
	return ctx.update.Message.From.ID
}

// GetUsername returns the username which was set using WithUsername method
func (ctx TelegramUpdateContext) GetUsername() string {
	return ctx.username
}

// HasUsername returns true if any non-empty username was set using WithUsername method
func (ctx TelegramUpdateContext) HasUsername() bool {
	return len(ctx.username) > 0
}

// GetChat returns the underlying chat info, which was the update came from
func (ctx TelegramUpdateContext) GetChat() *tgbotapi.Chat {
	return ctx.update.Message.Chat
}

// GetChatId returns the underlying chat id, which was the update came from
func (ctx TelegramUpdateContext) GetChatId() int64 {
	return ctx.update.Message.Chat.ID
}

// NewResponseMessage initializes a response message based in the chat which the update came from
func (ctx TelegramUpdateContext) NewResponseMessage(msgContent string) tgbotapi.Chattable {
	return tgbotapi.NewMessage(ctx.update.Message.Chat.ID, msgContent)
}
