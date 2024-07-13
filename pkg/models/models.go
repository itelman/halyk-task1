package models

import (
	"io/ioutil"
	"net/http"

	"github.com/gofrs/uuid"
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Response struct {
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

func (r *Response) Set(request Request) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return err
	}

	for key, value := range request.Headers {
		httpReq.Header.Set(key, value)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	r.ID = id.String()
	r.Status = httpResp.StatusCode
	r.Headers = httpResp.Header
	r.Length = len(body)

	return nil
}
