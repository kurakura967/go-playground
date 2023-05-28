package es

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type esClient struct {
	client *elasticsearch.Client
}

func NewEsClient() (*esClient, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &esClient{
		client: es,
	}, nil
}

func (e *esClient) Search(ctx context.Context, req esapi.Request) ([]byte, error) {
	res, err := req.Do(ctx, e.client)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		var e errorResponse
		if e := json.NewDecoder(res.Body).Decode(&e); e != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[%d] %s: %s", res.StatusCode, e.Error.Type, e.Error.Reason)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type errorResponse struct {
	Error esError `json:"error"`
}

type esError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}
