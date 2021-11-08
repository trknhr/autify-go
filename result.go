package autify

import (
	"context"
	"fmt"
	"time"
)

type ResultOptions struct {
	Page       *int
	PerPage    *int
	TestPlanID *int
}

type Result struct {
	ID           int64      `json:"id,omitempty"`
	Status       string     `json:"status,omitempty"`
	Duration     int        `json:"duration,omitempty"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	FinishedAt   *time.Time `json:"finished_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	ReviewNeeded bool       `json:"review_needed,omitempty"`
	TestPlan     TestPlan   `json:"test_plan,omitempty"`
}

type TestPlan struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ResultDetail struct {
	ID                        int64                      `json:"id,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	FinishedAt                *time.Time                 `json:"finished_at,omitempty"`
	StartedAt                 *time.Time                 `json:"started_at,omitempty"`
	Duration                  int                        `json:"duration,omitempty"`
	TestPlanCapabilityResults []TestPlanCapabilityResult `json:"test_plan_capability_results,omitempty"`
	TestPlan                  TestPlan                   `json:"test_plan,omitempty"`
	CreatedAt                 time.Time                  `json:"created_at,omitempty"`
	UpdatedAt                 time.Time                  `json:"updated_at,omitempty"`
}

type TestPlanCapabilityResult struct {
	ID              int64            `json:"id,omitempty"`
	Capability      Capability       `json:"capability,omitempty"`
	TestCaseResults []TestCaseResult `json:"test_case_results,omitempty"`
}

type Capability struct {
	Browser        string    `json:"browser,omitempty"`
	BrowserVersion string    `json:"browser_version,omitempty"`
	Device         string    `json:"device,omitempty"`
	Resolution     string    `json:"resolution,omitempty"`
	OS             string    `json:"os,omitempty"`
	OSVersion      string    `json:"os_version,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type TestCaseResult struct {
	ID           int64      `json:"id,omitempty"`
	Duration     int        `json:"duration,omitempty"`
	ProjectURL   string     `json:"project_url,omitempty"`
	ReviewNeeded int        `json:"review_needed,omitempty"`
	Status       string     `json:"status,omitempty"`
	TestCaseID   int64      `json:"test_case_id,omitempty"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	FinishedAt   *time.Time `json:"finished_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
}

func (c *Client) Results(ctx context.Context, projectID int64, opt *ResultOptions) ([]Result, error) {
	url := fmt.Sprintf("projects/%d/results", projectID)

	if opt != nil {
		if params := buildQuery(Page(opt.Page), PerPage(opt.PerPage), TestPlanID(opt.TestPlanID)); params != "" {
			url += "?" + params
		}
	}

	var result []Result

	err := c.get(ctx, url, &result)

	if err != nil {
		return []Result{}, err
	}

	return result, nil
}

func (c *Client) Result(ctx context.Context, projectID int64, resultID int64) (*ResultDetail, error) {
	url := fmt.Sprintf("projects/%d/results/%d", projectID, resultID)

	var result ResultDetail

	err := c.get(ctx, url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
