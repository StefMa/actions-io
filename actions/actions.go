package actions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/yaml.v3"
)

type GitHubConfig struct {
	Repository     string // Format as owner/repo
	Branch         string // Either `main` or a tag
	ActionFilename string // Either `action.yml` or `action.yaml`
	ActionConfig   ActionConfig
}

type ActionConfig struct {
	Inputs  map[string]ActionInput  `yaml:"inputs"`
	Outputs map[string]ActionOutput `yaml:"outputs"`
}

type ActionInput struct {
	Description        string `yaml:"description"`
	Required           bool   `yaml:"required"`
	Default            string `yaml:"default"`
	DeprecationMessage string `yaml:"deprecationMessage"`
}

type ActionOutput struct {
	Description string `yaml:"description"`
}

func Config(query url.Values) (*GitHubConfig, error) {
	owner, repo, branch := extractOwnerRepoBranch(query)

	filename, body, err := fetchActionFileContent(owner, repo, branch)
	if err != nil {
		return nil, fmt.Errorf("error fetching action file content. see %w", err)
	}

	var config ActionConfig
	if err := yaml.Unmarshal(body, &config); err != nil {
		return nil, fmt.Errorf("error parsing action file. see %w", err)
	}

	return &GitHubConfig{
		Repository:     owner + "/" + repo,
		Branch:         branch,
		ActionConfig:   config,
		ActionFilename: filename,
	}, nil
}

func extractOwnerRepoBranch(query url.Values) (owner, repo, branch string) {
	owner = query.Get("owner")
	repo = query.Get("repo")
	repoParts := strings.Split(repo, "@")

	branch = "main"
	if len(repoParts) == 2 {
		repo = repoParts[0]
		branch = repoParts[1]
	}
	return owner, repo, branch
}

func fetchActionFileContent(owner, repo, branch string) (filename string, body []byte, err error) {
	filename = "action.yml"
	actionsFileUrl := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", owner, repo, branch, filename)

	resp, err := http.Get(actionsFileUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		filename = "action.yaml"
		actionsFileUrl = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", owner, repo, branch, filename)
		resp, err = http.Get(actionsFileUrl)
		if err != nil || resp.StatusCode != http.StatusOK {
			return "", nil, fmt.Errorf("error fetching action file. see %w", err)
		}
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error reading action file. see %w", err)
	}
	return filename, body, nil
}
