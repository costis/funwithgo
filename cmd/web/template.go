package main

import (
	"html/template"
	"path/filepath"
)

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tpl, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		tpl, err = tpl.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}
		tpl, err = tpl.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}

		cache[name] = tpl
	}

	return cache, nil
}
