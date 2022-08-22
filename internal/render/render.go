package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ad9311/hitomgr/internal/sess"
)

const (
	layoutsPath  = "./web/templates/layouts/*.layout.html"
	viewsPath    = "./web/templates/views/*.view.html"
	partialsPath = "./web/templates/partials/*.partial.html"
)

var viewsCache map[string]*template.Template

// Data ...
var Data *sess.Data
var cache bool

// Init ...
func Init(serverCache bool, sessionData *sess.Data) error {
	cache = serverCache
	Data = sessionData

	vc, err := deafultViewsCache()
	if err != nil {
		return err
	}
	viewsCache = vc

	return nil
}

// WriteView ...
func WriteView(w http.ResponseWriter, tmpl string) error {
	templateMap, err := deafultViewsCache()
	if err != nil {
		return err
	}

	v, exist := templateMap[tmpl]
	if !exist {
		return fmt.Errorf("template %s does not exist", tmpl)
	}

	buff := new(bytes.Buffer)
	err = v.Execute(buff, Data)

	_, err = buff.WriteTo(w)

	return nil
}

func deafultViewsCache() (map[string]*template.Template, error) {
	if cache {
		return viewsCache, nil
	}

	vc, err := loadViewsCache()
	if err != nil {
		return viewsCache, err
	}

	viewsCache = vc
	return viewsCache, nil
}

func loadViewsCache() (map[string]*template.Template, error) {
	vc := map[string]*template.Template{}
	views, err := filepath.Glob(viewsPath)
	if err != nil {
		return vc, err
	}

	for _, v := range views {
		name := filepath.Base(v)
		newView, err := template.New(name).Funcs(templateFuncMap()).ParseFiles(v)
		if err != nil {
			return vc, err
		}

		layouts, err := filepath.Glob(layoutsPath)
		if err != nil {
			return vc, err
		}

		partials, err := filepath.Glob(partialsPath)
		if err != nil {
			return vc, err
		}

		if (len(layouts) > 0) && (len(partials) > 0) {
			newView, err = newView.ParseGlob(layoutsPath)
			newView, err = newView.ParseGlob(partialsPath)
			if err != nil {
				return vc, err
			}
		}
		vc[name] = newView
	}

	return vc, err
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatDate": func(date time.Time) string {
			return date.Format("Mon Jan 02, 03:04:05 PM")
		},
		"formatShortDate": func(date time.Time) string {
			return date.Format("2006-Jan-02")
		},
	}
}
