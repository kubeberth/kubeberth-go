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
	APIPathISOImages = "isoimages"
)

func (c *Client) GetAllISOImages(ctx context.Context) ([]ResponseISOImage, error) {
	req, err := c.newRequest(ctx, http.MethodGet, APIPathISOImages, nil)
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

	var ret []ResponseISOImage
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetISOImage(ctx context.Context, name string) (*ResponseISOImage, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path.Join(APIPathISOImages, name), nil)
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

	var ret ResponseISOImage
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateISOImage(ctx context.Context, isoimage *RequestISOImage) (*ResponseISOImage, error) {
	b, err := json.Marshal(isoimage)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, APIPathISOImages, bytes.NewBuffer(b))
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

	var ret ResponseISOImage
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) UpdateISOImage(ctx context.Context, name string, isoimage *RequestISOImage) (*ResponseISOImage, error) {
	b, err := json.Marshal(isoimage)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path.Join(APIPathISOImages, name), bytes.NewBuffer(b))
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

	var ret ResponseISOImage
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) DeleteISOImage(ctx context.Context, name string) (bool, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path.Join(APIPathISOImages, name), nil)
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
