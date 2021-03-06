package kubeberth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
)

const (
	APIPathArchives = "archives"
)

func (c *Client) GetAllArchives(ctx context.Context) ([]ResponseArchive, error) {
	req, err := c.newRequest(ctx, http.MethodGet, APIPathArchives, nil)
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

	var ret []ResponseArchive
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetArchive(ctx context.Context, name string) (*ResponseArchive, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path.Join(APIPathArchives, name), nil)
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

	var ret ResponseArchive
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateArchive(ctx context.Context, archive *RequestArchive) (*ResponseArchive, error) {
	b, err := json.Marshal(archive)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, APIPathArchives, bytes.NewBuffer(b))
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

	var ret ResponseArchive
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) UpdateArchive(ctx context.Context, name string, archive *RequestArchive) (*ResponseArchive, error) {
	b, err := json.Marshal(archive)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path.Join(APIPathArchives, name), bytes.NewBuffer(b))
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

	var ret ResponseArchive
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) DeleteArchive(ctx context.Context, name string) (bool, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path.Join(APIPathArchives, name), nil)
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
