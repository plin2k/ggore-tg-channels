package parser

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/plin2k/ggore-tg-channels/internal/models"

	"github.com/PuerkitoBio/goquery"
)

type service struct {
	client *http.Client
}

func New() *service {
	return &service{
		client: http.DefaultClient,
	}
}

//func (s *service) Get

const (
	ErrorContentNotFound = "content not found"
)

// Parse ...
func (s *service) Parse(source models.Source) (models.Contents, error) {

	var contents models.Contents

	res, err := s.client.Get(source.Link)
	if err != nil {
		return nil, fmt.Errorf("parse requestSourse : %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("parse requestSourse : not equal to 200")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("parse goquery.NewDocumentFromReader : %w", err)
	}

	selection := doc.Find(source.Rule.Block)
	if len(selection.Nodes) == 0 {
		err = errors.New(source.Link + " | " + ErrorContentNotFound)
		return nil, err
	}

	selection.Each(func(i int, s *goquery.Selection) {

		var content = models.Content{
			Title:            getText(s, source.Rule.Title),
			Link:             getLink(s, source.Rule.Link),
			ShortDescription: getText(s, source.Rule.ShortDescription),
			Author:           getText(s, source.Rule.Author),
			Rating:           getText(s, source.Rule.Rating),
			Date:             getDate(s, source.Rule.Date),
			Tags:             getArray(s, source.Rule.TagsBlock, source.Rule.Tags),
		}

		content.AddTags(source.Tags...)

		if content.IsNotEmpty() {
			contents = append(contents, content)
		}
	})
	return contents, nil
}

func getText(s *goquery.Selection, rule string) string {
	return strings.TrimSpace(s.Find(rule).Text())
}

func getLink(s *goquery.Selection, rule models.SourceRuleLink) string {
	if link, exists := s.Find(rule.Href).Attr("href"); exists {
		if rule.Prefix != "" {
			return rule.Prefix + link
		}
		return link
	}
	return ""
}

func getArray(s *goquery.Selection, blockRule string, itemRule string) []string {
	var array []string

	selection := s.Find(blockRule)
	if len(selection.Nodes) == 0 {
		return array
	}

	selection.Each(func(i int, s *goquery.Selection) {
		array = append(array, s.Find(itemRule).Text())
	})

	return array
}

func getDate(s *goquery.Selection, rule models.SourceRuleDate) time.Time {
	var dateStr string
	if dt, ok := s.Find(rule.Time).Attr(rule.Attribute); ok {
		dateStr = dt
	} else {
		dateStr = s.Find(rule.Time).Text()
	}

	date, err := time.Parse(rule.Layout, strings.TrimSpace(dateStr))
	if err != nil {
		return time.Time{}
	}
	return date
}
