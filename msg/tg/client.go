// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ysicing/ext/logger"
	"strings"
)

type BotConfig struct {
	ApiProxy string
	BotToken string
	Debug    bool
}

type BotClinet struct {
	Client   *tgbotapi.BotAPI
	UserID   int64
	ChanName string
}

func NewBot(cfg *BotConfig) *tgbotapi.BotAPI {
	endpoint := tgbotapi.APIEndpoint
	if len(cfg.ApiProxy) != 0 && strings.HasPrefix(cfg.ApiProxy, "http") {
		endpoint = cfg.ApiProxy
	}
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(cfg.BotToken, endpoint)
	if err != nil {
		logger.Slog.Fatalf("create bot conn: ", err)
		return nil
	}
	bot.Debug = cfg.Debug
	return bot
}

func (bc *BotClinet) SendMsg(msg string) {
	if bc.UserID > 0 {
		tgumsg := tgbotapi.NewMessage(bc.UserID, msg)
		_, err := bc.Client.Send(tgumsg)
		if err != nil {
			logger.Slog.Errorf("send user %v msg err: %v", bc.UserID, err)
		}
	}
	if len(bc.ChanName) != 0 {
		tgchanmsg := tgbotapi.NewMessageToChannel(bc.ChanName, msg)
		_, err := bc.Client.Send(tgchanmsg)
		if err != nil {
			logger.Slog.Errorf("send chan %v msg err: %v", bc.ChanName, err)
		}
	}
}
