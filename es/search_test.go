package es

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type mockRequest struct {
	mockStatusCode   int
	mockResponseBody io.ReadCloser
}

func (mock mockRequest) Do(ctx context.Context, transport esapi.Transport) (*esapi.Response, error) {
	return &esapi.Response{
		StatusCode: mock.mockStatusCode,
		Body:       mock.mockResponseBody,
	}, nil
}

func TestSearch(t *testing.T) {
	client, err := NewEsClient()
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		name             string
		statusCode       int
		responseFilePath string
		err              error
	}{
		{
			"OK_Response",
			200,
			"testdata/ok_es_response.json",
			nil,
		},
		{
			"NG_Response",
			404,
			"testdata/ng_es_response.json",
			fmt.Errorf("[404] index_not_found_exception: no such index [test]"),
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			bt, err := os.ReadFile(tt.responseFilePath)
			if err != nil {
				t.Fatal(err)
			}
			r := io.NopCloser(bytes.NewBuffer(bt))
			req := mockRequest{
				mockStatusCode:   tt.statusCode,
				mockResponseBody: r,
			}
			_, err = client.Search(context.Background(), req)
			if err != tt.err {
				t.Errorf("got: %v, want: %v", err, tt.err)
			}
		})
	}
}
