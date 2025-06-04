package httpcli

import (
	"context"
	"io"
	"net/http"
	"sync"
)

type request struct {
	*http.Request
	reader       io.Reader
	rawBody      []byte
	readBodyErr  error
	readBodyOnce sync.Once
}

type RequestOption func(req *request)

func WrapRequest(req *http.Request) *request {
	return &request{
		Request: req,
		reader:  req.Body,
	}
}

func NewRequest(ctx context.Context, method, url string, body io.Reader, opts ...RequestOption) (*request, error) {
	if v, ok := body.(iBody); ok {
		reqBody, err := v.Create()
		if err != nil {
			return nil, err
		}

		body = reqBody

	}
}
