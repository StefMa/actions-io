package api

import (
	"fmt"
	"net/http"

	"github.com/stefma/action-inputs/actions"
)

func BadgeHandler(w http.ResponseWriter, r *http.Request) {
	config, err := actions.Config(r.URL.Query())
	if err != nil {
		http.Error(w, "Error parsing config", http.StatusInternalServerError)
		return
	}

	il := len(config.ActionConfig.Inputs)
	ol := len(config.ActionConfig.Outputs)

	url := fmt.Sprintf("https://img.shields.io/badge/Actions_I/O-%d/%d-blue", il, ol)

	http.Redirect(w, r, url, http.StatusSeeOther)
}
