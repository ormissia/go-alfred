package client

import (
	"encoding/json"
	"strconv"

	"go-alfred/output"
)

type ZhihuClient struct {
	Icon         string
	Host         string
	TrendingPath string
}

func (c *ZhihuClient) IconPath() string {
	return c.Icon
}

func (c *ZhihuClient) TrendingUrl() string {
	return c.Host + c.TrendingPath
}

func (c *ZhihuClient) TrendingFormat(body []byte) ([]output.UnifiedOutput, error) {
	res := make([]output.UnifiedOutput, 0)

	resp := new(ZhihuTrendingResp)
	if err := json.Unmarshal(body, resp); err != nil {
		return res, err
	}

	prefix := "https://www.zhihu.com/question/"
	for _, data := range resp.Data {
		item := output.UnifiedOutput{
			Msg:      data.Target.Title,
			Detail:   data.Target.Excerpt,
			Url:      prefix + strconv.Itoa(data.Target.Id),
			IconPath: c.IconPath(),
		}
		res = append(res, item)
	}
	return res, nil
}

type ZhihuTrendingResp struct {
	Data []struct {
		Type      string `json:"type"`
		StyleType string `json:"style_type"`
		Id        string `json:"id"`
		CardId    string `json:"card_id"`
		Target    struct {
			Id            int    `json:"id"`
			Title         string `json:"title"`
			Url           string `json:"url"`
			Type          string `json:"type"`
			Created       int    `json:"created"`
			AnswerCount   int    `json:"answer_count"`
			FollowerCount int    `json:"follower_count"`
			Author        struct {
				Type      string `json:"type"`
				UserType  string `json:"user_type"`
				Id        string `json:"id"`
				UrlToken  string `json:"url_token"`
				Url       string `json:"url"`
				Name      string `json:"name"`
				Headline  string `json:"headline"`
				AvatarUrl string `json:"avatar_url"`
			} `json:"author"`
			BoundTopicIds []int  `json:"bound_topic_ids"`
			CommentCount  int    `json:"comment_count"`
			IsFollowing   bool   `json:"is_following"`
			Excerpt       string `json:"excerpt"`
		} `json:"target"`
		AttachedInfo string `json:"attached_info"`
		DetailText   string `json:"detail_text"`
		Trend        int    `json:"trend"`
		Debut        bool   `json:"debut"`
		Children     []struct {
			Type      string `json:"type"`
			Thumbnail string `json:"thumbnail"`
		} `json:"children"`
	} `json:"data"`
	Paging struct {
		IsEnd    bool   `json:"is_end"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"paging"`
	FreshText  string `json:"fresh_text"`
	DisplayNum int    `json:"display_num"`
}
