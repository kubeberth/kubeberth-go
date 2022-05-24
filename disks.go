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
	APIPathDisks = "disks"
)

func (c *Client) GetAllDisks(ctx context.Context) ([]Disk, error) {
	req, err := c.newRequest(ctx, http.MethodGet, APIPathDisks, nil)
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

	var ret []Disk
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetDisk(ctx context.Context, name string) (*Disk, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path.Join(APIPathDisks, name), nil)
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

	var ret Disk
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateDisk(ctx context.Context, disk *Disk) (*Disk, error) {
	b, err := json.Marshal(disk)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, APIPathDisks, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(res.Status)
	}

	var ret Disk
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) UpdateDisk(ctx context.Context, name string, disk *Disk) (*Disk, error) {
	b, err := json.Marshal(disk)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path.Join(APIPathDisks, name), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(res.Status)
	}

	var ret Disk
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) DeleteDisk(ctx context.Context, name string) (bool, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path.Join(APIPathDisks, name), nil)
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
