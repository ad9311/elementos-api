package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/ad9311/hitomgr/internal/sess"
)

const (
	layoutsPath  = "./web/templates/**/*.layout.html"
	viewsPath    = "./web/templates/**/*.view.html"
	partialsPath = "./web/templates/**/*.partial.html"
)

var viewsCache map[string]*template.Template

// App ...
var App *sess.App
var cache bool

// Init ...
func Init(serverCache bool, sessionData *sess.App) error {
	cache = serverCache
	App = sessionData

	vc, err := deafultViewsCache()
	if err != nil {
		return err
	}
	viewsCache = vc

	return nil
}

// WriteView ...
func WriteView(w http.ResponseWriter, ID string) error {
	templateMap, err := deafultViewsCache()
	if err != nil {
		return err
	}

	v, exist := templateMap[ID]
	if !exist {
		return fmt.Errorf("template %s does not exist", ID)
	}

	buff := new(bytes.Buffer)
	err = v.Execute(buff, App)

	_, err = buff.WriteTo(w)
	if err != nil {
		return err
	}

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
		file := filepath.Base(v)
		newView, err := template.New(file).Funcs(templateFuncMap()).ParseFiles(v)
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
		vc[viewID(v)] = newView
	}

	return vc, err
}

func viewID(path string) string {
	dir := strings.Split(filepath.Dir(path), "/")
	action := strings.Split(filepath.Base(path), ".")
	return fmt.Sprintf("%s_%s", dir[len(dir)-1], action[0])
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatDate": func(date time.Time) string {
			return date.Format("Mon Jan 02, 03:04:05 PM")
		},
		"formatShortDate": func(date time.Time) string {
			return date.Format("2006-Jan-02")
		},
		"emptySlice": func(i interface{}) bool {
			switch reflect.TypeOf(i).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(i)
				if s.Len() > 0 {
					return false
				}
			default:
				return true
			}

			return true
		},
	}
}
