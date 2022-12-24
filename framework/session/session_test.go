package session

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {

	s := &Session{
		CookieLifetime: "100",
		CookieDomain:   "localhost",
		CookieName:     "goravel",
		CookieSecure:   "false",
		SessionType:    "cookie",
		CookiePersist:  "true",
	}

	var sm *scs.SessionManager

	ses := s.InitSession()

	var sessKind reflect.Kind
	var sesType reflect.Type

	rv := reflect.ValueOf(ses)

	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For loop", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sesType = rv.Type()
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("invalid type or kind. kind:", rv.Kind(), "type:", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned testing cookie session. Expected", reflect.ValueOf(sm).Kind(), "and got", sessKind)
	}

	if sesType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned testing cookie session. Expected", reflect.ValueOf(sm).Type(), "and got", sesType)
	}

}
