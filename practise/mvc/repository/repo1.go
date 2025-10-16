package repo

import (
	"fmt"
	"net/http"
)

type IRepository1 interface {
	Do(req *http.Request) (*http.Response, error)
}

type Repository1 struct {
	client http.Client
}

func NewRepository(client http.Client) *Repository1 {
	return &Repository1{
		client,
	}
}

func (repo Repository1) Do(req *http.Request) (*http.Response, error) {
	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API call, err: %v", err)
	}
	return resp, nil
}
