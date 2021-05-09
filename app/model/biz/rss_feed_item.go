package biz

import "github.com/gogf/gf/os/gtime"

type RssFeedItemData struct {
	Id              string
	ChannelId       string
	Title           string
	ChannelDesc     string
	Link            string
	Date            *gtime.Time
	Author          string
	InputDate       *gtime.Time
	ChannelImageUrl string
	ChannelTitle    string
}