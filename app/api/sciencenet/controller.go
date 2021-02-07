package sciencenet

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gcharset"
	"rsshub/app/dao"
	"rsshub/lib"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(respString string) (items []dao.RSSItem) {
	respString, _ = gcharset.Convert("UTF-8", "gbk", respString)
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.Find("table", "class", "tablebox").FindAll("tr")
	baseUrl := "http://blog.sciencenet.cn/"
	for index, article := range articleList {
		if index == 0 {
			continue
		}
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string
		aTagList := article.FindAll("a")
		for _, aTag := range aTagList {
			if aTag.Attrs()["title"] != "" {
				title, _ = aTag.Attrs()["title"]
				link = baseUrl + aTag.Attrs()["href"]
			} else {
				author = aTag.Text()
			}
		}

		tdDocList := article.FindAll("td")
		lastTdDoc := tdDocList[len(tdDocList)-1]
		time = lastTdDoc.Text()

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: lib.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "recommend",
			Title:     "精选博文"},
		"hot": {
			ChannelId: "hot",
			Title:     "热门博文"},
		"new": {
			ChannelId: "new",
			Title:     "最新博文"},
	}

	return Links
}
