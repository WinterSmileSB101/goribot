# Goribot
一个分布式友好的轻量的 Golang 爬虫框架。

[完整文档 | Document](https://imagician.net/goribot/)

[![Codecov](https://img.shields.io/codecov/c/gh/zhshch2002/goribot)](https://codecov.io/gh/zhshch2002/goribot)
[![go-report](https://goreportcard.com/badge/github.com/zhshch2002/goribot)](https://goreportcard.com/report/github.com/zhshch2002/goribot)
![license](https://img.shields.io/github/license/zhshch2002/goribot)
![code-size](https://img.shields.io/github/languages/code-size/zhshch2002/goribot.svg)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fzhshch2002%2Fgoribot.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fzhshch2002%2Fgoribot?ref=badge_shield)
[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/zhshch2002/goribot)

## 🚀Feature
* 优雅的 API
* 整洁、有趣的文档
* 高速
* 友善的分布式支持
* 丰富的扩展支持
* 轻量，适于学习或快速开箱搭建

## 👜获取 Goribot
```sh
go get -u github.com/zhshch2002/goribot
```
::: tip
Goribot 包含一个历史开发版本，如果您需要使用过那个版本，请拉取 Tag 为 v0.0.1 版本。
:::

## ⚡建立你的第一个项目
```Go
package main

import (
	"fmt"
	"github.com/zhshch2002/goribot"
)

func main() {
	s := goribot.NewSpider()

	s.AddTask(
		goribot.GetReq("https://httpbin.org/get"),
		func(ctx *goribot.Context) {
			fmt.Println(ctx.Resp.Text)
			fmt.Println(ctx.Resp.Json("headers.User-Agent"))
		},
	)

	s.Run()
}
```

## 🎉完成
至此你已经可以使用 Goribot 了。更多内容请从 [开始使用](./get-start) 了解。

## 🙏感谢
按字母顺序排序。

* [ants](https://github.com/panjf2000/ants)
* [chardet](https://github.com/saintfish/chardet)
* [colly](https://github.com/gocolly/colly)
* [gjson](https://github.com/tidwall/gjson)
* [goquery](https://github.com/PuerkitoBio/goquery)
* [go-logging](https://github.com/op/go-logging)
* [go-redis](https://github.com/go-redis/redis)

万分感谢以上项目的帮助🙏。

## 📃TODO

* ~~分布式支持~~
* 扩展
  * Json、CVS 数据收集
  * site Limiter
  * 随机代理
  * 错误重试
  * 过滤响应码
* English Document