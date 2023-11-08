package tmpl

import (
	"embed"
	"html/template"
)

func RegisterTemplates(tm *TemplateManager, templateFS embed.FS) error {
	// index
	t, err := template.ParseFS(templateFS,
		"templates/base.tmpl",
		"templates/head.tmpl",
		"templates/index.tmpl",
	)
	if err != nil {
		return err
	}

	tm.AddTemplate("index", t)

	// c
	t, err = template.ParseFS(templateFS,
		"templates/base.tmpl",
		"templates/head.tmpl",
		"templates/cloud.tmpl",
		"templates/tags.tmpl",
	)
	if err != nil {
		return err
	}

	tm.AddTemplate("cloud", t)

	// tags
	t, err = template.ParseFS(templateFS,
		"templates/tags.tmpl",
	)
	if err != nil {
		return err
	}

	tm.AddTemplate("tags", t)

	return nil
}
