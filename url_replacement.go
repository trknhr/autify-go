package autify

import (
	"context"
	"fmt"
	"time"
)

type URLReplacementResponse struct {
	ID             int64     `json:"id,omitempty"`
	TestPlanID     int64     `json:"test_plan_id,omitempty"`
	PatternURL     string    `json:"pattern_url,omitempty"`
	ReplacementURL string    `json:"replacement_url,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type URLReplacementOptions struct {
	PatternURL     string `json:"pattern_url,omitempty"`
	ReplacementURL string `json:"replacement_url,omitempty"`
}

func (c *Client) CreateNewURLReplacement(ctx context.Context, testPlanID int64, opt *URLReplacementOptions) (*URLReplacementResponse, error) {
	url := fmt.Sprintf("test_plans/%d/url_replacements", testPlanID)

	var result URLReplacementResponse

	err := c.post(ctx, url, opt, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) URLReplacementList(ctx context.Context, testPlanID int64) ([]URLReplacementResponse, error) {
	url := fmt.Sprintf("test_plans/%d/url_replacements", testPlanID)

	var result []URLReplacementResponse

	err := c.get(ctx, url, &result)

	if err != nil {
		return []URLReplacementResponse{}, err
	}

	return result, nil
}

func (c *Client) DeleteURLReplacement(ctx context.Context, testPlanID, urlReplacementID int64) (string, error) {
	url := fmt.Sprintf("test_plans/%d/url_replacements/%d", testPlanID, urlReplacementID)

	var result string

	err := c.delete(ctx, url, &result)

	if err != nil {
		return "", err
	}

	return result, nil
}
