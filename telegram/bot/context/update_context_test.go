package context

import (
	"fmt"
	"github.com/EscanBE/go-lib/telegram/bot"
	"github.com/EscanBE/go-lib/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strings"
	"testing"
)

func TestNewTelegramUpdateContext(t *testing.T) {
	t.Run("init with correct input", func(t *testing.T) {
		update := tgbotapi.Update{}
		tBot := bot.TelegramBot{}
		got := NewTelegramUpdateContext(update, tBot)
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
		tBot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, tBot)
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
		tBot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, tBot)
		if ctx == nil {
			t.Errorf("expect init")
			return
		}
		tBot = ctx.GetBot()
	})
}

func TestTelegramUpdateContext_ExposeUpdate(t *testing.T) {
	t.Run("expect no error", func(t *testing.T) {
		update := tgbotapi.Update{}
		tBot := bot.TelegramBot{}
		ctx := NewTelegramUpdateContext(update, tBot)
		if ctx == nil {
			t.Errorf("expect init")
			return
		}
		_ = ctx.ExposeUpdate()
	})
}

func TestTelegramUpdateContext_GetUserId(t *testing.T) {
	t.Run("get correct id", func(t *testing.T) {
		uid := rand.Int63()
		ctx := TelegramUpdateContext{
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					From: &tgbotapi.User{
						ID: uid,
					},
				},
			},
		}

		if ctx.GetUserId() != uid {
			t.Errorf("GetUserId() returned wrong value %d, want %d", ctx.GetUserId(), uid)
		}
	})
}

func TestTelegramUpdateContext_GetUsername(t *testing.T) {
	for i := 1; i <= 3; i++ {
		name := fmt.Sprintf("test-%d", i)
		t.Run(name, func(t *testing.T) {
			ctx := TelegramUpdateContext{}
			ctx = *ctx.WithUsername(name)
			if got := ctx.GetUsername(); got != name {
				t.Errorf("GetUsername() = %v, want %v", got, name)
			}
		})
	}
}

func TestTelegramUpdateContext_HasUsername(t *testing.T) {
	t.Run("init empty", func(t *testing.T) {
		ctx := TelegramUpdateContext{
			username: "",
		}
		if ctx.HasUsername() {
			t.Errorf("HasUsername() = true, want false")
			return
		}
		ctx.WithUsername("x")
		if !ctx.HasUsername() {
			t.Errorf("HasUsername() = false, want true")
			return
		}
	})

	t.Run("init with value", func(t *testing.T) {
		ctx := TelegramUpdateContext{
			username: "x",
		}
		if !ctx.HasUsername() {
			t.Errorf("HasUsername() = false, want true")
			return
		}
	})
}

func TestTelegramUpdateContext_GetChat(t *testing.T) {
	t.Run("get correct chat", func(t *testing.T) {
		chat := &tgbotapi.Chat{}
		ctx := TelegramUpdateContext{
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: chat,
				},
			},
		}

		if ctx.GetChat() != chat {
			t.Errorf("GetChat() returned wrong value %v, want %v", ctx.GetChat(), chat)
		}
	})
}

func TestTelegramUpdateContext_GetChatId(t *testing.T) {
	t.Run("get correct chat id", func(t *testing.T) {
		cid := rand.Int63()
		chat := &tgbotapi.Chat{
			ID: cid,
		}
		ctx := TelegramUpdateContext{
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: chat,
				},
			},
		}

		if ctx.GetChatId() != cid {
			t.Errorf("GetChatId() returned wrong value %d, want %d", ctx.GetChatId(), cid)
		}
	})
}

func TestTelegramUpdateContext_NewResponseMessage(t *testing.T) {
	t.Run("for correct chat id", func(t *testing.T) {
		cid := rand.Int63()
		content := "test"

		chat := &tgbotapi.Chat{
			ID: cid,
		}
		ctx := TelegramUpdateContext{
			update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Chat: chat,
				},
			},
		}
		chattable := ctx.NewResponseMessage(content)
		if m, ok := chattable.(tgbotapi.MessageConfig); ok {
			if m.ChatID != cid {
				t.Errorf("wrong chat id %d, want %d", m.ChatID, cid)
				return
			}

			if m.Text != content {
				t.Errorf("wrong message content [%s], want [%s]", m.Text, content)
			}
		} else {
			t.Errorf("wrong Chattable type, expect %T, found %T", tgbotapi.MessageConfig{}, chattable)
			return
		}
	})
}

func setUpTestCmdUpdate(command string) TelegramUpdateContext {
	cid := rand.Int63()
	uid := rand.Int63()
	isCommand := len(command) > 0
	return TelegramUpdateContext{
		update: tgbotapi.Update{
			Message: &tgbotapi.Message{
				From: &tgbotapi.User{
					ID: uid,
				},
				Chat: &tgbotapi.Chat{
					ID: cid,
				},
				Text: utils.ConditionalString(isCommand, fmt.Sprintf("/%s", command), "hello world"),
				Entities: []tgbotapi.MessageEntity{
					{
						Type:   utils.ConditionalString(isCommand, "bot_command", "not_command"),
						Offset: 0,
						Length: func() int {
							spaceIdx := strings.Index(command, " ")
							if spaceIdx < 0 {
								return len(command) + 1
							}
							return len(command[:strings.Index(command, " ")]) + 1
						}(),
					},
				},
			},
		},
	}
}

func TestTelegramUpdateContext_GetCommand(t *testing.T) {
	tests := []string{"help", "version", "me", ""}
	for _, cmd := range tests {
		t.Run(cmd, func(t *testing.T) {
			ctx := setUpTestCmdUpdate(cmd)
			if got := ctx.GetCommand(); got != cmd {
				t.Errorf("GetCommand() = %v, want %v", got, cmd)
			}
		})
	}
}

func TestTelegramUpdateContext_GetCommandArg(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "help me",
			want:  "me",
		},
		{
			input: "help",
			want:  "",
		},
		{
			input: "set 1",
			want:  "1",
		},
		{
			input: "",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ctx := setUpTestCmdUpdate(tt.input)
			if got := ctx.GetCommandArg(); got != tt.want {
				t.Errorf("GetCommandArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
