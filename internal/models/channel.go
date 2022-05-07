package models

import (
	"encoding/xml"
)

type Channel struct {
	XMLName xml.Name `xml:"xml"`

	ChannelID int64 `xml:"channel_id"`

	Title string `xml:"title"`
	Name  string `xml:"name"`

	Sources     Sources     `xml:"source"`
	SourceRules SourceRules `xml:"rule"`

	ArticleStructure ArticleStructure `xml:"article_structure"`
}

type Channels []Channel

func (c *Channel) GetSourceRule(s Source) SourceRule {
	for _, sr := range c.SourceRules {
		if s.RuleName == sr.Name {
			return sr
		}
	}
	return SourceRule{}
}

func (c *Channel) ConstructSourceRules() {
	for i, s := range c.Sources {
		sr := c.GetSourceRule(s)
		c.Sources[i].Rule = &sr
	}
}

func (c *Channel) Construct() {
	c.ConstructSourceRules()
	c.Sources.ConstructTags()
}

func (cs Channels) Construct() {
	for i := range cs {
		cs[i].Construct()
	}
}

func (cs Channels) GetLimitMap(limit int) map[int64]int {
	var m = make(map[int64]int, len(cs))

	for _, c := range cs {
		m[c.ChannelID] = limit
	}

	return m
}
