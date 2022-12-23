package render

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (r *Render) Page(w http.ResponseWriter, req *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(r.Renderer) {
	case "go":
		return r.GoPage(w, req, view, data)
	case "jet":
		return r.JetPage(w, req, view, variables, data)
	}

	return nil
}

func (r *Render) GoPage(w http.ResponseWriter, req *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", r.RootPath, view))
	if err != nil {
		return err
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

func (r *Render) JetPage(w http.ResponseWriter, req *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	t, err := r.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		log.Print(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
