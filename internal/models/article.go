package models

type Article struct {
	ID          int64  `json:"id"`
	ChannelID   int64  `json:"channel_id"`
	Signature   string `json:"signature"`
	HTMLMessage string `json:"html_message"`
}

type Articles []Article
