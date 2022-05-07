package models

import (
	"fmt"
	"strings"
)

type Row struct {
	Field string `xml:",chardata"`

	Prefix  string `xml:"prefix,attr"`
	Postfix string `xml:"postfix,attr"`

	TopMargin    int `xml:"top_margin,attr"`
	BottomMargin int `xml:"bottom_margin,attr"`

	Bold      bool `xml:"bold,attr"`
	Italic    bool `xml:"italic,attr"`
	Underline bool `xml:"underline,attr"`

	CSS string `xml:"css,attr"`

	TextType string `xml:"text_type,attr"`

	Features []string `xml:"features,attr"`
}

type Rows []Row

type ArticleStructure struct {
	Row Rows `xml:"row"`
}

func (r *Row) TextTypeHTML(text string) string {
	if r == nil {
		return text
	}

	switch strings.ToLower(r.TextType) {
	case "title":
		return fmt.Sprintf("<h1>%s</h1>", text)
	case "heading":
		return fmt.Sprintf("<h2>%s</h2>", text)
	case "subheading":
		return fmt.Sprintf("<h3>%s</h3>", text)
	default:
		return text

	}
}

func (r *Row) CSSHTML(text string) string {
	if r == nil {
		return text
	}

	if len(r.CSS) > 1 {
		return fmt.Sprintf("<div style=\"%s\">%s</div>", r.CSS, text)
	}

	return text
}

func (r *Row) BoldHTML(text string) string {
	if r == nil {
		return text
	}

	if r.Bold {
		return fmt.Sprintf("<b>%s</b>", text)
	}

	return text
}

func (r *Row) ItalicHTML(text string) string {
	if r == nil {
		return text
	}

	if r.Italic {
		return fmt.Sprintf("<i>%s</i>", text)
	}

	return text
}

func (r *Row) UnderlineHTML(text string) string {
	if r == nil {
		return text
	}

	if r.Underline {
		return fmt.Sprintf("<u>%s</u>", text)
	}

	return text
}

func (r *Row) PrefixHTML(text string) string {
	if r == nil {
		return text
	}

	if len(r.Prefix) > 0 {
		return fmt.Sprintf("%s%s", r.Prefix, text)
	}

	return text
}

func (r *Row) PostfixHTML(text string) string {
	if r == nil {
		return text
	}

	if len(r.Postfix) > 0 {
		return fmt.Sprintf("%s%s", text, r.Prefix)
	}

	return text
}

func (r *Row) TopMarginHTML(text string) string {
	if r == nil {
		return text
	}

	if r.TopMargin < 1 {
		return text
	}

	return fmt.Sprintf("%s%s", strings.Repeat("\n\r", r.TopMargin), text)
}

func (r *Row) BottomMarginHTML(text string) string {
	if r == nil {
		return text
	}

	if r.BottomMargin < 1 {
		return text
	}

	return fmt.Sprintf("%s%s", text, strings.Repeat("\n\r", r.BottomMargin))
}

func (r *Row) HandleContent(text string) string {
	if r == nil {
		return text
	}

	if len(text) < 1 {
		return ""
	}

	var result string = text

	result = r.PrefixHTML(result)
	result = r.PostfixHTML(result)

	result = r.BoldHTML(result)
	result = r.ItalicHTML(result)
	result = r.UnderlineHTML(result)

	result = r.TextTypeHTML(result)

	result = r.CSSHTML(result)

	result = r.TopMarginHTML(result)
	result = r.BottomMarginHTML(result)

	return result + "\n\r"
}
