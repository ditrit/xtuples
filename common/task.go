package common

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type TaskConfig struct {
	Name         string `yaml:"name"`
	Agents       Agents `yaml:"agents"`
	Keys         []string
	AgentByNames map[string]Agent
}

type Agent struct {
	Filter    Filter    `yaml:"filter"`
	Job       Job       `yaml:"job"`
	OnSuccess OnSuccess `yaml:"on_success"`
}

type Agents []Agent

type Job struct {
	Name  string `yaml:"name"`
	Scale int    `yaml:"scale"`
}

type OnSuccess struct {
	OldState string `yaml:"old_state" json:"old_state,omitempty"`
	State    string `yaml:"state" json:"state,omitempty"`
	Step     string `yaml:"step" json:"step,omitempty"`
}

type Filter struct {
	Step  string   `yaml:"step" json:"step"`
	Empty []string `yaml:"empty" json:"empty,omitempty"`
	OneOf string   `yaml:"one_of" json:"one_of,omitempty"`
}

// parseTask parses a yaml string into a TaskConfig struct
func ParseConfig(configPath string) *TaskConfig {
	config, err := ReadFromFile(configPath)
	if err != nil {
		log.Fatal("Failed to read config from file:", err)
	}

	var cfg TaskConfig
	err = yaml.Unmarshal([]byte(config), &cfg)

	if err != nil {
		log.Fatal("Failed to parse config file", err)
	}

	cfg.Keys = cfg.GetKeys()
	cfg.AgentByNames = make(map[string]Agent)
	for _, agent := range cfg.Agents {
		cfg.AgentByNames[agent.Job.Name] = agent
	}

	return &cfg
}

func (a Agent) JsonFilter() (string, error) {
	jsonBytes, err := json.Marshal(a.Filter)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (a Agent) JsonOnSuccess() (string, error) {
	a.OnSuccess.OldState = a.Filter.OneOf
	fmt.Println("a.Filter.OneOf")
	fmt.Println(a.Filter.OneOf)
	fmt.Println("a.OnSuccess")
	fmt.Println(a.OnSuccess)
	fmt.Println("a.OnSuccess.OldState")
	fmt.Println(a.OnSuccess.OldState)
	fmt.Println("a.OnSuccess.State")
	fmt.Println(a.OnSuccess.State)

	jsonBytes, err := json.Marshal(a.OnSuccess)
	if err != nil {
		return "", err
	}
	fmt.Println("json :" + string(jsonBytes))
	return string(jsonBytes), nil
}

func (a Agent) GetKeys() []string {
	keys := make([]string, 0)
	uniqueKeys := make(map[string]bool)

	if len(a.Filter.Empty) > 0 {
		for _, value := range a.Filter.Empty {
			uniqueKeys[value] = true
		}
	}

	if a.Filter.OneOf != "" {
		uniqueKeys[a.Filter.OneOf] = true
	}

	if a.OnSuccess.State != "" {
		uniqueKeys[a.OnSuccess.State] = true
	}

	for value := range uniqueKeys {
		keys = append(keys, value)
	}

	return keys
}

func (t TaskConfig) GetKeys() []string {
	keys := make([]string, 0)
	uniqueKeys := make(map[string]bool)

	for _, agent := range t.Agents {
		agentKeys := agent.GetKeys()
		for _, key := range agentKeys {
			uniqueKeys[key] = true
		}
	}

	for _, key := range []string{"error", "retries", "step", "values"} {
		uniqueKeys[key] = true
	}

	for value := range uniqueKeys {
		keys = append(keys, value)
	}

	return keys
}
