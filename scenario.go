package autify

import (
	"context"
	"fmt"
	"time"
)

type ScenarioList struct {
	ID         int64     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	ProjectURL string    `json:"project_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Scenario struct {
	Name       string    `json:"name,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ProjectURL string    `json:"project_url,omitempty"`
	Steps      []Step    `json:"steps,omitempty"`
}

type Step struct {
	ID          int64       `json:"id,omitempty"`
	RowOrder    int         `json:"row_order,omitempty"`
	SourceType  string      `json:"source_type,omitempty"`
	StepKeyword StepKeyword `json:"step_keyword,omitempty"`
}

type StepKeyword struct {
	ID                int64             `json:"id,omitempty"`
	StepArguments     []StepArgument    `json:"step_arguments,omitempty"`
	TranslatedKeyword TranslatedKeyword `json:"translated_keyword,omitempty"`
}

type TranslatedKeyword struct {
	ID               int64     `json:"id,omitempty"`
	AbstractTemplate string    `json:"abstract_template,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	Name             string    `json:"name,omitempty"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type StepArgument struct {
	ID    int64  `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

func (c *Client) ListScenario(ctx context.Context, projectID int64, page int) ([]ScenarioList, error) {
	url := fmt.Sprintf("projects/%d/scenarios", projectID)

	if params := buildQuery(Page(&page)); params != "" {
		url += "?" + params
	}

	var targetScenarios []ScenarioList
	err := c.get(ctx, url, &targetScenarios)

	if err != nil {
		return []ScenarioList{}, err
	}

	return targetScenarios, nil
}

func (c *Client) Scenario(ctx context.Context, projectID int64, id int64) (*Scenario, error) {
	url := fmt.Sprintf("projects/%d/scenarios/%d", projectID, id)

	var result Scenario

	err := c.get(ctx, url, &result)

	if err != nil {
		return nil, nil
	}

	return &result, nil
}
