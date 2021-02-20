// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type TGBotv1 struct {
	Bot *tgbotapi.BotAPI
}

func NewTGBot(token string, debug bool, api ...string) *TGBotv1 {
	endpoint := tgbotapi.APIEndpoint
	if len(api) != 0 && strings.HasPrefix(api[0], "https://") {
		endpoint = api[0]
	}
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, endpoint)
	if err != nil {
		return nil
	}
	bot.Debug = debug
	return &TGBotv1{Bot: bot}
}

func (bot TGBotv1) SendMsg(msg string, ischan bool, senduser ...interface{}) error {
	if ischan {
		tgmsg := tgbotapi.NewMessageToChannel(senduser[0].(string), msg)
		_, err := bot.Bot.Send(tgmsg)
		return err
	}
	tgmsg := tgbotapi.NewMessage(senduser[0].(int64), msg)
	_, err := bot.Bot.Send(tgmsg)
	return err
}

func (bot TGBotv1) SendFile(filepath string, senduser ...interface{}) error {
	tgfile := tgbotapi.NewDocumentUpload(senduser[0].(int64), filepath)
	_, err := bot.Bot.Send(tgfile)
	return err
}