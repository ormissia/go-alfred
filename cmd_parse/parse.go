package cmd_parse

import (
	"go-alfred/client"
	"go-alfred/output"
)

func Execute(cmd string) {
	if c := client.TrendingClientFactory(cmd); c != nil {
		trending, err := client.GetTrending(c)
		if err != nil {
			// TODO 改为输出错误
			autocomplete()
			return
		}
		output.OutputToAlfred(trending)
		return
	}
	autocomplete()
	return
}
func autocomplete() {
	trendingMap := client.TrendingClientAll()
	outputs := make([]output.UnifiedOutput, 0, len(trendingMap))
	// TODO 排序
	outputs = append(outputs, output.UnifiedOutput{
		Msg:      client.Bilibili,
		Detail:   client.Bilibili,
		Url:      "",
		IconPath: trendingMap[client.Bilibili].IconPath(),
	})
	outputs = append(outputs, output.UnifiedOutput{
		Msg:      client.Weibo,
		Detail:   client.Weibo,
		Url:      "",
		IconPath: trendingMap[client.Weibo].IconPath(),
	})
	output.OutputToAlfred(outputs)
}
