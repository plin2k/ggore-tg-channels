package models

import (
	"math/rand"
	"strings"
	"time"
)

type Source struct {
	Link     string   `xml:",chardata"`
	RuleName string   `xml:"rule,attr"`
	Tags     []string `xml:"tags,attr"`

	Rule *SourceRule
}

type Sources []Source

func (s *Source) ConstructTags() {
	s.Tags = strings.Split(s.Tags[0], ",")
}

func (ss Sources) ConstructTags() {
	for i := range ss {
		ss[i].ConstructTags()
	}
}

func (ss Sources) Shuffle() Sources {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ss), func(i, j int) { ss[i], ss[j] = ss[j], ss[i] })

	return ss
}
