package feed

import (
	"rsshub/app/dao"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gorilla/feeds"
)

func GenerateRSS(data dao.RSSFeed, rsshubLink string) string {

	feed := &feeds.Feed{
		Title:       data.Title,
		Link:        &feeds.Link{Href: data.Link},
		Description: data.Description,
		Author:      &feeds.Author{Name: data.Author},
		Image:       &feeds.Image{Url: data.ImageUrl},
		Created:     gconv.Time(data.Created),
	}

	feed.Items = make([]*feeds.Item, 0)
	itemList := data.Items

	for _, item := range itemList {
		feedItem := feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Author:      &feeds.Author{Name: item.Author},
			Enclosure:   &feeds.Enclosure{Url: item.Thumbnail},
		}
		if item.Created != "" {
			if gtime.NewFromStr(item.Created) != nil {
				feedItem.Created = gtime.NewFromStr(item.Created).Time
			}
		} else {
			nowTime := gtime.Now().Format("Y-m-d H:i:s.u")
			if gtime.NewFromStr(nowTime) != nil {
				feedItem.Created = gtime.NewFromStr(nowTime).Time
			}
		}
		feed.Items = append(feed.Items, &feedItem)
	}

	if result, err := feed.ToRss(); err == nil {
		return result
	} else {
		return ""
	}
}

func GenerateDescription(imageLink, content string) (description string) {
	var imageHtml string
	var contentHtml string
	var htmlString string
	if imageLink != "" && content != "" {
		imageHtml = "<img src=" + imageLink + " style='width:100%' >"
		contentHtml = "<div style='margin-top: 8px' >" + content + "</div>"
		htmlString = "<meta name='referrer' content='no-referrer' /><div style='position: relative;text-align: left;'>" + imageHtml + contentHtml + "</div>"
	} else if imageLink != "" && content == "" {
		imageHtml = "<img src=" + imageLink + " style='width:100%' >"
		htmlString = "<div style='position: relative;text-align: left;'>" + imageHtml + "</div>"
	} else {
		contentHtml = "<div >" + content + "</div>"
		htmlString = "<div style='position: relative;text-align: left;'>" + contentHtml + "</div>"
	}
	description = htmlString
	return
}
