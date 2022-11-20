package bot

import (
	"fmt"
	"github.com/EscanBE/go-lib/logging"
	"github.com/EscanBE/go-lib/types"
	"github.com/EscanBE/go-lib/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

type TelegramBot struct {
	bot    *tgbotapi.BotAPI
	logger logging.Logger
}

func NewBot(telegramBotToken string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		bot: bot,
	}, nil
}

func (b *TelegramBot) WithLogger(logger logging.Logger) *TelegramBot {
	b.logger = logger
	return b
}

func (b *TelegramBot) EnableDebug(enable bool) *TelegramBot {
	b.bot.Debug = enable
	return b
}

func (b *TelegramBot) StopReceivingUpdates() {
	b.bot.StopReceivingUpdates()
}

func (b *TelegramBot) GetBotAPI() *tgbotapi.BotAPI {
	return b.bot
}

func (b *TelegramBot) GetBotUsername() string {
	return b.bot.Self.UserName
}

func (b *TelegramBot) GetUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *TelegramBot) Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error) {
	return b.bot.Send(chattable)
}

func (b *TelegramBot) SendMessage(msgContent string, chatId int64) (tgbotapi.Message, error) {
	return b.Send(tgbotapi.NewMessage(chatId, msgContent))
}

func (b *TelegramBot) SendMessageToMultipleChats(msgContent string, chatIds []int64, perUserDuration *time.Duration, extraLogTags ...interface{}) error {
	if len(msgContent) < 1 {
		return fmt.Errorf("message content is empty")
	}
	if len(chatIds) < 1 {
		return fmt.Errorf("input chat ID list is empty")
	}

	cntSent := 0
	errors := make(map[string]bool)
	for _, chatId := range chatIds {
		var cancellationToken *types.CancellationToken
		if perUserDuration == nil {
			cancellationToken = nil
		} else {
			token := types.NewCancellationTokenSourceWithTimeoutDuration(*perUserDuration).GetCancellationToken()
			cancellationToken = &token
		}

		for {
			msg := tgbotapi.NewMessage(chatId, msgContent)
			_, err := b.bot.Send(msg)

			if err != nil {
				b.logError(
					"failed to send Telegram message to user",
					append(extraLogTags, "chat-id", chatId, "error", err.Error()),
				)

				errors[fmt.Sprintf("\"%s\"", err.Error())] = true

				time.Sleep(300 * time.Millisecond)

				if cancellationToken == nil || cancellationToken.IsExpired() {
					break
				}
			} else {
				cntSent++
				time.Sleep(100 * time.Millisecond)
				break
			}
		}
	}

	if cntSent == len(chatIds) && len(errors) < 1 {
		b.logInfo(
			"successfully sent Telegram message to multiple chats",
			append(extraLogTags, "count-total", len(chatIds)),
		)
		return nil
	} else {
		b.logError(
			"failed to send Telegram message to multiple chats",
			append(extraLogTags, "sent", cntSent, "count-total", len(chatIds)),
		)

		var err error
		if len(errors) > 0 {
			err = fmt.Errorf("failed to send telegram message, errors: [%s], sent %d/%d", strings.Join(utils.GetKeys(errors), ", "), cntSent, len(chatIds))
		} else if cntSent < 1 {
			err = fmt.Errorf("failed to send any telegram message")
		} else {
			err = fmt.Errorf("failed to send telegram message, sent %d/%d", cntSent, len(chatIds))
		}

		return err
	}
}

func (b *TelegramBot) logInfo(msg string, keyVals []interface{}) {
	if b.logger == nil {
		return
	}
	b.logger.Info(msg, keyVals...)
}

func (b *TelegramBot) logError(msg string, keyVals []interface{}) {
	if b.logger == nil {
		return
	}
	b.logger.Error(msg, keyVals...)
}
