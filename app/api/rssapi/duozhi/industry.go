package duozhi

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetIndustryNews(req *ghttp.Request) {
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getIndustryNewsLinks()[linkType]

	cacheKey := "DUOZHI_INDUSTRY_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.duozhi.com/industry/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "多知 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "多知网 - 独立商业视角 新锐教育观察",
		ImageUrl:    "http://www.duozhi.com/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		rssData.Items = commonParser(resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*6)
	_ = req.Response.WriteXmlExit(rssStr)
}
