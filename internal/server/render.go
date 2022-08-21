package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	layoutsPath  = "./web/templates/layouts/*.layout.html"
	viewsPath    = "./web/templates/views/*.view.html"
	partialsPath = "./web/templates/partials/*.partial.html"
)

func writeView(w http.ResponseWriter, tmpl string) error {
	templateMap, err := deafultViewsCache()
	if err != nil {
		return err
	}

	v, exist := templateMap[tmpl]
	if !exist {
		return fmt.Errorf("template %s does not exist", tmpl)
	}

	buff := new(bytes.Buffer)
	err = v.Execute(buff, app.Data)

	_, err = buff.WriteTo(w)

	return nil
}

func deafultViewsCache() (map[string]*template.Template, error) {
	if app.config.ServerCache {
		return app.viewsCache, nil
	}

	viewsCache, err := loadViewsCache()
	if err != nil {
		return app.viewsCache, err
	}

	app.viewsCache = viewsCache
	return app.viewsCache, nil
}

func loadViewsCache() (map[string]*template.Template, error) {
	var tmplFuncs = template.FuncMap{
		"userSignedIn": func() bool {
			if app.Data.CurrentUser != nil {
				return true
			}

			return false
		},
	}

	vc := map[string]*template.Template{}
	views, err := filepath.Glob(viewsPath)
	if err != nil {
		return vc, err
	}

	for _, v := range views {
		name := filepath.Base(v)
		newView, err := template.New(name).Funcs(tmplFuncs).ParseFiles(v)
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
