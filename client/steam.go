package client

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"go-alfred/output"
)

type SteamClient struct {
	Icon         string
	Host         string
	TrendingPath string
}

func (c *SteamClient) IconPath() string {
	return c.Icon
}

func (c *SteamClient) TrendingUrl() string {
	return c.Host + c.TrendingPath
}

func (c *SteamClient) TrendingFormat(body []byte) ([]output.UnifiedOutput, error) {
	res := make([]output.UnifiedOutput, 0)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// prefix := "https://search.bilibili.com/all?keyword="
	doc.Find(".search_result_row").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find(".title").Text()
		price := s.Find(".search_price").Text()
		price = strings.TrimSpace(price)
		href, _ := s.Attr("href")
		item := output.UnifiedOutput{
			Msg:      title,
			Detail:   price,
			Url:      href,
			IconPath: "resources/steam_logo.png",
		}
		res = append(res, item)
	})

	return res, nil
}
