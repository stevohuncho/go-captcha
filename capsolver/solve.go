package capsolver

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Payload interface{}

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
			solution, err := c.getTaskResult(taskID)
			if err != nil {
				if strings.Contains(err.Error(), "errorCode=") {
					return "", err
				}
				continue
			}
			return solution, nil
		}
	}
}
