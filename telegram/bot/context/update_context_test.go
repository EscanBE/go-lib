package context

import (
	"github.com/EscanBE/go-lib/telegram/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestNewTelegramUpdateContext(t *testing.T) {
	t.Run("init with correct input", func(t *testing.T) {
		update := tgbotapi.Update{}
		bot := bot.TelegramBot{}
		got := NewTelegramUpdateContext(update, bot)
		if got == nil {
			t.Errorf("expect success")
			return
		}
		if got.username != "" {
			t.Errorf("wrong default username")
			return
		}
	})
}

func TestTelegramUpdateContext_WithUsername(t *testing.T) {
	t.Run("with correct username", func(t *testing.T) {
		update := tgbotapi.Update{}
		bot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, bot)
		if ctx.username != "" {
			t.Errorf("wrong default username")
			return
		}
		ctx = ctx.WithUsername("1")
		if ctx.username != "1" {
			t.Errorf("expect set name correct")
			return
		}
	})
}

func TestTelegramUpdateContext_GetBot(t *testing.T) {
	t.Run("expect no error", func(t *testing.T) {
		update := tgbotapi.Update{}
		bot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, bot)
		if ctx == nil {
			t.Errorf("expect init")
			return
		}
		bot = ctx.GetBot()
	})
}

func TestTelegramUpdateContext_ExposeUpdate(t *testing.T) {
	t.Run("expect no error", func(t *testing.T) {
		update := tgbotapi.Update{}
		bot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, bot)
		if ctx == nil {
			t.Errorf("expect init")
			return
		}
		_ = ctx.ExposeUpdate()
	})
}
