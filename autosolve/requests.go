package autosolve

import (
	"fmt"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

func (c *Client) auth() error {
	c.authWg = &sync.WaitGroup{}
	c.authWg.Add(1)
	defer c.authWg.Done()
	resp, err := c.hClient.R().SetQueryParam(
		"apiKey", c.apikey,
	).Get("https://autosolve-dashboard-api.aycd.io/api/v1/auth/generate-token?")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	auth := gjson.Parse(resp.String())
	c.authToken = auth.Get("token").Str
	c.authExp = time.Unix(auth.Get("expiresAt").Int(), 0)
	return nil
}

func (c *Client) createTask(payload any) (string, error) {
	c.checkAuth()
	resp, err := c.hClient.R().SetHeader(formatAuth(c.authToken)).SetBody(payload).Post("https://autosolve-api.aycd.io/api/v1/tasks/create")
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	return "", nil
}

func (c *Client) getTasks() error {
	c.checkAuth()
	resp, err := c.hClient.R().SetHeader(formatAuth(c.authToken)).Get("https://autosolve-api.aycd.io/api/v1/tasks")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	tasks := gjson.Parse(resp.String()).Array()
	for _, t := range tasks {
		c.tasks.Set(t.Get("taskId").Str, task{
			TaskId:    t.Get("taskId").Str,
			CreatedAt: t.Get("createdAt").Int(),
			Status:    t.Get("status").Str,
			Token:     t.Get("token").Str,
		})
	}
	return nil
}

func (c *Client) checkAuth() error {
	if time.Until(c.authExp) < time.Second*30 {
		if c.authWg != nil {
			c.authWg.Wait()
			return c.checkAuth()
		} else {
			err := c.auth()
			if err != nil {
				return err
			}
		}
	}
	c.authWg = nil
	return nil
}

func formatAuth(authToken string) (string, string) {
	return "Authorization", fmt.Sprintf("Token %s", authToken)
}
