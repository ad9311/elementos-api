package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	layouts   = "./web/views/layouts/*.layout.html"
	templates = "./web/views/templates/*.template.html"
)

func writeTemplate(w http.ResponseWriter, tmpl string) error {
	templateMap, err := defaultTemplateCache()
	if err != nil {
		return err
	}

	t, exist := templateMap[tmpl]
	if !exist {
		return fmt.Errorf("template %s does not exist", tmpl)
	}

	buff := new(bytes.Buffer)
	err = t.Execute(buff, app)

	_, err = buff.WriteTo(w)

	return nil
}

func defaultTemplateCache() (map[string]*template.Template, error) {
	if app.config.ServerCache {
		return app.templateCache, nil
	}

	templateCache, err := loadTemplateCache()
	if err != nil {
		return app.templateCache, err
	}

	app.templateCache = templateCache
	return app.templateCache, nil
}

func loadTemplateCache() (map[string]*template.Template, error) {
	var tmplFuncs = template.FuncMap{}

	cache := map[string]*template.Template{}
	tmpls, err := filepath.Glob(templates)
	if err != nil {
		return cache, err
	}

	for _, t := range tmpls {
		name := filepath.Base(t)
		tmplDef, err := template.New(name).Funcs(tmplFuncs).ParseFiles(t)
		if err != nil {
			return cache, err
		}

		lays, err := filepath.Glob(layouts)
		if err != nil {
			return cache, err
		}

		if len(lays) > 0 {
			tmplDef, err = tmplDef.ParseGlob(layouts)
			if err != nil {
				return cache, err
			}
		}
		cache[name] = tmplDef
	}

	return cache, err
}
