package actions

import (
	"net/url"
	"testing"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name            string
		owner           string
		repo            string
		expectedRepo    string
		expectedBranch  string
		expectedInputs  int
		expectedOutputs int
		wantErr         bool
	}{
		{"ValidConfig1", "Blackjacx", "backlog-notifier", "Blackjacx/backlog-notifier", "main", 7, 0, false},
		{"ValidConfig2", "docker", "build-push-action", "docker/build-push-action", "master", 33, 3, false},
		{"ValidConfig3", "actions", "setup-go@v2", "actions/setup-go", "v2", 4, 0, false},
		{"ValidConfig4", "StefMa", "Upgrade-Go-Action@v1.0.0", "StefMa/Upgrade-Go-Action", "v1.0.0", 3, 0, false},
		{"ValidConfig5", "actions", "setup-java@v3.13.0", "actions/setup-java", "v3.13.0", 19, 4, false},
		{"InvalidConfig", "not-exist", "seriously@vNotAvailable", "not-exist/seriously", "vNotAvailable", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := url.Values{}
			query.Set("owner", tt.owner)
			query.Set("repo", tt.repo)

			config, err := Config(query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if config.Repository != tt.expectedRepo {
					t.Errorf("Config() Repository = %v, want %v", config.Repository, tt.expectedRepo)
				}

				if config.Branch != tt.expectedBranch {
					t.Errorf("Config() Branch = %v, want %v", config.Branch, tt.expectedBranch)
				}

				if config.ActionFilename != "action.yml" && config.ActionFilename != "action.yaml" {
					t.Errorf("Config() ActionFilename = %v, want action.yml or action.yaml", config.ActionFilename)
				}

				if len(config.ActionConfig.Inputs) != tt.expectedInputs {
					t.Errorf("Config() ActionConfig.Inputs size = %d, want %d", len(config.ActionConfig.Inputs), tt.expectedInputs)
				}
				if len(config.ActionConfig.Outputs) != tt.expectedOutputs {
					t.Errorf("Config() ActionConfig.Outputs size = %d, want %d", len(config.ActionConfig.Outputs), tt.expectedOutputs)
				}
			}
		})
	}
}
