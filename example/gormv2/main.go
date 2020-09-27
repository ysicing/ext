// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/kunnos/zap"
	"github.com/ysicing/ext/exlog/gormv2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	logger := gormv2.New(zap.L())
	logger.SetAsDefault()
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{Logger: logger})
	if err != nil {
		log.Println(err.Error())
	}
	db.AutoMigrate(&Product{})
	db.Debug().Create(&Product{Code: "D43", Price: 100})
	var product Product
	db.Debug().First(&product, 1)         // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	tx := db.Session(&gorm.Session{Logger: logger})
	tx.Find(&Product{})
}
