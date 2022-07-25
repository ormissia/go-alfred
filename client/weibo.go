package client

import (
	"encoding/json"
	"net/url"

	"go-alfred/output"
)

type WeiboClient struct {
	Icon         string
	Host         string
	TrendingPath string
}

func (c *WeiboClient) IconPath() string {
	return c.Icon
}

func (c *WeiboClient) TrendingUrl() string {
	return c.Host + c.TrendingPath
}

func (c *WeiboClient) TrendingFormat(body []byte) ([]output.UnifiedOutput, error) {
	res := make([]output.UnifiedOutput, 0)

	resp := new(WeiboTrendingResp)
	if err := json.Unmarshal(body, resp); err != nil {
		return res, err
	}

	// TODO 判断返回code

	// resp.Data.Realtime = resp.Data.Realtime[:10]
	prefix := "https://s.weibo.com/weibo?q="
	tail := "&Refer=top"
	for _, data := range resp.Data.Realtime {
		item := output.UnifiedOutput{
			Msg:      data.Word,
			Detail:   data.WordScheme,
			Url:      prefix + url.PathEscape(data.WordScheme) + tail,
			IconPath: "resources/weibo_logo.png",
		}
		res = append(res, item)
	}
	return res, nil
}

type WeiboTrendingResp struct {
	Ok   int `json:"ok"`
	Data struct {
		Hotgov struct {
			IconDescColor      string `json:"icon_desc_color"`
			SmallIconDesc      string `json:"small_icon_desc"`
			Note               string `json:"note"`
			IsHot              int    `json:"is_hot"`
			TopicFlag          int    `json:"topic_flag"`
			IsGov              int    `json:"is_gov"`
			Word               string `json:"word"`
			Mid                string `json:"mid"`
			IconDesc           string `json:"icon_desc"`
			Name               string `json:"name"`
			SmallIconDescColor string `json:"small_icon_desc_color"`
		} `json:"hotgov"`
		Realtime []struct {
			IconDescColor      string `json:"icon_desc_color"`
			Category           string `json:"category"`
			SmallIconDescColor string `json:"small_icon_desc_color"`
			Note               string `json:"note"`
			RawHot             int    `json:"raw_hot"`
			Emoticon           string `json:"emoticon"`
			SubjectLabel       string `json:"subject_label"`
			Word               string `json:"word"`
			IsBoom             int    `json:"is_boom"`
			IconDesc           string `json:"icon_desc"`
			Num                int    `json:"num"`
			Realpos            int    `json:"realpos"`
			Flag               int    `json:"flag"`
			OnboardTime        int    `json:"onboard_time"`
			StarWord           int    `json:"star_word"`
			FunWord            int    `json:"fun_word"`
			TopicFlag          int    `json:"topic_flag"`
			Mid                string `json:"mid"`
			WordScheme         string `json:"word_scheme"`
			SmallIconDesc      string `json:"small_icon_desc"`
			SubjectQuerys      string `json:"subject_querys"`
			LabelName          string `json:"label_name"`
			Expand             int    `json:"expand"`
			ChannelType        string `json:"channel_type"`
			Rank               int    `json:"rank"`
			FlagDesc           string `json:"flag_desc"`
			IconWidth          int    `json:"icon_width"`
			AdChannel          int    `json:"ad_channel"`
			IconHeight         int    `json:"icon_height"`
			IconType           string `json:"icon_type"`
			IsAd               int    `json:"is_ad"`
			Id                 int    `json:"id"`
			DotIcon            int    `json:"dot_icon"`
			Icon               string `json:"icon"`
			Sort               int    `json:"sort"`
			IsAbtest           int    `json:"is_abtest"`
			IsTopic            int    `json:"is_topic"`
			IsStar             int    `json:"is_star"`
			UpdatedTime        string `json:"updated_time"`
			EndTime            int    `json:"end_time"`
			Poi                int    `json:"poi"`
			Name               string `json:"name"`
			AdType             string `json:"ad_type"`
			StartTime          int    `json:"start_time"`
			Logo               string `json:"logo"`
			IsNew              int    `json:"is_new"`
			IsHot              int    `json:"is_hot"`
			DiscussRoomid      int    `json:"discuss_roomid"`
			DiscussRoomGray    int    `json:"discuss_room_gray"`
		} `json:"realtime"`
	} `json:"data"`
}
