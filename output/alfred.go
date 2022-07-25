package output

import (
	"encoding/json"
	"fmt"
)

func OutputToAlfred(outputs []UnifiedOutput) {
	alfred := Alfred{
		Items: make([]Item, 0, len(outputs)),
	}
	for _, output := range outputs {
		item := Item{
			Title:        output.Msg,
			Subtitle:     output.Detail,
			Arg:          output.Url,
			Autocomplete: output.Msg,
			Icon: Icon{
				Path: output.IconPath,
			},
		}
		alfred.Items = append(alfred.Items, item)
	}
	outputBytes, err := json.Marshal(alfred)
	if err != nil {
		return
	}
	fmt.Println(string(outputBytes))
}

type Alfred struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title        string `json:"title"`        // 大字
	Subtitle     string `json:"subtitle"`     // 下面小字
	Arg          string `json:"arg"`          // 到下一级的输出，比如回车复制的数据就是这条。这里可以填写url，回车之后打开链接
	Autocomplete string `json:"autocomplete"` // tab自动填充，比如"hot bi"自动填充为"hot bilibili"
	Icon         Icon   `json:"icon"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
