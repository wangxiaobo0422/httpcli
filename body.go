package httpcli

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

const (
	ContentTypeJson     = "application/json"
	ContentTypeForm     = "application/x-www-form-urlencoded"
	ContentTypeTextHtml = "text/html"
)

type iBody interface {
	ContentType() string
	Create() (io.Reader, error)
}

type JsonBody struct {
	io.Reader
	Val any
}

func ToJsonBody(v any) *JsonBody {
	return &JsonBody{
		Val: v,
	}
}

func (j *JsonBody) ContentType() string {
	return ContentTypeJson
}

func (j *JsonBody) Create() (io.Reader, error) {
	bs, err := json.Marshal(j.Val)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(bs), nil
}

type FormBody struct {
	io.Reader
	Val url.Values
}

func ToFormBody(v url.Values) *FormBody {
	return &FormBody{
		Val: v,
	}
}

func (f *FormBody) ContentType() string {
	return ContentTypeForm
}

func (f *FormBody) Create() (io.Reader, error) {
	return bytes.NewBufferString(f.Val.Encode()), nil
}
