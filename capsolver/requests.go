package capsolver

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type createTaskPayload struct {
	ClientKey string `json:"clientKey"`
	Task      any    `json:"task"`
}

func (c *Client) createTask(payload any) (string, error) {
	var endpoint string
	switch payload.(type) {
	case *akamaiWebPayload:
		endpoint = "/akamaiweb/invoke"
	case **akamaiWebPayload:
		endpoint = "/akamaibmp/invoke"
	default:
		endpoint = "/createTask"
	}
	resp, err := c.hClient.R().SetHeaders(map[string]string{
		"Host":         "api.capsolver.com",
		"Content-Type": "application/json",
	}).SetBody(createTaskPayload{
		ClientKey: c.apikey,
		Task:      payload,
	}).Post(endpoint)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	jsonResp := gjson.Parse(resp.String())
	if int(jsonResp.Get("errorId").Int()) != 0 {
		return "", fmt.Errorf("%s: %s", jsonResp.Get("errorCode").Str, jsonResp.Get("errorDescription").Str)
	}
	taskID := jsonResp.Get("taskId").Str
	if len(taskID) != 36 {
		return "", fmt.Errorf("invalid taskID: %v", taskID)
	}
	return taskID, nil
}

type getTaskResultPayload struct {
	ClientKey string `json:"clientKey"`
	TaskID    string `json:"taskId"`
}

func (c *Client) getTaskResult(taskID string) (string, error) {
	resp, err := c.hClient.R().SetHeaders(map[string]string{
		"Host":         "api.capsolver.com",
		"Content-Type": "application/json",
	}).SetBody(getTaskResultPayload{
		ClientKey: c.apikey,
		TaskID:    taskID,
	}).Post("/getTaskResult")
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	jsonResp := gjson.Parse(resp.String())
	if int(jsonResp.Get("errorId").Int()) != 0 {
		return "", fmt.Errorf("errorCode=%s: %s", jsonResp.Get("errorCode").Str, jsonResp.Get("errorDescription").Str)
	}
	status := jsonResp.Get("status").Str
	if status == "ready" {
		return jsonResp.Get("solution").Raw, nil
	}
	return "", fmt.Errorf("invalid status: %s", status)
}

type getBalancePayload struct {
	ClientKey string `json:"clientKey"`
}

func (c *Client) Balance() (float64, error) {
	resp, err := c.hClient.R().SetHeaders(map[string]string{
		"Host":         "api.capsolver.com",
		"Content-Type": "application/json",
	}).SetBody(getBalancePayload{
		ClientKey: c.apikey,
	}).Post("/getBalance")
	if err != nil {
		return 0, err
	}
	if resp.StatusCode() != 200 {
		return 0, fmt.Errorf("bad response: %v", resp.StatusCode())
	}
	jsonResp := gjson.Parse(resp.String())
	if int(jsonResp.Get("errorId").Int()) != 0 {
		return 0, fmt.Errorf("errorCode=%s: %s", jsonResp.Get("errorCode").Str, jsonResp.Get("errorDescription").Str)
	}
	return jsonResp.Get("balance").Float(), nil
}
