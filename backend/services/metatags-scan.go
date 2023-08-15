package services

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type MetaTags struct {
	TimeStamp time.Time
	Tags      []Tags
}

type Tags struct {
	Name    string
	Content string
}

// MetaTagsAnalyser analyzes the meta tags in a goquery.Document for potential tracking-related information
func MetaTagsAnalyser(doc *goquery.Document, keywords ...string) MetaTags {
	metaTagsRes := MetaTags{
		TimeStamp: time.Now(),
	}

	doc.Find("meta[name='referrer'], meta[name='description']").Each(func(index int, element *goquery.Selection) {
		name := element.AttrOr("name", "")
		content := element.AttrOr("content", "")

		if len(keywords) == 0 {
			// No specific keywords provided, check for any content
			metaTagsRes.Tags = append(metaTagsRes.Tags, Tags{Name: name, Content: content})
		} else {
			for _, keyword := range keywords {
				if strings.Contains(strings.ToLower(content), strings.ToLower(keyword)) {
					metaTagsRes.Tags = append(metaTagsRes.Tags, Tags{Name: name, Content: content})
					break
				}
			}
		}
	})

	return metaTagsRes
}
