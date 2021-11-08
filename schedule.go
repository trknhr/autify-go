package autify

import (
	"context"
	"fmt"
)

type ScheduleResult struct {
	Data ScheduleResultDetail `json:"data,omitempty"`
}

type ScheduleResultDetail struct {
	ID         string            `json:"id,omitempty"`
	Type       string            `json:"type,omitempty"`
	Attributes ScheduleAttribute `json:"attributes,omitempty"`
}
type ScheduleAttribute struct {
	ID int64 `json:"id,omitempty"`
}

func (c *Client) CreateSchedule(ctx context.Context, scheduleID int64) (*ScheduleResult, error) {
	url := fmt.Sprintf("schedules/%d", scheduleID)

	var result ScheduleResult

	err := c.post(ctx, url, nil, &result)

	if err != nil {
		return nil, nil
	}

	return &result, nil
}
