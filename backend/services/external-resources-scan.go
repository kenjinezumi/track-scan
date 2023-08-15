package services

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ExternalResources struct {
	TimeStamp time.Time
	Resources []Resource
}

type Resource struct {
	TagName          string
	Src              string
	PotentialTracker bool // Indicates if the resource is a potential tracker
}

// ExternalResourcesAnalyser analyzes external resources in a goquery.Document for potential tracking-related information
func ExternalResourcesAnalyser(doc *goquery.Document, knownNetworks ...string) ExternalResources {
	externalRes := ExternalResources{
		TimeStamp: time.Now(),
	}

	doc.Find("script[src], link[rel='stylesheet'], img[src], iframe[src], embed[src], object[data]").Each(func(index int, element *goquery.Selection) {
		tagName := element.Get(0).Data
		src, exists := element.Attr("src")

		if exists {
			externalResource := Resource{TagName: tagName, Src: src}

			// Check if the domain belongs to known ad or tracking networks
			for _, network := range knownNetworks {
				if strings.Contains(src, network) {
					externalResource.PotentialTracker = true
					externalRes.Resources = append(externalRes.Resources, externalResource)
					break
				}
			}

		}
	})

	return externalRes
}
