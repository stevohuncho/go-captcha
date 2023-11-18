package autosolve

import (
	"context"
	"fmt"
	"time"
)

type Payload struct {
	TaskId           string            `json:"taskId"`
	Url              string            `json:"url"`
	SiteKey          string            `json:"siteKey"`
	Version          int               `json:"version"`
	Action           string            `json:"action"`
	MinScore         float32           `json:"minScore"`
	Proxy            string            `json:"proxy"`
	ProxyRequired    bool              `json:"proxyRequired"`
	UserAgent        string            `json:"userAgent"`
	Cookies          string            `json:"cookies"`
	RenderParameters map[string]string `json:"renderParameters"`
	Metadata         map[string]string `json:"metadata"`
}

func (c *Client) Solve(payload Payload, polling int, ctx context.Context) (string, error) {
	taskID, err := c.createTask(payload)
	if err != nil {
		return "", err
	}
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("timed out")
		case <-time.After(time.Millisecond * time.Duration(polling)):
			task, ok := c.tasks.Get(taskID)
			if ok {
				if task.Status == "success" {
					c.tasks.Delete(taskID)
					return task.Token, nil
				}
			}
		}
	}
}
