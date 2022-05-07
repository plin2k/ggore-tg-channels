package models

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Content struct {
	ID               int       `db:"id"`
	Title            string    `db:"title"`
	Link             string    `db:"link"`
	ShortDescription string    `db:"short_description"`
	Author           string    `db:"author"`
	Rating           string    `db:"rating"`
	Date             time.Time `db:"date"`
	Tags             []string  `db:"tags"`
}

type Contents []Content

func (cs Contents) Reverse() {
	for i, j := 0, len(cs)-1; i < j; i, j = i+1, j-1 {
		cs[i], cs[j] = cs[j], cs[i]
	}
}

func (cs Contents) Shuffle() Contents {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cs), func(i, j int) { cs[i], cs[j] = cs[j], cs[i] })

	return cs
}

func (c *Content) AddTags(tags ...string) {
	c.Tags = append(tags, c.Tags...)
}

func (c *Content) GetSignature(channelName string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(c.Title+c.Link+channelName)))
}

func (cs Contents) ToArticles(channel *Channel) Articles {
	articles := make(Articles, len(cs))

	for i := range cs {
		articles[i] = cs[i].ToArticle(channel)
	}

	return articles
}

func (c *Content) ToArticle(channel *Channel) Article {
	var article Article

	article.Signature = c.GetSignature(channel.Name)
	article.ChannelID = channel.ChannelID

	article.HTMLMessage = c.ToHTMLByStructure(channel.ArticleStructure)

	return article
}

func (c *Content) IsNotEmpty() bool {
	if c.Title != "" && c.Link != "" {
		return true
	}

	return false
}

var specialCharsReplacer = strings.NewReplacer(
	"-", "_",
	".", "_",
	",", "",
	" ", "_",
	"#", "",
	"\\", "_",
	"/", "_",
	"|", "_",
	"@", "",
	"$", "",
	"%", "",
	"^", "",
	"&", "",
	"*", "",
	"(", "",
	")", "",
	"~", "",
	"`", "",
	"\"", "",
	"[", "",
	"]", "",
	"{", "",
	"}", "",
	"§", "",
	"±", "",
	">", "",
	"<", "",
	";", "",
	":", "",
	"'", "",
	"?", "",
	"+", "",
	"=", "",
	"»", "",
	"«", "",
)

func (c *Content) RemoveDoubleTags() {
	inResult := make(map[string]struct{})
	var result []string

	for _, tag := range c.Tags {
		if _, ok := inResult[tag]; !ok {
			inResult[tag] = struct{}{}
			result = append(result, tag)
		}
	}

	c.Tags = result
}

func (c *Content) GetTagsString() string {
	if len(c.Tags) == 0 {
		return ""
	}

	var tags strings.Builder

	for i := range c.Tags {
		tags.WriteString("#")
		tags.WriteString(specialCharsReplacer.Replace(c.Tags[i]))
		tags.WriteString(" ")
	}

	return tags.String()
}

var monthEnToRu = strings.NewReplacer(
	"January", "Январь",
	"February", "Февраль",
	"March", "Март",
	"April", "Апрель",
	"May", "Май",
	"June", "Июнь",
	"July", "Июль",
	"August", "Август",
	"September", "Сентябрь",
	"October", "Октябрь",
	"November", "Ноябрь",
	"December", "Декабрь",
)

func (c *Content) GetDateString() string {
	if c.Date.Year() > 2015 {
		return monthEnToRu.Replace(c.Date.Format("2 January 2006"))
	}

	return ""
}

func (c *Content) ToHTMLByStructure(structure ArticleStructure) string {
	var result strings.Builder

	for _, row := range structure.Row {
		switch strings.ToLower(row.Field) {
		case "title":
			result.WriteString(row.HandleContent(c.Title))
		case "tags":
			c.RemoveDoubleTags()
			result.WriteString(row.HandleContent(c.GetTagsString()))
		case "link":
			result.WriteString(row.HandleContent(c.Link))
		case "short_description":
			result.WriteString(row.HandleContent(c.ShortDescription))
		case "author":
			result.WriteString(row.HandleContent(c.Author))
		case "rating":
			result.WriteString(row.HandleContent(c.Rating))
		case "date":
			result.WriteString(row.HandleContent(c.GetDateString()))
		default:
			result.WriteString(row.HandleContent(row.Field))
		}
	}

	return result.String()
}

//func handleContentRulePartlyEnglish(str string) string {
//	if len(str) == 0 {
//		return str
//	}
//
//	strArr := strings.Fields(str)
//
//	newStr := strArr[0]
//
//	for _, word := range strArr[1:] {
//		if rand.Intn(10) > 0 {
//			enStr, err := repo.Translate.GetTranslateFromRU(strings.ToLower(word))
//			if err != nil {
//				log.Fatalln(err)
//			}
//
//			if len(enStr) > 0 {
//				newStr += " " + enStr
//
//				continue
//			}
//		}
//
//		newStr += " " + word
//	}
//
//	return newStr
//}
