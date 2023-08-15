// services/csp_inspector.go

package services

import (
	"net/http"
	"regexp"
	"strings"
	"time"
)

type CSPInspector struct{}

func NewCSPInspector() *CSPInspector {
	return &CSPInspector{}
}

type CSPReport struct {
	TimeStamp       time.Time
	Directives      map[string][]string
	TrackingDomains []string
}

func (c *CSPInspector) InspectCSP(resp *http.Response) CSPReport {
	cspHeaders := resp.Header.Values("Content-Security-Policy")
	directives := make(map[string][]string)
	trackingDomains := make([]string, 0)

	for _, cspHeader := range cspHeaders {
		for directive, values := range c.extractDirectives(cspHeader) {
			directives[directive] = append(directives[directive], values...)
		}
	}

	return CSPReport{
		TimeStamp:       time.Now(),
		Directives:      directives,
		TrackingDomains: trackingDomains,
	}
}

func (c *CSPInspector) extractDirectives(cspHeader string) map[string][]string {
	directives := make(map[string][]string)

	directivePattern := regexp.MustCompile(`(?i)\b(\w+)\s+(['"][^'"]+['"]|[^;]+)`)
	matches := directivePattern.FindAllStringSubmatch(cspHeader, -1)

	for _, match := range matches {
		directive := strings.TrimSpace(match[1])
		values := strings.Fields(match[2])
		for _, value := range values {
			if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
				directives[directive] = append(directives[directive], value)
			}
		}
	}

	return directives
}
