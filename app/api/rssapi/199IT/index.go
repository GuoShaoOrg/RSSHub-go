package _199IT

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *controller) Get199ITIndex(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "199IT_INDEX"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://www.199it.com/newly"
	headers := getHeaders()
	rssData := dao.RSSFeed{
		Title:       "199it",
		Link:        apiUrl,
		ImageUrl:    "https://www.199it.com/favicon.ico",
		Tag:         []string{"互联网", "IT", "科技"},
		Description: "互联网数据资讯网-199IT",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(apiUrl); err == nil {
		rssItems := parseArticle(resp.ReadAllString())
		rssData.Items = rssItems
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "199IT_INDEX", rssStr)
	g.Redis().DoVar("EXPIRE", "199IT_INDEX", 60*10)
	_ = req.Response.WriteXmlExit(rssStr)
}
