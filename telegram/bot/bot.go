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

// TelegramBot wraps the bot and provide some utilities
type TelegramBot struct {
	bot    *tgbotapi.BotAPI
	logger logging.Logger
}

// NewBot returns a new instance of TelegramBot, provide some utilities
//goland:noinspection GoUnusedExportedFunction
func NewBot(telegramBotToken string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		bot: bot,
	}, nil
}

// WithLogger injects a logger into TelegramBot, enable bot to be able to logging
func (b *TelegramBot) WithLogger(logger logging.Logger) *TelegramBot {
	b.logger = logger
	return b
}

// EnableDebug enables debugging mode of bot
func (b *TelegramBot) EnableDebug(enable bool) *TelegramBot {
	b.bot.Debug = enable
	return b
}

// StopReceivingUpdates stops the go routine which receives updates
func (b *TelegramBot) StopReceivingUpdates() {
	b.bot.StopReceivingUpdates()
}

// ExposeBotAPI exposes the underlying tgbotapi.BotAPI instance
func (b *TelegramBot) ExposeBotAPI() *tgbotapi.BotAPI {
	return b.bot
}

// GetBotUsername returns username of the Telegram bot
func (b *TelegramBot) GetBotUsername() string {
	return b.bot.Self.UserName
}

// GetUpdatesChannel returns default tgbotapi.UpdatesChannel for the tgbotapi.BotAPI, timeout is 60, offset is 0
func (b *TelegramBot) GetUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

// Send delivers a chat
func (b *TelegramBot) Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error) {
	return b.bot.Send(chattable)
}

// SendMessage delivers a message to destination chat id
func (b *TelegramBot) SendMessage(msgContent string, chatId int64) (tgbotapi.Message, error) {
	return b.Send(tgbotapi.NewMessage(chatId, msgContent))
}

// SendMessageToMultipleChats delivers a message to multiple chats
func (b *TelegramBot) SendMessageToMultipleChats(msgContent string, chatIds []int64, perUserDuration *time.Duration, extraLogTags ...interface{}) error {
	if len(msgContent) < 1 {
		return fmt.Errorf("message content is empty")
	}
	if len(chatIds) < 1 {
		return fmt.Errorf("input chat ID list is empty")
	}

	chatIds = utils.GetUniqueElements(chatIds...)

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

// logDebug uses the supplied logger to perform logging at Debug level
func (b *TelegramBot) logDebug(msg string, keyVals []interface{}) {
	if b.logger == nil {
		return
	}
	b.logger.Debug(msg, keyVals...)
}

// logInfo uses the supplied logger to perform logging at Info level
func (b *TelegramBot) logInfo(msg string, keyVals []interface{}) {
	if b.logger == nil {
		return
	}
	b.logger.Info(msg, keyVals...)
}

// logError uses the supplied logger to perform logging at Error level
func (b *TelegramBot) logError(msg string, keyVals []interface{}) {
	if b.logger == nil {
		return
	}
	b.logger.Error(msg, keyVals...)
}
