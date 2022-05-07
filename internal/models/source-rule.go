package models

type SourceRule struct {
	Name             string         `xml:"name,attr"`
	Block            string         `xml:"block"`
	TagsBlock        string         `xml:"tags_block"`
	Tags             string         `xml:"tags"`
	Title            string         `xml:"title"`
	Link             SourceRuleLink `xml:"link"`
	ShortDescription string         `xml:"short_description"`
	Author           string         `xml:"author"`
	Rating           string         `xml:"rating"`
	Date             SourceRuleDate `xml:"date"`
}

type SourceRuleLink struct {
	Href   string `xml:",chardata"`
	Prefix string `xml:"prefix,attr"`
}

type SourceRuleDate struct {
	Time      string `xml:",chardata"`
	Layout    string `xml:"layout,attr"`
	Attribute string `xml:"attribute,attr"`
}

type SourceRules []SourceRule
