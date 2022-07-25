package client

import (
	"go-alfred/output"
	"go-alfred/util"
)

var clientMap = map[string]TrendingClient{
	Bilibili: &BilibiliClient{
		Icon:         "resources/bilibili_logo.png",
		Host:         "https://api.bilibili.com",
		TrendingPath: "/x/web-interface/search/square?limit=50",
	},
	Weibo: &WeiboClient{
		Icon:         "resources/weibo_logo.png",
		Host:         "https://weibo.com",
		TrendingPath: "/ajax/side/hotSearch",
	},
}

func TrendingClientAll() map[string]TrendingClient {
	return clientMap
}

func TrendingClientFactory(clientType string) TrendingClient {
	return clientMap[clientType]
}

type TrendingClient interface {
	IconPath() string
	TrendingUrl() string
	TrendingFormat(body []byte) ([]output.UnifiedOutput, error)
}

func GetTrending(client TrendingClient) ([]output.UnifiedOutput, error) {
	body, err := util.HttpGet(client.TrendingUrl())
	if err != nil {
		return nil, err
	}
	return client.TrendingFormat(body)
}
