package client

import (
	"go-alfred/output"
	"go-alfred/util"
)

var clientMap = map[string]TrendingClient{
	Weibo: &WeiboClient{
		Icon:         "resources/weibo_logo.png",
		Host:         "https://weibo.com",
		TrendingPath: "/ajax/side/hotSearch",
	},
	Zhihu: &ZhihuClient{
		Icon:         "resources/zhihu_logo.png",
		Host:         "https://www.zhihu.com",
		TrendingPath: "/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true",
	},
	Steam: &SteamClient{
		Icon:         "resources/steam_logo.png",
		Host:         "https://store.steampowered.com",
		TrendingPath: "/search/results?force_infinite=1&os=win&filter=topsellers",
	},
	Bilibili: &BilibiliClient{
		Icon:         "resources/bilibili_logo.png",
		Host:         "https://api.bilibili.com",
		TrendingPath: "/x/web-interface/search/square?limit=50",
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
