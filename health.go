package docling

import (
	"context"
	"net/http"
)

func (c *Client) Health(ctx context.Context) (HealthResponse, error) {
	r, err := c.NewRequest(ctx, http.MethodGet, "health", nil)
	if err != nil {
		return HealthResponse{}, err
	}
	var resp HealthResponse
	err = c.Do(r, &resp)
	if err != nil {
		return HealthResponse{}, err
	}
	return resp, nil
}

type HealthResponse struct {
	Status string `json:"status"`
}
