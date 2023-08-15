package services

import (
	"net/http"
	"time"
)

// Cookie represents a single cookie
type Cookie struct {
	Name   string
	Value  string
	Domain string
}

// CookiesList stores information about cookies and their presence
type CookiesList struct {
	TimeStamp     time.Time
	HeaderCookie  bool
	Cookies       []Cookie
}

// AnalyzeCookies analyzes the cookies in an HTTP response
func AnalyzeCookies(resp *http.Response) CookiesList {
	cookiesResponse := CookiesList{
		TimeStamp:    time.Now(),
		HeaderCookie: checkCookies(resp),
	}

	if cookiesResponse.HeaderCookie {
		cookiesResponse.Cookies = extractCookies(resp.Cookies())
	}

	return cookiesResponse
}

func checkCookies(resp *http.Response) bool {
	cookieHeaders := resp.Header["Set-Cookie"]
	return len(cookieHeaders) > 0
}

func extractCookies(httpCookies []*http.Cookie) []Cookie {
	cookiesList := make([]Cookie, len(httpCookies))
	for i, cookie := range httpCookies {
		cookiesList[i] = Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Domain: cookie.Domain,
		}
	}
	return cookiesList
}
