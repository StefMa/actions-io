package api

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"sort"

	"github.com/stefma/action-inputs/actions"
)

//go:embed templates/index.html
var htmlTemplate string

//go:embed templates/style.css
var css string

type htmlTemplateData struct {
	Name   string
	Inputs []struct {
		Key   string
		Value actions.ActionInput
	}
	Outputs []struct {
		Key   string
		Value actions.ActionOutput
	}
	ActionFileURL template.URL
	BadgeURL      template.URL
	CSS           template.CSS
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	config, err := actions.Config(r.URL.Query())
	if err != nil {
		http.Error(w, "Error parsing config", http.StatusInternalServerError)
		return
	}

	data := htmlTemplateData{
		Name:          fmt.Sprintf("%s@%s", config.Repository, config.Branch),
		Inputs:        sortMapKeys(config.ActionConfig.Inputs),
		Outputs:       sortMapKeys(config.ActionConfig.Outputs),
		ActionFileURL: template.URL(fmt.Sprintf("https://github.com/%s/blob/%s/%s", config.Repository, config.Branch, config.ActionFilename)),
		BadgeURL:      template.URL(fmt.Sprintf("https://%s/badge/%s@%s", r.Host, config.Repository, config.Branch)),
		CSS:           template.CSS(css),
	}

	tmpl, err := template.New("webpage").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func sortMapKeys[T any](m map[string]T) []struct {
	Key   string
	Value T
} {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Strings(keys)

	// Create a sorted slice of key-value pairs
	sortedSlice := make([]struct {
		Key   string
		Value T
	}, len(keys))
	for i, k := range keys {
		sortedSlice[i] = struct {
			Key   string
			Value T
		}{Key: k, Value: m[k]}
	}

	return sortedSlice
}
