package docling

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) PollTaskStatus(ctx context.Context, taskID string) (AsyncResponse, error) {
	r, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("status/poll/%s", taskID), nil)
	if err != nil {
		return AsyncResponse{}, err
	}
	var resp AsyncResponse
	err = c.Do(r, &resp)
	if err != nil {
		return AsyncResponse{}, err
	}
	return resp, nil
}

func (c *Client) GetConvertTaskResult(ctx context.Context, taskID string) (ConvertResponse, error) {
	r, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("result/%s", taskID), nil)
	if err != nil {
		return ConvertResponse{}, err
	}
	var resp ConvertResponse
	err = c.Do(r, &resp)
	if err != nil {
		return ConvertResponse{}, err
	}
	return resp, nil
}
