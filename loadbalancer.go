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
	APIPathLoadBalancers = "loadbalancers"
)

func (c *Client) GetAllLoadBalancers(ctx context.Context) ([]ResponseLoadBalancer, error) {
	req, err := c.newRequest(ctx, http.MethodGet, APIPathLoadBalancers, nil)
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

	var ret []ResponseLoadBalancer
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetLoadBalancer(ctx context.Context, name string) (*ResponseLoadBalancer, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path.Join(APIPathLoadBalancers, name), nil)
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

	var ret ResponseLoadBalancer
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateLoadBalancer(ctx context.Context, loadbalancer *RequestLoadBalancer) (*ResponseLoadBalancer, error) {
	b, err := json.Marshal(loadbalancer)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, APIPathLoadBalancers, bytes.NewBuffer(b))
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

	var ret ResponseLoadBalancer
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) UpdateLoadBalancer(ctx context.Context, name string, loadbalancer *RequestLoadBalancer) (*ResponseLoadBalancer, error) {
	b, err := json.Marshal(loadbalancer)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path.Join(APIPathLoadBalancers, name), bytes.NewBuffer(b))
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

	var ret ResponseLoadBalancer
	if err := decodeBody(res, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) DeleteLoadBalancer(ctx context.Context, name string) (bool, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, path.Join(APIPathLoadBalancers, name), nil)
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
