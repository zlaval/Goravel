package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool = false, false

	minutes, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		minutes = 60
	}
	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	} else {
		secure = false
	}

	session := scs.New()

	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = s.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {
	case "redis":
	case "mariadb, mysql":
	default:
		//Cookie
	}

	return session
}
