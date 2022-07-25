package client

import (
	"encoding/json"
	"net/url"

	"go-alfred/output"
)

type BilibiliClient struct {
	Icon         string
	Host         string
	TrendingPath string
}

func (c *BilibiliClient) IconPath() string {
	return c.Icon
}

func (c *BilibiliClient) TrendingUrl() string {
	return c.Host + c.TrendingPath
}

func (c *BilibiliClient) TrendingFormat(body []byte) ([]output.UnifiedOutput, error) {
	res := make([]output.UnifiedOutput, 0)

	resp := new(BilibiliResp)
	if err := json.Unmarshal(body, resp); err != nil {
		return res, err
	}

	// TODO 判断返回code

	// resp.Data.Trending.List = resp.Data.Trending.List[:10]
	prefix := "https://search.bilibili.com/all?keyword="
	for _, data := range resp.Data.Trending.List {
		item := output.UnifiedOutput{
			Msg:      data.Keyword,
			Detail:   data.ShowName,
			Url:      prefix + url.PathEscape(data.Keyword),
			IconPath: "resources/bilibili_logo.png",
		}
		res = append(res, item)
	}
	return res, nil
}

type BilibiliResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Trending struct {
			Title   string `json:"title"`
			Trackid string `json:"trackid"`
			List    []struct {
				Keyword  string `json:"keyword"`
				ShowName string `json:"show_name"`
				Icon     string `json:"icon"`
				Uri      string `json:"uri"`
				Goto     string `json:"goto"`
			} `json:"list"`
		} `json:"trending"`
	} `json:"data"`
}
