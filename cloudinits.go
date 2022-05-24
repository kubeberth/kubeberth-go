package kubeberth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"

	"github.com/kubeberth/berth-apiserver/pkg/cloudinits"
)

const (
	APIPathCloudInits = "cloudinits"
)

type CloudInit cloudinits.CloudInit

func (c *Client) GetAllCloudInits(ctx context.Context) ([]CloudInit, error) {
	req, err := c.newRequest(ctx, http.MethodGet, APIPathCloudInits, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var ret []CloudInit
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetCloudInit(ctx context.Context, name string) (*CloudInit, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path.Join(APIPathCloudInits, name), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var ret CloudInit
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateCloudInit(ctx context.Context, cloudinit *CloudInit) (*CloudInit, error) {
	b, err := json.Marshal(cloudinit)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, APIPathCloudInits, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var ret CloudInit
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) UpdateCloudInit(ctx context.Context, name string, cloudinit *CloudInit) (*CloudInit, error) {
	b, err := json.Marshal(cloudinit)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path.Join(APIPathCloudInits, name), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	var ret CloudInit
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) DeleteCloudInit(ctx context.Context, name string) (bool, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path.Join(APIPathCloudInits, name), nil)
	if err != nil {
		return false, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, errors.New(res.Status)
	}

	return true, nil
}
