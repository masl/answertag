package tmpl

import (
	"html/template"
	"net/http"
	"sync"
)

type TemplateManager struct {
	mu        sync.Mutex
	templates map[string]*template.Template
}

func NewTemplateManager() *TemplateManager {
	return &TemplateManager{
		templates: make(map[string]*template.Template),
	}
}

func (tm *TemplateManager) AddTemplate(name string, tmpl *template.Template) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.templates[name] = tmpl
}

type RenderTemplateOptions struct {
	Layout string
}

func (tm *TemplateManager) RenderTemplate(w http.ResponseWriter, name string, data interface{}, opts *RenderTemplateOptions) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// get template from template manager map by name
	t, ok := tm.templates[name]
	if !ok {
		return ErrTemplateDoesNotExist
	}

	// set base as default layout

	if opts == nil {
		opts = &RenderTemplateOptions{
			Layout: "base",
		}
	}

	// execute template
	if err := t.ExecuteTemplate(w, opts.Layout, data); err != nil {
		return err
	}

	return nil
}
