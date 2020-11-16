// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/msg/tg"
	"github.com/ysicing/ext/utils/extime"
)

func main() {
	tgcfg := tg.BotConfig{
		ApiProxy: "https://botapi.hk2.godu.dev/bot%s/%s",
		BotToken: "botxxx",
		Debug:    false,
	}
	tgbot := tg.BotClinet{
		Client:   tg.NewBot(&tgcfg),
		UserID:   12306,
		ChanName: "",
	}
	tgbot.SendMsg(extime.GetToday())
}
